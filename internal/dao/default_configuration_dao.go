package dao

import (
	"context"
	"database/sql"
	"errors"
	"filler/internal/database"
	"filler/internal/dto"
	"filler/internal/model"

	"github.com/jackc/pgx/v5"
)

func CreateDefaultConfiguration(ctx context.Context, d dto.ConfigurationCreate) (int64, error) {
	conn := database.Get()
	query := `INSERT INTO public.default_configuration (indicator, device_component_id) VALUES ($1, $2) RETURNING id`
	var dcID sql.NullInt64
	if d.DeviceComponentID != nil {
		dcID.Int64 = *d.DeviceComponentID
		dcID.Valid = true
	}
	var id int64
	err := conn.QueryRow(ctx, query, d.IndicatorID, dcID).Scan(&id)
	return id, err
}

func GetDetailedDefaultConfigByID(ctx context.Context, id int64) (*model.DeviceIndicator, *model.DeviceComponent, error) {
	conn := database.Get()
	query := `
		WITH RECURSIVE target_configs AS (
			SELECT id, indicator, device_component_id FROM public.default_configuration WHERE id = $1
		),
		device_tree AS (
			SELECT dc.id, dc.model, dc.internal_order, dc.parent, tc.id AS cfg_id
			FROM public.device_component dc
			JOIN target_configs tc ON dc.id = tc.device_component_id
			UNION ALL
			SELECT c.id, c.model, c.internal_order, c.parent, dt.cfg_id
			FROM public.device_component c
			JOIN device_tree dt ON c.parent = dt.id
		)
		SELECT 
			cfg.id,
			i.id, i.description, i.object_id, i.contact, i.name, i.location, i.services,
			dt.id, dt.model, dt.internal_order, dt.parent,
			json_strip_nulls(json_agg(json_build_object(
				'id', m.id,
				'indicator_id', m.indicator,
				'param_id', m.param,
				'frequency', m.frequency,
				'coefficient', m.coefficient,
				'enum', m.enum
			))) FILTER (WHERE m.id IS NOT NULL) AS mappings_json
		FROM target_configs cfg
		LEFT JOIN public.device_indicator i ON cfg.indicator = i.id
		LEFT JOIN device_tree dt ON cfg.id = dt.cfg_id
		LEFT JOIN public.device_component_mapping dcm ON dt.id = dcm.device_component_id
		LEFT JOIN public.mapping m ON dcm.mapping_id = m.id
		GROUP BY cfg.id, i.id, i.description, i.object_id, i.contact, i.name, i.location, i.services, dt.id, dt.model, dt.internal_order, dt.parent`
	rows, err := conn.Query(ctx, query, id)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	flatRows, err := ScanGenericConfigRows(rows)
	if err != nil || len(flatRows) == 0 {
		return nil, nil, err
	}
	_, indMap, dcMap := AssembleConfigurations(flatRows)
	return indMap[id], dcMap[id], nil
}

func UpdateDefaultConfiguration(ctx context.Context, id int64, d dto.ConfigurationUpdate) (int64, error) {
	conn := database.Get()
	query := `UPDATE public.default_configuration SET indicator = $1, device_component_id = $2 WHERE id = $3 RETURNING id`
	var dcID sql.NullInt64
	if d.DeviceComponentID != nil {
		dcID.Int64 = *d.DeviceComponentID
		dcID.Valid = true
	}
	var updatedID int64
	err := conn.QueryRow(ctx, query, d.IndicatorID, dcID, id).Scan(&updatedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return updatedID, nil
}

func DeleteDefaultConfiguration(ctx context.Context, id int64) (bool, error) {
	conn := database.Get()
	query := `DELETE FROM public.default_configuration WHERE id = $1`
	commandTag, err := conn.Exec(ctx, query, id)
	if err != nil {
		return false, err
	}
	return commandTag.RowsAffected() > 0, nil
}

func GetExpandedDefaultConfigurations(ctx context.Context) ([]model.DefaultConfiguration, error) {
	conn := database.Get()
	query := `
		WITH RECURSIVE device_tree AS (
			SELECT dc.id, dc.model, dc.internal_order, dc.parent, tc.id AS cfg_id
			FROM public.device_component dc
			JOIN public.default_configuration tc ON dc.id = tc.device_component_id
			UNION ALL
			SELECT c.id, c.model, c.internal_order, c.parent, dt.cfg_id
			FROM public.device_component c
			JOIN device_tree dt ON c.parent = dt.id
		)
		SELECT 
			cfg.id,
			i.id, i.description, i.object_id, i.contact, i.name, i.location, i.services,
			dt.id, dt.model, dt.internal_order, dt.parent,
			json_strip_nulls(json_agg(json_build_object(
				'id', m.id,
				'indicator_id', m.indicator,
				'param_id', m.param,
				'frequency', m.frequency,
				'coefficient', m.coefficient,
				'enum', m.enum
			))) FILTER (WHERE m.id IS NOT NULL) AS mappings_json
		FROM public.default_configuration cfg
		LEFT JOIN public.device_indicator i ON cfg.indicator = i.id
		LEFT JOIN device_tree dt ON cfg.id = dt.cfg_id
		LEFT JOIN public.device_component_mapping dcm ON dt.id = dcm.device_component_id
		LEFT JOIN public.mapping m ON dcm.mapping_id = m.id
		GROUP BY cfg.id, i.id, i.description, i.object_id, i.contact, i.name, i.location, i.services, dt.id, dt.model, dt.internal_order, dt.parent
		ORDER BY cfg.id`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flatRows, err := ScanGenericConfigRows(rows)
	if err != nil {
		return nil, err
	}
	ids, indMap, dcMap := AssembleConfigurations(flatRows)
	var list []model.DefaultConfiguration
	for _, id := range ids {
		list = append(list, model.DefaultConfiguration{ID: id, Indicator: indMap[id], DeviceComponent: dcMap[id]})
	}
	return list, nil
}

func BindDefaultConfigurationThreshold(ctx context.Context, d dto.BindDefaultConfigThresholdRequest) error {
	conn := database.Get()
	query := `INSERT INTO public.default_configuration_threshold (default_configuration_id, threshold_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := conn.Exec(ctx, query, d.DefaultConfigurationID, d.ThresholdID)
	return err
}

func UnbindDefaultConfigurationThreshold(ctx context.Context, defaultConfigID, thresholdID int64) (bool, error) {
	conn := database.Get()
	query := `DELETE FROM public.default_configuration_threshold WHERE default_configuration_id = $1 AND threshold_id = $2`
	commandTag, err := conn.Exec(ctx, query, defaultConfigID, thresholdID)
	if err != nil {
		return false, err
	}
	return commandTag.RowsAffected() > 0, nil
}
