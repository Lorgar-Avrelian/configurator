package dao

import (
	"context"
	"database/sql"
	"filler/internal/database"
	"filler/internal/logger"
	"filler/internal/model"
)

// LoadEnumsFromDB вычитывает все справочники из БД и инициализирует реестр моделей
func LoadEnumsFromDB(ctx context.Context) error {
	conn := database.Get()
	accessMap := make(map[int16]string)
	aRows, _ := conn.Query(ctx, `SELECT id, value FROM public.access`)
	defer aRows.Close()
	for aRows.Next() {
		var id int16
		var val string
		_ = aRows.Scan(&id, &val)
		accessMap[id] = val
	}
	varTypeMap := make(map[int16]string)
	vRows, _ := conn.Query(ctx, `SELECT id, value FROM public.var_type`)
	defer vRows.Close()
	for vRows.Next() {
		var id int16
		var val string
		_ = vRows.Scan(&id, &val)
		varTypeMap[id] = val
	}
	pollMap := make(map[int16]string)
	pRows, _ := conn.Query(ctx, `SELECT id, value FROM public.polling_frequency`)
	defer pRows.Close()
	for pRows.Next() {
		var id int16
		var val string
		_ = pRows.Scan(&id, &val)
		pollMap[id] = val
	}
	model.LoadRegistries(accessMap, varTypeMap, pollMap)
	asn1Map := make(map[int16]string)
	asRows, err := conn.Query(ctx, `SELECT id, value FROM public.asn1_type`)
	if err == nil {
		defer asRows.Close()
		for asRows.Next() {
			var id int16
			var val string
			_ = asRows.Scan(&id, &val)
			asn1Map[id] = val
		}
	}
	statusMap := make(map[int16]string)
	stRows, err := conn.Query(ctx, `SELECT id, value FROM public.oid_status`)
	if err == nil {
		defer stRows.Close()
		for stRows.Next() {
			var id int16
			var val string
			_ = stRows.Scan(&id, &val)
			statusMap[id] = val
		}
	}
	oidAccessMap := make(map[int16]string)
	oaRows, err := conn.Query(ctx, `SELECT id, value FROM public.oid_access`)
	if err == nil {
		defer oaRows.Close()
		for oaRows.Next() {
			var id int16
			var val string
			_ = oaRows.Scan(&id, &val)
			oidAccessMap[id] = val
		}
	}
	var vendors []map[string]interface{}
	vdRows, err := conn.Query(ctx, `SELECT id, name, directory FROM public.vendor`)
	if err == nil {
		defer vdRows.Close()
		for vdRows.Next() {
			var id int64
			var name string
			var dir sql.NullString
			if err := vdRows.Scan(&id, &name, &dir); err == nil {
				vendors = append(vendors, map[string]interface{}{
					"id":        id,
					"name":      name,
					"directory": dir.String,
				})
			}
		}
	}
	model.LoadOidRegistries(asn1Map, statusMap, oidAccessMap, vendors)
	logger.Info("Все справочники OID успешно загружены из БД (%d asn1, %d statuses, %d access, %d vendors)", len(asn1Map), len(statusMap), len(oidAccessMap), len(vendors))
	return nil
}
