package dao

import (
	"configurator/internal/database"
	"configurator/internal/dto"
	"configurator/internal/model"
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
)

type thresholdFlatRow struct {
	ID                  int64
	SourceModel         int64
	SourceInternalOrder int64
	SourceParam         string
	Value               string
	Type                int16
	Operator            int16
	Enabled             bool
	TargetParam         sql.NullString
	TargetDevice        sql.NullInt64
	Level               int16
	PrevOperator        int16
	PreviousID          sql.NullInt64
}

func scanThresholdRows(rows pgx.Rows) ([]thresholdFlatRow, error) {
	var list []thresholdFlatRow
	for rows.Next() {
		var r thresholdFlatRow
		err := rows.Scan(
			&r.ID, &r.SourceModel, &r.SourceInternalOrder, &r.SourceParam, &r.Value,
			&r.Type, &r.Operator, &r.Enabled, &r.TargetParam, &r.TargetDevice,
			&r.Level, &r.PrevOperator, &r.PreviousID,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, nil
}

func assembleThresholdChain(flatRows []thresholdFlatRow, startID int64) *model.Threshold {
	nodes := make(map[int64]*model.Threshold)
	for _, r := range flatRows {
		t := &model.Threshold{
			ID:                  r.ID,
			SourceModel:         r.SourceModel,
			SourceInternalOrder: r.SourceInternalOrder,
			SourceParam:         r.SourceParam,
			Value:               r.Value,
			Type:                model.VarType(r.Type),
			Operator:            model.LogicOperator(r.Operator),
			Enabled:             r.Enabled,
			Level:               model.AlarmLevel(r.Level),
			PrevOperator:        model.LogicOperator(r.PrevOperator),
		}
		if r.TargetParam.Valid {
			t.TargetParam = &r.TargetParam.String
		}
		if r.TargetDevice.Valid {
			t.TargetDevice = &r.TargetDevice.Int64
		}
		if r.PreviousID.Valid {
			t.PreviousID = &r.PreviousID.Int64
		}
		nodes[r.ID] = t
	}
	for _, node := range nodes {
		if node.PreviousID != nil {
			if prevNode, ok := nodes[*node.PreviousID]; ok {
				node.PreviousThreshold = prevNode
			}
		}
	}
	return nodes[startID]
}

func CreateThreshold(ctx context.Context, d dto.ThresholdCreate) (int64, error) {
	conn := database.Get()
	query := `INSERT INTO public.threshold (source_model, source_internal_order, source_param, value, type, operator, enabled, target_param, target_device, level, prev_operator, previous) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id`
	typeIdx := model.ParseVarType(d.Type)
	opIdx := model.ParseLogicOperator(d.Operator)
	lvlIdx := model.ParseAlarmLevel(d.Level)
	prevOpIdx := model.ParseLogicOperator(d.PrevOperator)
	enabled := true
	if d.Enabled != nil {
		enabled = *d.Enabled
	}
	var prevID sql.NullInt64
	if d.PreviousID != nil {
		prevID.Int64 = *d.PreviousID
		prevID.Valid = true
	}
	var id int64
	err := conn.QueryRow(ctx, query, d.SourceModel, d.SourceInternalOrder, d.SourceParam, d.Value, int16(typeIdx), int16(opIdx), enabled, toNullString(d.TargetParam), toNullInt64(d.TargetDevice), int16(lvlIdx), int16(prevOpIdx), prevID).Scan(&id)
	return id, err
}

func GetDetailedThresholdByID(ctx context.Context, id int64) (*model.Threshold, error) {
	conn := database.Get()
	query := `
		WITH RECURSIVE threshold_chain AS (
			SELECT id, source_model, source_internal_order, source_param, value, type, operator, enabled, target_param, target_device, level, prev_operator, previous 
			FROM public.threshold WHERE id = $1
			UNION ALL
			SELECT t.id, t.source_model, t.source_internal_order, t.source_param, t.value, t.type, t.operator, t.enabled, t.target_param, t.target_device, t.level, t.prev_operator, t.previous
			FROM public.threshold t
			JOIN threshold_chain tc ON t.id = tc.previous
		)
		SELECT id, source_model, source_internal_order, source_param, value, type, operator, enabled, target_param, target_device, level, prev_operator, previous FROM threshold_chain`
	rows, err := conn.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flatRows, err := scanThresholdRows(rows)
	if err != nil || len(flatRows) == 0 {
		return nil, err
	}
	return assembleThresholdChain(flatRows, id), nil
}

func UpdateThreshold(ctx context.Context, id int64, d dto.ThresholdUpdate) (int64, error) {
	conn := database.Get()
	query := `UPDATE public.threshold SET source_model = $1, source_internal_order = $2, source_param = $3, value = $4, type = $5, operator = $6, enabled = $7, target_param = $8, target_device = $9, level = $10, prev_operator = $11, previous = $12 WHERE id = $13 RETURNING id`
	typeIdx := model.ParseVarType(d.Type)
	opIdx := model.ParseLogicOperator(d.Operator)
	lvlIdx := model.ParseAlarmLevel(d.Level)
	prevOpIdx := model.ParseLogicOperator(d.PrevOperator)
	enabled := true
	if d.Enabled != nil {
		enabled = *d.Enabled
	}
	var prevID sql.NullInt64
	if d.PreviousID != nil {
		prevID.Int64 = *d.PreviousID
		prevID.Valid = true
	}
	var updatedID int64
	err := conn.QueryRow(ctx, query, d.SourceModel, d.SourceInternalOrder, d.SourceParam, d.Value, int16(typeIdx), int16(opIdx), enabled, toNullString(d.TargetParam), toNullInt64(d.TargetDevice), int16(lvlIdx), int16(prevOpIdx), prevID, id).Scan(&updatedID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return updatedID, nil
}

func DeleteThreshold(ctx context.Context, id int64) (bool, error) {
	conn := database.Get()
	query := `
		WITH RECURSIVE target_thresholds AS (
			SELECT id FROM public.threshold WHERE id = $1
			UNION ALL
			SELECT t.id FROM public.threshold t
			JOIN target_thresholds tt ON t.previous = tt.id
		),
		pre_update_refs AS (
			UPDATE public.threshold
			SET previous = NULL
			WHERE previous IN (SELECT id FROM target_thresholds)
			RETURNING id
		)
		DELETE FROM public.threshold
		WHERE id IN (SELECT id FROM target_thresholds)`
	commandTag, err := conn.Exec(ctx, query, id)
	if err != nil {
		return false, err
	}
	return commandTag.RowsAffected() > 0, nil
}

func GetAllThresholdsExpanded(ctx context.Context) ([]model.Threshold, error) {
	conn := database.Get()
	query := `SELECT id, source_model, source_internal_order, source_param, value, type, operator, enabled, target_param, target_device, level, prev_operator, previous FROM public.threshold`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flatRows, err := scanThresholdRows(rows)
	if err != nil {
		return nil, err
	}
	var list []model.Threshold
	for _, r := range flatRows {
		chain := assembleThresholdChain(flatRows, r.ID)
		if chain != nil {
			list = append(list, *chain)
		}
	}
	return list, nil
}

func toNullInt64(i *int64) sql.NullInt64 {
	if i == nil {
		return sql.NullInt64{}
	}
	return sql.NullInt64{Int64: *i, Valid: true}
}
