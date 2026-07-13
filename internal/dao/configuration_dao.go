package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"configurator/internal/util"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func scanConfigurationRows(rows pgx.Rows) ([]Configuration, error) {
	var list []Configuration
	list = []Configuration{}
	var cfg Configuration
	var uPtr *uuid.UUID
	var err error
	for rows.Next() {
		uPtr = nil
		err = rows.Scan(&cfg.ID, &cfg.IndicatorID, &cfg.IndDescription, &cfg.IndContact, &cfg.IndName, &cfg.IndLocation, &cfg.IndServices, &cfg.DeviceComponentID, &cfg.CompModel, &cfg.CompInternalOrder, &cfg.CompParent, &cfg.CompTitle, &cfg.CompNameEn, &cfg.CompNameRu, &cfg.MapID, &cfg.MapIndicatorID, &cfg.MapParamID, &cfg.MapFrequency, &cfg.MapValue, &cfg.MapCoefficient, &cfg.MapEnum, &cfg.MapPosition, &cfg.MapFrom, &cfg.MapPositionType, &uPtr, &cfg.MapIndDotter, &cfg.MapOidMibID, &cfg.MapOidMibPath, &cfg.MapOidMibName, &cfg.MapOidMibVendor, &cfg.MapOidType, &cfg.MapOidName, &cfg.MapOidNumber, &cfg.MapOidDotter, &cfg.MapOidDescriptor, &cfg.MapOidSyntax, &cfg.MapOidEnum, &cfg.MapOidStatus, &cfg.MapOidAccess, &cfg.MapOidUnits, &cfg.MapOidDescription, &cfg.MapOidCategory, &cfg.MapParamTitle, &cfg.MapParamNameEn, &cfg.MapParamNameRu, &cfg.MapParamType, &cfg.MapParamValue, &cfg.MapParamDescEn, &cfg.MapParamDescRu, &cfg.MapParamUnitsEn, &cfg.MapParamUnitsRu, &cfg.MapParamAccess, &cfg.MapParamSaved, &cfg.MapParamVisible)
		if err != nil {
			return nil, err
		}
		cfg.MapIndOidID = uPtr
		list = append(list, cfg)
	}
	return list, rows.Err()
}

func CreateConfiguration(ctx context.Context, indicatorID int64, devCompID *int64) ([]Configuration, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `WITH RECURSIVE ins AS (
		INSERT INTO public.configuration ("indicator","device_component_id") VALUES ($1,$2) RETURNING "id","indicator","device_component_id"
	), dt AS (
		SELECT ins."id" AS "config_id", dc."id" AS "dc_id", dc."model", dc."internal_order", dc."parent" FROM ins JOIN public.device_component dc ON ins."device_component_id" = dc."id"
		UNION ALL
		SELECT dt."config_id", dc."id", dc."model", dc."internal_order", dc."parent" FROM public.device_component dc JOIN dt ON dc."parent" = dt."dc_id"
	), mb AS (
		SELECT dt."config_id", dt."dc_id", m."id" AS "m_id", m."indicator" AS "m_ind", m."param" AS "m_par", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM dt JOIN public.device_component_mapping dcm ON dt."dc_id" = dcm."device_component_id" JOIN public.mapping m ON dcm."mapping_id" = m."id"
		UNION ALL
		SELECT mb."config_id", mb."dc_id", m."id", m."indicator", m."param", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM public.mapping m JOIN mb ON m."from" = mb."m_id"
	)
	SELECT dc."id", dc."indicator", di."description", di."contact", di."name", di."location", di."services", dt."dc_id", dt."model", dt."internal_order", dt."parent", c."title", c."name_en", c."name_ru", mb."m_id", mb."m_ind", mb."m_par", mb."frequency", mb."value", mb."coefficient", mb."enum", mb."position", mb."from", mb."position_type", pi."oid_id", pi."dotter_notation", o."mib", mib."path", mib."name", mib."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
	FROM ins dc
	LEFT JOIN public.device_indicator di ON dc."indicator" = di."id"
	LEFT JOIN dt ON dc."id" = dt."config_id"
	LEFT JOIN public.component c ON dt."model" = c."id"
	LEFT JOIN mb ON dt."dc_id" = mb."dc_id" AND dt."config_id" = mb."config_id"
	LEFT JOIN public.param_indicator pi ON mb."m_ind" = pi."id"
	LEFT JOIN public.oid o ON pi."oid_id" = o."id"
	LEFT JOIN public.mib mib ON o."mib" = mib."id"
	LEFT JOIN public.param p ON mb."m_par" = p."id"`
	var sqlDc interface{}
	sqlDc = util.Int64PtrToSqlNullInt64(devCompID)
	rows, err = conn.Query(ctx, query, indicatorID, sqlDc)
	if err != nil {
		logger.Error("Failed to execute data-modifying create configuration query: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []Configuration
	res, err = scanConfigurationRows(rows)
	return res, err
}

func GetConfigurationByID(ctx context.Context, id int64) ([]Configuration, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `WITH RECURSIVE dt AS (
		SELECT dc."id" AS "config_id", dc."indicator" AS "ind_id", dcom."id" AS "dc_id", dcom."model", dcom."internal_order", dcom."parent" FROM public.configuration dc JOIN public.device_component dcom ON dc."device_component_id" = dcom."id" WHERE dc."id" = $1
		UNION ALL
		SELECT dt."config_id", dt."ind_id", dc."id", dc."model", dc."internal_order", dc."parent" FROM public.device_component dc JOIN dt ON dc."parent" = dt."dc_id"
	), mb AS (
		SELECT dt."config_id", dt."dc_id", m."id" AS "m_id", m."indicator" AS "m_ind", m."param" AS "m_par", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM dt JOIN public.device_component_mapping dcm ON dt."dc_id" = dcm."device_component_id" JOIN public.mapping m ON dcm."mapping_id" = m."id"
		UNION ALL
		SELECT mb."config_id", mb."dc_id", m."id", m."indicator", m."param", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM public.mapping m JOIN mb ON m."from" = mb."m_id"
	)
	SELECT dc."id", dc."indicator", di."description", di."contact", di."name", di."location", di."services", dt."dc_id", dt."model", dt."internal_order", dt."parent", c."title", c."name_en", c."name_ru", mb."m_id", mb."m_ind", mb."m_par", mb."frequency", mb."value", mb."coefficient", mb."enum", mb."position", mb."from", mb."position_type", pi."oid_id", pi."dotter_notation", o."mib", mib."path", mib."name", mib."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
	FROM public.configuration dc
	LEFT JOIN public.device_indicator di ON dc."indicator" = di."id"
	LEFT JOIN dt ON dc."id" = dt."config_id"
	LEFT JOIN public.component c ON dt."model" = c."id"
	LEFT JOIN mb ON dt."dc_id" = mb."dc_id" AND dt."config_id" = mb."config_id"
	LEFT JOIN public.param_indicator pi ON mb."m_ind" = pi."id"
	LEFT JOIN public.oid o ON pi."oid_id" = o."id"
	LEFT JOIN public.mib mib ON o."mib" = mib."id"
	LEFT JOIN public.param p ON mb."m_par" = p."id"
	WHERE dc."id" = $1`
	rows, err = conn.Query(ctx, query, id)
	if err != nil {
		logger.Error("Failed to execute hierarchical select configuration query: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []Configuration
	res, err = scanConfigurationRows(rows)
	return res, err
}

func GetAllConfigurations(ctx context.Context) ([]Configuration, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `WITH RECURSIVE dt AS (
		SELECT dc."id" AS "config_id", dc."indicator" AS "ind_id", dcom."id" AS "dc_id", dcom."model", dcom."internal_order", dcom."parent" FROM public.configuration dc JOIN public.device_component dcom ON dc."device_component_id" = dcom."id"
		UNION ALL
		SELECT dt."config_id", dt."ind_id", dc."id", dc."model", dc."internal_order", dc."parent" FROM public.device_component dc JOIN dt ON dc."parent" = dt."dc_id"
	), mb AS (
		SELECT dt."config_id", dt."dc_id", m."id" AS "m_id", m."indicator" AS "m_ind", m."param" AS "m_par", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM dt JOIN public.device_component_mapping dcm ON dt."dc_id" = dcm."device_component_id" JOIN public.mapping m ON dcm."mapping_id" = m."id"
		UNION ALL
		SELECT mb."config_id", mb."dc_id", m."id", m."indicator", m."param", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM public.mapping m JOIN mb ON m."from" = mb."m_id"
	)
	SELECT dc."id", dc."indicator", di."description", di."contact", di."name", di."location", di."services", dt."dc_id", dt."model", dt."internal_order", dt."parent", c."title", c."name_en", c."name_ru", mb."m_id", mb."m_ind", mb."m_par", mb."frequency", mb."value", mb."coefficient", mb."enum", mb."position", mb."from", mb."position_type", pi."oid_id", pi."dotter_notation", o."mib", mib."path", mib."name", mib."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
	FROM public.configuration dc
	LEFT JOIN public.device_indicator di ON dc."indicator" = di."id"
	LEFT JOIN dt ON dc."id" = dt."config_id"
	LEFT JOIN public.component c ON dt."model" = c."id"
	LEFT JOIN mb ON dt."dc_id" = mb."dc_id" AND dt."config_id" = mb."config_id"
	LEFT JOIN public.param_indicator pi ON mb."m_ind" = pi."id"
	LEFT JOIN public.oid o ON pi."oid_id" = o."id"
	LEFT JOIN public.mib mib ON o."mib" = mib."id"
	LEFT JOIN public.param p ON mb."m_par" = p."id"
	ORDER BY dc."id" ASC`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch all configurations hierarchically: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []Configuration
	res, err = scanConfigurationRows(rows)
	return res, err
}

func UpdateConfiguration(ctx context.Context, id int64, indicatorID int64, devCompID *int64) ([]Configuration, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `WITH RECURSIVE upd AS (
		UPDATE public.configuration SET "indicator" = $1, "device_component_id" = $2 WHERE "id" = $3 RETURNING "id", "indicator", "device_component_id"
	), dt AS (
		SELECT upd."id" AS "config_id", dc."id" AS "dc_id", dc."model", dc."internal_order", dc."parent" FROM upd JOIN public.device_component dc ON upd."device_component_id" = dc."id"
		UNION ALL
		SELECT dt."config_id", dc."id", dc."model", dc."internal_order", dc."parent" FROM public.device_component dc JOIN dt ON dc."parent" = dt."dc_id"
	), mb AS (
		SELECT dt."config_id", dt."dc_id", m."id" AS "m_id", m."indicator" AS "m_ind", m."param" AS "m_par", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM dt JOIN public.device_component_mapping dcm ON dt."dc_id" = dcm."device_component_id" JOIN public.mapping m ON dcm."mapping_id" = m."id"
		UNION ALL
		SELECT mb."config_id", mb."dc_id", m."id", m."indicator", m."param", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type" FROM public.mapping m JOIN mb ON m."from" = mb."m_id"
	)
	SELECT dc."id", dc."indicator", di."description", di."contact", di."name", di."location", di."services", dt."dc_id", dt."model", dt."internal_order", dt."parent", c."title", c."name_en", c."name_ru", mb."m_id", mb."m_ind", mb."m_par", mb."frequency", mb."value", mb."coefficient", mb."enum", mb."position", mb."from", mb."position_type", pi."oid_id", pi."dotter_notation", o."mib", mib."path", mib."name", mib."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
	FROM upd dc
	LEFT JOIN public.device_indicator di ON dc."indicator" = di."id"
	LEFT JOIN dt ON dc."id" = dt."config_id"
	LEFT JOIN public.component c ON dt."model" = c."id"
	LEFT JOIN mb ON dt."dc_id" = mb."dc_id" AND dt."config_id" = mb."config_id"
	LEFT JOIN public.param_indicator pi ON mb."m_ind" = pi."id"
	LEFT JOIN public.oid o ON pi."oid_id" = o."id"
	LEFT JOIN public.mib mib ON o."mib" = mib."id"
	LEFT JOIN public.param p ON mb."m_par" = p."id"`
	var sqlDc interface{}
	sqlDc = util.Int64PtrToSqlNullInt64(devCompID)
	rows, err = conn.Query(ctx, query, indicatorID, sqlDc, id)
	if err != nil {
		logger.Error("Failed to execute data-modifying update configuration query: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []Configuration
	res, err = scanConfigurationRows(rows)
	return res, err
}

func DeleteConfiguration(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var seqQuery string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	conn = database.Get()
	query = `DELETE FROM public.configuration WHERE "id" = $1`
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete configuration ID %d: %v", id, err)
		return false, err
	}
	var affected int64
	affected = commandTag.RowsAffected()
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.configuration', 'id'), COALESCE(MAX("id"), 0) + 1, false) FROM public.configuration`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Error("Failed to reset configuration sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}
