package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func scanConfigInProcessRows(rows pgx.Rows) ([]ConfigInProcessDao, error) {
	var list []ConfigInProcessDao
	list = []ConfigInProcessDao{}
	var cfg ConfigInProcessDao
	var err error
	for rows.Next() {
		err = rows.Scan(&cfg.Host, &cfg.Port, &cfg.Protocol)
		if err != nil {
			return nil, err
		}
		list = append(list, cfg)
	}
	return list, rows.Err()
}

func GetAllConfigInProcess(ctx context.Context) ([]ConfigInProcessDao, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT "host", "port", "protocol" FROM public.config_in_process`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch all config_in_process entries: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []ConfigInProcessDao
	res, err = scanConfigInProcessRows(rows)
	return res, err
}

func SearchConfigInProcess(ctx context.Context, host *string, port *int32) ([]ConfigInProcessDao, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT "host", "port", "protocol" FROM public.config_in_process WHERE ($1::VARCHAR IS NULL OR "host" = $1) AND ($2::INTEGER IS NULL OR "port" = $2)`
	rows, err = conn.Query(ctx, query, host, port)
	if err != nil {
		logger.Error("Failed to search config_in_process entries: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []ConfigInProcessDao
	res, err = scanConfigInProcessRows(rows)
	return res, err
}
