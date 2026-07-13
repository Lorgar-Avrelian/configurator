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
	var queryStr string
	var targetStr string
	var err error
	for rows.Next() {
		err = rows.Scan(
			&cfg.ID,
			&cfg.Name,
			&cfg.Description,
			&cfg.Author,
			&cfg.Created,
			&queryStr,
			&targetStr,
			&cfg.Value,
		)
		if err != nil {
			logger.Errorf("Failed to scan threshold row: %v", err)
			return nil, err
		}
		cfg.Query = []byte(queryStr)
		cfg.Target = []byte(targetStr)
		list = append(list, cfg)
	}
	err = rows.Err()
	if err != nil {
		logger.Errorf("Rows iterator encountered an error: %v", err)
		return nil, err
	}
	return list, nil
}

func CreateThreshold(ctx context.Context, t ThresholdDao) (int64, error) {
	var conn *pgxpool.Pool
	var query string
	var insertedID int64
	var err error
	conn = database.Get()
	query = `INSERT INTO public.threshold (
		"name",
		"description",
		"author",
		"created",
		"query",
		"target",
		"value"
	) VALUES ($1, $2, $3, NOW(), $4::jsonb, $5::jsonb, $6)
	RETURNING "id"`
	err = conn.QueryRow(ctx, query, t.Name, t.Description, t.Author, string(t.Query), string(t.Target), t.Value).Scan(&insertedID)
	if err != nil {
		logger.Errorf("Failed to insert threshold into DB: %v", err)
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
	query = `SELECT
		"id",
		"name",
		"description",
		"author",
		"created",
		"query"::text,
		"target"::text,
		"value"
	FROM public.threshold
	WHERE "id" = $1`
	rows, err = conn.Query(ctx, query, id)
	if err != nil {
		logger.Errorf("Failed to query threshold ID %d: %v", id, err)
		return nil, err
	}
	defer rows.Close()
	var res []ThresholdDao
	res, err = scanThresholdRows(rows)
	if err != nil {
		logger.Errorf("Failed to scan rows for threshold ID %d: %v", id, err)
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
	query = `SELECT
		"id",
		"name",
		"description",
		"author",
		"created",
		"query"::text,
		"target"::text,
		"value"
	FROM public.threshold
	ORDER BY "id" ASC`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Errorf("Failed to fetch all thresholds from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res []ThresholdDao
	res, err = scanThresholdRows(rows)
	if err != nil {
		logger.Errorf("Failed to scan all thresholds rows: %v", err)
		return nil, err
	}
	return res, nil
}

func UpdateThreshold(ctx context.Context, id int64, t ThresholdDao) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	conn = database.Get()
	query = `UPDATE public.threshold
	SET
		"name" = $1,
		"description" = $2,
		"author" = $3,
		"query" = $4::jsonb,
		"target" = $5::jsonb,
		"value" = $6
	WHERE "id" = $7`
	commandTag, err = conn.Exec(ctx, query, t.Name, t.Description, t.Author, string(t.Query), string(t.Target), t.Value, id)
	if err != nil {
		logger.Errorf("Failed to update threshold ID %d: %v", id, err)
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
	query = `DELETE FROM public.threshold
	WHERE "id" = $1`
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Errorf("Failed to delete threshold ID %d: %v", id, err)
		return false, err
	}
	var affected int64
	affected = commandTag.RowsAffected()
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.threshold', 'id'), COALESCE(MAX("id"), 0) + 1, false) FROM public.threshold`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Errorf("Failed to reset threshold sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}
