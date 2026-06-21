package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func BindDeviceMapping(ctx context.Context, dcID int64, mID int64) error {
	var conn *pgxpool.Pool
	var query string
	var err error
	conn = database.Get()
	query = `INSERT INTO public.device_component_mapping ("device_component_id", "mapping_id")
			 VALUES ($1, $2)
			 ON CONFLICT ("device_component_id", "mapping_id") DO NOTHING`
	_, err = conn.Exec(ctx, query, dcID, mID)
	if err != nil {
		logger.Error("Failed to bind device component ID %d to mapping ID %d: %v", dcID, mID, err)
		return err
	}
	return nil
}

func UnbindDeviceMapping(ctx context.Context, dcID int64, mID int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	var affected int64
	conn = database.Get()
	query = `DELETE FROM public.device_component_mapping
			 WHERE "device_component_id" = $1 AND "mapping_id" = $2`
	commandTag, err = conn.Exec(ctx, query, dcID, mID)
	if err != nil {
		logger.Error("Failed to unbind device component ID %d from mapping ID %d: %v", dcID, mID, err)
		return false, err
	}
	affected = commandTag.RowsAffected()
	return affected > 0, nil
}
