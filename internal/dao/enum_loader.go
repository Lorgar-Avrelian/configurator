package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"configurator/internal/model"
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LoadEnumsFromDB(ctx context.Context) error {
	var conn *pgxpool.Pool
	var accessMap map[int16]string
	var alarmMap map[int16]string
	var asn1Map map[int16]string
	var logicOperators []map[string]interface{}
	var oidAccessMap map[int16]string
	var statusMap map[int16]string
	var pollMap map[int16]string
	var varTypeMap map[int16]string
	var oidTypeMap map[int16]string
	var pollingProtocolMap map[int16]string
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
	logicOperators, err = fetchLogicOperators(ctx, conn)
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
	pollingProtocolMap, err = fetchSimpleEnum(ctx, conn, "public.polling_protocol")
	if err != nil {
		return err
	}
	model.LoadRegistries(accessMap, varTypeMap, pollMap, asn1Map, statusMap, oidAccessMap, logicOperators, alarmMap, vendors, oidTypeMap, pollingProtocolMap)
	var logMsg string
	logMsg = fmt.Sprintf("\n%-17s | %s\n------------------+-------\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d\n%-17s | %d",
		"Registry Name", "Count",
		"Access", len(accessMap),
		"Alarm Level", len(alarmMap),
		"ASN.1 Type", len(asn1Map),
		"Logic Operator", len(logicOperators),
		"OID Access", len(oidAccessMap),
		"OID Status", len(statusMap),
		"OID Type", len(oidTypeMap),
		"Polling Frequency", len(pollMap),
		"Polling Protocol", len(pollingProtocolMap),
		"Variable Type", len(varTypeMap),
		"Vendor", len(vendors),
	)
	logger.Info("System registries successfully loaded from DB:%s", logMsg)
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

func fetchLogicOperators(ctx context.Context, conn *pgxpool.Pool) ([]map[string]interface{}, error) {
	var query string
	query = `SELECT "id", "value", "type", "precedence", "arity" FROM public.logic_operator`
	var rows pgx.Rows
	var err error
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch logic operators from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []map[string]interface{}
	list = []map[string]interface{}{}
	var id int16
	var value string
	var opType string
	var precedence int16
	var arity int16
	for rows.Next() {
		err = rows.Scan(&id, &value, &opType, &precedence, &arity)
		if err != nil {
			return nil, err
		}
		var item map[string]interface{}
		item = map[string]interface{}{
			"id":         id,
			"value":      value,
			"type":       opType,
			"precedence": precedence,
			"arity":      arity,
		}
		list = append(list, item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return list, nil
}
