package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateDeviceComponentRaw(ctx context.Context, d DeviceComponentDao) (int64, error) {
	var conn *pgxpool.Pool
	var query string
	var insertedID int64
	var err error
	conn = database.Get()
	query = `INSERT INTO public.device_component ("model", "internal_order", "parent")
			 VALUES ($1, $2, $3)
			 RETURNING "id"`
	err = conn.QueryRow(ctx, query, d.Model, d.InternalOrder, d.Parent).Scan(&insertedID)
	if err != nil {
		logger.Error("Failed to raw insert device component into DB: %v", err)
		return 0, err
	}
	return insertedID, nil
}

func GetDeviceComponentByID(ctx context.Context, id int64) ([]DeviceComponent, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []DeviceComponent
	var dc DeviceComponent
	conn = database.Get()
	query = `WITH RECURSIVE root_search AS (
			     SELECT "id", "parent"
			     FROM public.device_component
			     WHERE "id" = $1
			     UNION ALL
			     SELECT m."id", m."parent"
			     FROM public.device_component m
			         JOIN root_search rs ON m."id" = rs."parent"
			 ),
			 target_root AS (
			     SELECT "id"
			     FROM root_search
			     WHERE "parent" IS NULL
			     LIMIT 1
			 ),
			 component_branch AS (
			     SELECT "id", "model", "internal_order", "parent"
			     FROM public.device_component
			     WHERE "id" = (SELECT "id" FROM target_root)
			     UNION ALL
			     SELECT m."id", m."model", m."internal_order", m."parent"
			     FROM public.device_component m
			         JOIN component_branch cb ON m."parent" = cb."id"
			 )
			 SELECT cb."id", cb."model", cb."internal_order", cb."parent",
			        c."title", c."name_en", c."name_ru"
			 FROM component_branch cb
			     LEFT JOIN public.component c ON cb."model" = c."id"`
	rows, err = conn.Query(ctx, query, id)
	if err != nil {
		logger.Error("Failed to fetch recursive device component branch from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []DeviceComponent{}
	for rows.Next() {
		err = rows.Scan(&dc.ID, &dc.ModelID, &dc.InternalOrder, &dc.Parent, &dc.CompTitle, &dc.CompNameEn, &dc.CompNameRu)
		if err != nil {
			return nil, err
		}
		list = append(list, dc)
	}
	return list, rows.Err()
}

func GetAllDeviceComponents(ctx context.Context) ([]DeviceComponent, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []DeviceComponent
	var dc DeviceComponent
	conn = database.Get()
	query = `WITH RECURSIVE component_tree AS (
			     SELECT "id", "model", "internal_order", "parent"
			     FROM public.device_component
			     WHERE "parent" IS NULL
			     UNION ALL
			     SELECT m."id", m."model", m."internal_order", m."parent"
			     FROM public.device_component m
			         JOIN component_tree ct ON m."parent" = ct."id"
			 )
			 SELECT ct."id", ct."model", ct."internal_order", ct."parent",
			        c."title", c."name_en", c."name_ru"
			 FROM component_tree ct
			     LEFT JOIN public.component c ON ct."model" = c."id"`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch recursive device components from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []DeviceComponent{}
	for rows.Next() {
		err = rows.Scan(&dc.ID, &dc.ModelID, &dc.InternalOrder, &dc.Parent, &dc.CompTitle, &dc.CompNameEn, &dc.CompNameRu)
		if err != nil {
			return nil, err
		}
		list = append(list, dc)
	}
	return list, rows.Err()
}

func UpdateDeviceComponentRaw(ctx context.Context, id int64, d DeviceComponentDao) error {
	var conn *pgxpool.Pool
	var query string
	var err error
	conn = database.Get()
	query = `UPDATE public.device_component
			 SET "model" = $1, "internal_order" = $2, "parent" = $3
			 WHERE "id" = $4`
	_, err = conn.Exec(ctx, query, d.Model, d.InternalOrder, d.Parent, id)
	if err != nil {
		logger.Error("Failed to raw update device component ID %d: %v", id, err)
		return err
	}
	return nil
}

func DeleteDeviceComponent(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var seqQuery string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	var affected int64
	conn = database.Get()
	query = `DELETE FROM public.device_component
			 WHERE "id" = $1`
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete device component ID %d: %v", id, err)
		return false, err
	}
	affected = commandTag.RowsAffected()
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.device_component', 'id'), COALESCE(MAX("id"), 1))
				FROM public.device_component`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Error("Failed to reset device component sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}
