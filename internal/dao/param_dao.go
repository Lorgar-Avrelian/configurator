package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateParam(ctx context.Context, d ParamDao) (*ParamDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `INSERT INTO public.param ("title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible", "diagram")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
RETURNING "id", "title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible", "diagram"`
	var p ParamDao
	var err error
	err = conn.QueryRow(ctx, query, d.Title, d.NameEn, d.NameRu, d.Type, d.Value, d.DescriptionEn, d.DescriptionRu, d.UnitsEn, d.UnitsRu, d.Access, d.Saved, d.Visible, d.Diagram).Scan(&p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible, &p.Diagram)
	if err != nil {
		logger.Error("Failed to insert parameter into DB: %v", err)
		return nil, err
	}
	return &p, nil
}

func GetParamByID(ctx context.Context, id int64) (*ParamDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible", "diagram"
FROM public.param
WHERE "id" = $1`
	var p ParamDao
	var err error
	err = conn.QueryRow(ctx, query, id).Scan(&p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible, &p.Diagram)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Error("Failed to retrieve parameter ID %d: %v", id, err)
		return nil, err
	}
	return &p, nil
}

func UpdateParam(ctx context.Context, id int64, d ParamDao) (*ParamDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `UPDATE public.param
SET "title" = $1, "name_en" = $2, "name_ru" = $3, "type" = $4, "value" = $5, "description_en" = $6, "description_ru" = $7, "units_en" = $8, "units_ru" = $9, "access" = $10, "saved" = $11, "visible" = $12, "diagram" = $13
WHERE "id" = $14
RETURNING "id", "title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible", "diagram"`
	var p ParamDao
	var err error
	err = conn.QueryRow(ctx, query, d.Title, d.NameEn, d.NameRu, d.Type, d.Value, d.DescriptionEn, d.DescriptionRu, d.UnitsEn, d.UnitsRu, d.Access, d.Saved, d.Visible, d.Diagram, id).Scan(&p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Error("Failed to update parameter ID %d: %v", id, err)
		return nil, err
	}
	return &p, nil
}

func DeleteParam(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `DELETE FROM public.param WHERE "id" = $1`
	var commandTag interface{ RowsAffected() int64 }
	var err error
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete parameter ID %d: %v", id, err)
		return false, err
	}
	var affected int64
	affected = commandTag.RowsAffected()
	var seqQuery string
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.param', 'id'), COALESCE(MAX("id"), 0) + 1, false) FROM public.param`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Error("Failed to reset param sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}

func GetAllParams(ctx context.Context) ([]ParamDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible", "diagram"
FROM public.param
ORDER BY "id" ASC`
	var rows pgx.Rows
	var err error
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch all parameters from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []ParamDao
	list = []ParamDao{}
	var p ParamDao
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible, &p.Diagram)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}

func GetUnattachedParams(ctx context.Context) ([]ParamDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT p."id", p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible", p."diagram"
FROM public.param p
    LEFT JOIN public.component_param cp ON p."id" = cp."param_id"
WHERE cp."component_id" IS NULL
ORDER BY p."id" ASC`
	var rows pgx.Rows
	var err error
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch unattached parameters from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []ParamDao
	list = []ParamDao{}
	var p ParamDao
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible, &p.Diagram)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}

func SearchParams(ctx context.Context, queryText string) ([]ParamDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "type", "value", "description_en", "description_ru", "units_en", "units_ru", "access", "saved", "visible", "diagram"
FROM public.param
WHERE "title" ILIKE $1 OR "name_en" ILIKE $1 OR "name_ru" ILIKE $1 OR "description_en" ILIKE $1 OR "description_ru" ILIKE $1
ORDER BY "id" ASC`
	var rows pgx.Rows
	var err error
	var likePattern string
	likePattern = "%" + queryText + "%"
	rows, err = conn.Query(ctx, query, likePattern)
	if err != nil {
		logger.Error("Failed to search parameters for query '%s': %v", queryText, err)
		return nil, err
	}
	defer rows.Close()
	var list []ParamDao
	list = []ParamDao{}
	var p ParamDao
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible, &p.Diagram)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}

func GetComponentsByDirectParamID(ctx context.Context, paramID int64) ([]ComponentDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT c."id", c."title", c."name_en", c."name_ru", c."plural_name_en", c."plural_name_ru", c."base_component", c."description_en", c."description_ru", c."access"
FROM public.component c
    JOIN public.component_param cp ON c."id" = cp."component_id"
WHERE cp."param_id" = $1 ORDER BY c."id" ASC`
	var rows pgx.Rows
	var err error
	rows, err = conn.Query(ctx, query, paramID)
	if err != nil {
		logger.Error("Failed to fetch components by direct param ID %d: %v", paramID, err)
		return nil, err
	}
	defer rows.Close()
	var list []ComponentDao
	list = []ComponentDao{}
	var c ComponentDao
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Title, &c.NameEn, &c.NameRu, &c.PluralNameEn, &c.PluralNameRu, &c.BaseComponent, &c.DescriptionEn, &c.DescriptionRu, &c.Access)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, rows.Err()
}
