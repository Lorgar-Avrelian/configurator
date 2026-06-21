package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"configurator/internal/model"
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadEnumsFromDB(ctx context.Context) error {
	var conn *pgxpool.Pool
	var accessMap map[int16]string
	var alarmMap map[int16]string
	var asn1Map map[int16]string
	var logicMap map[int16]string
	var oidAccessMap map[int16]string
	var statusMap map[int16]string
	var pollMap map[int16]string
	var varTypeMap map[int16]string
	var oidTypeMap map[int16]string
	var vendors []map[string]interface{}
	var err error
	conn = database.Get()
	accessMap, err = fetchSimpleEnum(ctx, conn, "public.access")
	if err != nil {
		return err
	}
	alarmMap, err = fetchSimpleEnum(ctx, conn, "public.alarm_level")
	if err != nil {
		return err
	}
	asn1Map, err = fetchSimpleEnum(ctx, conn, "public.asn1_type")
	if err != nil {
		return err
	}
	logicMap, err = fetchSimpleEnum(ctx, conn, "public.logic_operator")
	if err != nil {
		return err
	}
	oidAccessMap, err = fetchSimpleEnum(ctx, conn, "public.oid_access")
	if err != nil {
		return err
	}
	statusMap, err = fetchSimpleEnum(ctx, conn, "public.oid_status")
	if err != nil {
		return err
	}
	pollMap, err = fetchSimpleEnum(ctx, conn, "public.polling_frequency")
	if err != nil {
		return err
	}
	varTypeMap, err = fetchSimpleEnum(ctx, conn, "public.var_type")
	if err != nil {
		return err
	}
	vendors, err = fetchVendorsList(ctx, conn)
	if err != nil {
		return err
	}
	oidTypeMap, err = fetchSimpleEnum(ctx, conn, "public.oid_type")
	if err != nil {
		return err
	}
	model.LoadRegistries(accessMap, varTypeMap, pollMap, asn1Map, statusMap, oidAccessMap, logicMap, alarmMap, vendors, oidTypeMap)
	logger.Info("System registries successfully loaded from DB: access=%d, alarms=%d, asn1=%d, logic=%d, oid_access=%d, status=%d, frequencies=%d, types=%d, vendors=%d, oid_types=%d", len(accessMap), len(alarmMap), len(asn1Map), len(logicMap), len(oidAccessMap), len(statusMap), len(pollMap), len(varTypeMap), len(vendors), len(oidTypeMap))
	return nil
}

func fetchSimpleEnum(ctx context.Context, conn *pgxpool.Pool, table string) (map[int16]string, error) {
	var query string
	query = `SELECT "id", "value" FROM ` + table
	var rows pgx.Rows
	var err error
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch enum from table %s: %v", table, err)
		return nil, err
	}
	defer rows.Close()
	var m map[int16]string
	m = make(map[int16]string)
	var id int16
	var val string
	for rows.Next() {
		err = rows.Scan(&id, &val)
		if err != nil {
			return nil, err
		}
		m[id] = val
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func fetchVendorsList(ctx context.Context, conn *pgxpool.Pool) ([]map[string]interface{}, error) {
	var query string
	query = `SELECT "id", "name", "number", "contact", "email", "directory" FROM public.vendor`
	var rows pgx.Rows
	var err error
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch vendors from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []map[string]interface{}
	list = []map[string]interface{}{}
	var id int64
	var name string
	var num int32
	var dir sql.NullString
	var email sql.NullString
	var contact sql.NullString
	for rows.Next() {
		err = rows.Scan(&id, &name, &num, &contact, &email, &dir)
		if err != nil {
			return nil, err
		}
		var item map[string]interface{}
		item = map[string]interface{}{
			"id":        id,
			"name":      name,
			"number":    num,
			"contact":   contact,
			"email":     email,
			"directory": dir,
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return list, nil
}
