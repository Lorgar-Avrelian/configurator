package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func BindParam(ctx context.Context, componentID int64, paramID int64) error {
	var conn *pgxpool.Pool
	var query string
	var err error
	conn = database.Get()
	query = `INSERT INTO public.component_param ("component_id", "param_id")
			 VALUES ($1, $2)
			 ON CONFLICT ("component_id", "param_id") DO NOTHING`
	_, err = conn.Exec(ctx, query, componentID, paramID)
	if err != nil {
		logger.Error("Failed to bind parameter ID %d to component ID %d: %v", paramID, componentID, err)
		return err
	}
	return nil
}

func UnbindParam(ctx context.Context, componentID int64, paramID int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	var affected int64
	conn = database.Get()
	query = `DELETE FROM public.component_param
			 WHERE "component_id" = $1 AND "param_id" = $2`
	commandTag, err = conn.Exec(ctx, query, componentID, paramID)
	if err != nil {
		logger.Error("Failed to unbind parameter ID %d from component ID %d: %v", paramID, componentID, err)
		return false, err
	}
	affected = commandTag.RowsAffected()
	return affected > 0, nil
}
