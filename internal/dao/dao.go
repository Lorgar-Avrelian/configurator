package dao

import (
	"context"
	"filler/internal/database"
	"filler/internal/logger"
	"filler/internal/model"
)

// LoadEnumsFromDB вычитывает все справочники из БД и инициализирует реестр моделей
func LoadEnumsFromDB(ctx context.Context) error {
	conn := database.Get()
	accessMap := make(map[int16]string)
	aRows, err := conn.Query(ctx, `SELECT id, value FROM public.access`)
	if err != nil {
		return err
	}
	defer aRows.Close()
	for aRows.Next() {
		var id int16
		var val string
		if err := aRows.Scan(&id, &val); err != nil {
			return err
		}
		accessMap[id] = val
	}
	varTypeMap := make(map[int16]string)
	vRows, err := conn.Query(ctx, `SELECT id, value FROM public.var_type`)
	if err != nil {
		return err
	}
	defer vRows.Close()
	for vRows.Next() {
		var id int16
		var val string
		if err := vRows.Scan(&id, &val); err != nil {
			return err
		}
		varTypeMap[id] = val
	}
	pollMap := make(map[int16]string)
	pRows, err := conn.Query(ctx, `SELECT id, value FROM public.polling_frequency`)
	if err != nil {
		return err
	}
	defer pRows.Close()
	for pRows.Next() {
		var id int16
		var val string
		if err := pRows.Scan(&id, &val); err != nil {
			return err
		}
		pollMap[id] = val
	}
	model.LoadRegistries(accessMap, varTypeMap, pollMap)
	logger.Info("Все справочники enum успешно загружены из БД в память приложения (%d access, %d types, %d frequencies)", len(accessMap), len(varTypeMap), len(pollMap))
	return nil
}
