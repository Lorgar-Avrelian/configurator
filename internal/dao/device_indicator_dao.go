package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateIndicator(ctx context.Context, d DeviceIndicatorDao) (*DeviceIndicatorDao, error) {
	var conn *pgxpool.Pool
	var query string
	var res DeviceIndicatorDao
	var err error
	conn = database.Get()
	query = `INSERT INTO public.device_indicator ("description", "object_id", "contact", "name", "location", "services")
			 VALUES ($1, $2, $3, $4, $5, $6)
			 ON CONFLICT ("description", "object_id", "contact", "name", "location", "services") 
			 DO UPDATE SET "object_id" = EXCLUDED."object_id"
			 RETURNING "id", "description", "object_id", "contact", "name", "location", "services"`
	err = conn.QueryRow(ctx, query, d.Description, d.ObjectID, d.Contact, d.Name, d.Location, d.Services).Scan(&res.ID, &res.Description, &res.ObjectID, &res.Contact, &res.Name, &res.Location, &res.Services)
	if err != nil {
		logger.Error("Failed to insert device indicator into DB: %v", err)
		return nil, err
	}
	return &res, nil
}

func GetIndicatorByID(ctx context.Context, id int64) (*DeviceIndicatorDao, error) {
	var conn *pgxpool.Pool
	var query string
	var res DeviceIndicatorDao
	var err error
	conn = database.Get()
	query = `SELECT "id", "description", "object_id", "contact", "name", "location", "services"
			 FROM public.device_indicator
			 WHERE "id" = $1`
	err = conn.QueryRow(ctx, query, id).Scan(&res.ID, &res.Description, &res.ObjectID, &res.Contact, &res.Name, &res.Location, &res.Services)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Error("Failed to retrieve device indicator ID %d: %v", id, err)
		return nil, err
	}
	return &res, nil
}

func GetAllIndicators(ctx context.Context) ([]DeviceIndicatorDao, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []DeviceIndicatorDao
	var item DeviceIndicatorDao
	conn = database.Get()
	query = `SELECT "id", "description", "object_id", "contact", "name", "location", "services"
			 FROM public.device_indicator
			 ORDER BY "id" ASC`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch all device indicators from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []DeviceIndicatorDao{}
	for rows.Next() {
		err = rows.Scan(&item.ID, &item.Description, &item.ObjectID, &item.Contact, &item.Name, &item.Location, &item.Services)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func UpdateIndicator(ctx context.Context, id int64, d DeviceIndicatorDao) (*DeviceIndicatorDao, error) {
	var conn *pgxpool.Pool
	var query string
	var res DeviceIndicatorDao
	var err error
	conn = database.Get()
	query = `UPDATE public.device_indicator
			 SET "description" = $1, "object_id" = $2, "contact" = $3, "name" = $4, "location" = $5, "services" = $6
			 WHERE "id" = $7
			 RETURNING "id", "description", "object_id", "contact", "name", "location", "services"`
	err = conn.QueryRow(ctx, query, d.Description, d.ObjectID, d.Contact, d.Name, d.Location, d.Services, id).Scan(&res.ID, &res.Description, &res.ObjectID, &res.Contact, &res.Name, &res.Location, &res.Services)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Error("Failed to update device indicator ID %d: %v", id, err)
		return nil, err
	}
	return &res, nil
}

func DeleteIndicator(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var seqQuery string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	var affected int64
	conn = database.Get()
	query = `DELETE FROM public.device_indicator
			 WHERE "id" = $1`
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete device indicator ID %d: %v", id, err)
		return false, err
	}
	affected = commandTag.RowsAffected()
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.device_indicator', 'id'), COALESCE(MAX("id"), 0) + 1, false)
				FROM public.device_indicator`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Error("Failed to reset device indicator sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}
