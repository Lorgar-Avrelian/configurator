package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDeviceSnmpByHostPort(ctx context.Context, host string, port int32) (*DeviceSnmpDao, error) {
	var conn *pgxpool.Pool
	var query string
	var d DeviceSnmpDao
	var err error
	conn = database.Get()
	query = `SELECT 
			    "host", 
			    "port", 
			    "community", 
			    "version", 
			    "login", 
			    "password", 
			    "authentication", 
			    "privacy", 
			    "id", 
			    "config" 
			 FROM public.device_snmp 
			 WHERE "host" = $1 
			   AND "port" = $2`
	err = conn.QueryRow(ctx, query, host, port).Scan(&d.Host, &d.Port, &d.Community, &d.Version, &d.Login, &d.Password, &d.Authentication, &d.Privacy, &d.ID, &d.Config)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Error("Failed to fetch device_snmp by host and port: %v", err)
		return nil, err
	}
	return &d, nil
}

func GetThresholdsByHostPort(ctx context.Context, host string, port int32) ([]ThresholdDao, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT t."id", 
                    t."name", 
                    t."description", 
                    t."author", 
                    t."created", 
                    t."query" 
             FROM public.threshold t 
                 JOIN public.affected_threshold af ON t."id" = af."threshold" 
             WHERE af."host" = $1 AND af."port" = $2 
             ORDER BY t."id" ASC`
	rows, err = conn.Query(ctx, query, host, port)
	if err != nil {
		logger.Error("Failed to fetch thresholds for device: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []ThresholdDao
	list = []ThresholdDao{}
	var t ThresholdDao
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.Name, &t.Description, &t.Author, &t.Created, &t.Query)
		if err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, rows.Err()
}

func GetComponentsWithTargetParamByHostPort(ctx context.Context, host string, port int32) ([]Component, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    c."id", 
			    c."title", 
			    c."name_en", 
			    c."name_ru", 
			    c."plural_name_en", 
			    c."plural_name_ru", 
			    c."base_component", 
			    c."description_en", 
			    c."description_ru", 
			    c."access", 
			    p."id", 
			    p."title", 
			    p."name_en", 
			    p."name_ru", 
			    p."type", 
			    p."value", 
			    p."description_en", 
			    p."description_ru", 
			    p."units_en", 
			    p."units_ru", 
			    p."access", 
			    p."saved", 
			    p."visible" 
			 FROM public.affected_param ap 
			 JOIN public.component c 
			     ON ap."component" = c."id" 
			 JOIN public.param p 
			     ON ap."param" = p."id" 
			 WHERE ap."host" = $1 
			   AND ap."port" = $2 
			 ORDER BY c."id" ASC, p."id" ASC`
	rows, err = conn.Query(ctx, query, host, port)
	if err != nil {
		logger.Error("Failed to fetch components with target params: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Component
	list = []Component{}
	var c Component
	var p Param
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Title, &c.NameEn, &c.NameRu, &c.PluralNameEn, &c.PluralNameRu, &c.BaseComponent, &c.DescriptionEn, &c.DescriptionRu, &c.Access, &p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible)
		if err != nil {
			return nil, err
		}
		c.Params = []Param{p}
		list = append(list, c)
	}
	return list, rows.Err()
}
