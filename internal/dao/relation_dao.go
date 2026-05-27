package dao

import (
	"context"
	"filler/internal/database"
	"filler/internal/dto"
)

func BindParam(ctx context.Context, d dto.BindParamRequest) error {
	conn := database.Get()
	query := `INSERT INTO public.component_param (component_id, param_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := conn.Exec(ctx, query, d.ComponentID, d.ParamID)
	return err
}

func UnbindParam(ctx context.Context, componentID, paramID int64) (bool, error) {
	conn := database.Get()
	query := `DELETE FROM public.component_param WHERE component_id = $1 AND param_id = $2`
	commandTag, err := conn.Exec(ctx, query, componentID, paramID)
	if err != nil {
		return false, err
	}
	return commandTag.RowsAffected() > 0, nil
}
