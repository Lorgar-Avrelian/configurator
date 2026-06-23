package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func scanThresholdRows(rows pgx.Rows) ([]ThresholdDao, error) {
	var list []ThresholdDao
	list = []ThresholdDao{}
	var cfg ThresholdDao
	var err error
	for rows.Next() {
		err = rows.Scan(&cfg.ID, &cfg.Name, &cfg.Description, &cfg.Author, &cfg.Created, &cfg.Query)
		if err != nil {
			return nil, err
		}
		list = append(list, cfg)
	}
	return list, rows.Err()
}

func CreateThreshold(ctx context.Context, t ThresholdDao) (int64, error) {
	var conn *pgxpool.Pool
	var query string
	var insertedID int64
	var err error
	conn = database.Get()
	query = `INSERT INTO public.threshold ("name", "description", "author", "created", "query") VALUES ($1, $2, $3, NOW(), $4) RETURNING "id"`
	err = conn.QueryRow(ctx, query, t.Name, t.Description, t.Author, t.Query).Scan(&insertedID)
	if err != nil {
		logger.Error("Failed to insert threshold into DB: %v", err)
		return 0, err
	}
	return insertedID, nil
}

func GetThresholdByID(ctx context.Context, id int64) (*ThresholdDao, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT "id", "name", "description", "author", "created", "query" FROM public.threshold WHERE "id" = $1`
	rows, err = conn.Query(ctx, query, id)
	if err != nil {
		logger.Error("Failed to query threshold ID %d: %v", id, err)
		return nil, err
	}
	defer rows.Close()
	var res []ThresholdDao
	res, err = scanThresholdRows(rows)
	if err != nil {
		return nil, err
	}
	if 0 >= len(res) {
		return nil, nil
	}
	return &res[0], nil
}

func GetAllThresholds(ctx context.Context) ([]ThresholdDao, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT "id", "name", "description", "author", "created", "query" FROM public.threshold ORDER BY "id" ASC`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch all thresholds from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []ThresholdDao
	res, err = scanThresholdRows(rows)
	return res, err
}

func UpdateThreshold(ctx context.Context, id int64, t ThresholdDao) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	conn = database.Get()
	query = `UPDATE public.threshold SET "name" = $1, "description" = $2, "author" = $3, "query" = $4 WHERE "id" = $5`
	commandTag, err = conn.Exec(ctx, query, t.Name, t.Description, t.Author, t.Query, id)
	if err != nil {
		logger.Error("Failed to update threshold ID %d: %v", id, err)
		return false, err
	}
	var affected int64
	affected = commandTag.RowsAffected()
	return affected > 0, nil
}

func DeleteThreshold(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var seqQuery string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	conn = database.Get()
	query = `DELETE FROM public.threshold WHERE "id" = $1`
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete threshold ID %d: %v", id, err)
		return false, err
	}
	var affected int64
	affected = commandTag.RowsAffected()
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.threshold', 'id'), COALESCE(MAX("id"), 1)) FROM public.threshold`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Error("Failed to reset threshold sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}
