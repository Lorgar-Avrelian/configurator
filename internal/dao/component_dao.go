package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetComponentByID(ctx context.Context, id int64) (*Component, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var comp Component
	var err error
	err = fetchComponentBase(ctx, conn, id, &comp)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	var params []Param
	params, err = fetchInheritedParams(ctx, conn, id)
	if err != nil {
		return nil, err
	}
	comp.Params = params
	return &comp, nil
}

func fetchComponentBase(ctx context.Context, conn *pgxpool.Pool, id int64, c *Component) error {
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access" FROM public.component WHERE "id" = $1`
	var err error
	err = conn.QueryRow(ctx, query, id).Scan(&c.ID, &c.Title, &c.NameEn, &c.NameRu, &c.PluralNameEn, &c.PluralNameRu, &c.BaseComponent, &c.DescriptionEn, &c.DescriptionRu, &c.Access)
	return err
}

func CreateComponent(ctx context.Context, d ComponentDao) (*Component, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `INSERT INTO public.component ("title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING "id", "title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access"`
	var comp Component
	var err error
	err = conn.QueryRow(ctx, query, d.Title, d.NameEn, d.NameRu, d.PluralNameEn, d.PluralNameRu, d.BaseComponent, d.DescriptionEn, d.DescriptionRu, d.Access).Scan(&comp.ID, &comp.Title, &comp.NameEn, &comp.NameRu, &comp.PluralNameEn, &comp.PluralNameRu, &comp.BaseComponent, &comp.DescriptionEn, &comp.DescriptionRu, &comp.Access)
	if err != nil {
		logger.Error("Failed to insert component into DB: %v", err)
		return nil, err
	}
	var params []Param
	params, err = fetchInheritedParams(ctx, conn, comp.ID)
	if err != nil {
		return nil, err
	}
	comp.Params = params
	return &comp, nil
}

func UpdateComponent(ctx context.Context, id int64, d ComponentDao) (*Component, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `UPDATE public.component SET "title" = $1, "name_en" = $2, "name_ru" = $3, "plural_name_en" = $4, "plural_name_ru" = $5, "base_component" = $6, "description_en" = $7, "description_ru" = $8, "access" = $9 WHERE "id" = $10 RETURNING "id", "title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access"`
	var comp Component
	var err error
	err = conn.QueryRow(ctx, query, d.Title, d.NameEn, d.NameRu, d.PluralNameEn, d.PluralNameRu, d.BaseComponent, d.DescriptionEn, d.DescriptionRu, d.Access, id).Scan(&comp.ID, &comp.Title, &comp.NameEn, &comp.NameRu, &comp.PluralNameEn, &comp.PluralNameRu, &comp.BaseComponent, &comp.DescriptionEn, &comp.DescriptionRu, &comp.Access)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Error("Failed to update component ID %d: %v", id, err)
		return nil, err
	}
	var params []Param
	params, err = fetchInheritedParams(ctx, conn, comp.ID)
	if err != nil {
		return nil, err
	}
	comp.Params = params
	return &comp, nil
}

func fetchInheritedParams(ctx context.Context, conn *pgxpool.Pool, componentID int64) ([]Param, error) {
	var query string
	query = `WITH RECURSIVE component_tree AS (
		SELECT "id", "base_component" FROM public.component WHERE "id" = $1
		UNION ALL
		SELECT c."id", c."base_component" FROM public.component c
		JOIN component_tree ct ON c."id" = ct."base_component"
	)
	SELECT p."id", p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
	FROM public.param p
	JOIN public.component_param cp ON p."id" = cp."param_id"
	JOIN component_tree ct ON cp."component_id" = ct."id"
	ORDER BY p."id" ASC`
	var rows pgx.Rows
	var err error
	rows = nil
	rows, err = conn.Query(ctx, query, componentID)
	if err != nil {
		logger.Error("Failed to query inherited params: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Param
	list = []Param{}
	var p Param
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.Title, &p.NameEn, &p.NameRu, &p.Type, &p.Value, &p.DescriptionEn, &p.DescriptionRu, &p.UnitsEn, &p.UnitsRu, &p.Access, &p.Saved, &p.Visible)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}

func GetAllComponents(ctx context.Context) ([]Component, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access" FROM public.component ORDER BY "id" ASC`
	var rows pgx.Rows
	var err error
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch all components from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Component
	list = []Component{}
	var c Component
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Title, &c.NameEn, &c.NameRu, &c.PluralNameEn, &c.PluralNameRu, &c.BaseComponent, &c.DescriptionEn, &c.DescriptionRu, &c.Access)
		if err != nil {
			return nil, err
		}
		var params []Param
		params, err = fetchInheritedParams(ctx, conn, c.ID)
		if err != nil {
			return nil, err
		}
		c.Params = params
		list = append(list, c)
	}
	return list, rows.Err()
}

func SearchComponents(ctx context.Context, queryText string) ([]Component, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT "id", "title", "name_en", "name_ru", "plural_name_en", "plural_name_ru", "base_component", "description_en", "description_ru", "access" FROM public.component WHERE "title" ILIKE $1 OR "name_en" ILIKE $1 OR "name_ru" ILIKE $1 OR "description_en" ILIKE $1 OR "description_ru" ILIKE $1 ORDER BY "id" ASC`
	var rows pgx.Rows
	var err error
	var likePattern string
	likePattern = "%" + queryText + "%"
	rows, err = conn.Query(ctx, query, likePattern)
	if err != nil {
		logger.Error("Failed to search components for query '%s': %v", queryText, err)
		return nil, err
	}
	defer rows.Close()
	var list []Component
	list = []Component{}
	var c Component
	for rows.Next() {
		err = rows.Scan(&c.ID, &c.Title, &c.NameEn, &c.NameRu, &c.PluralNameEn, &c.PluralNameRu, &c.BaseComponent, &c.DescriptionEn, &c.DescriptionRu, &c.Access)
		if err != nil {
			return nil, err
		}
		var params []Param
		params, err = fetchInheritedParams(ctx, conn, c.ID)
		if err != nil {
			return nil, err
		}
		c.Params = params
		list = append(list, c)
	}
	return list, rows.Err()
}

func DeleteComponent(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `WITH RECURSIVE component_tree AS (
		SELECT "id" FROM public.component WHERE "id" = $1
		UNION ALL
		SELECT c."id" FROM public.component c
		JOIN component_tree ct ON c."base_component" = ct."id"
	)
	DELETE FROM public.component WHERE "id" IN (SELECT "id" FROM component_tree)`
	var commandTag interface{ RowsAffected() int64 }
	var err error
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete component ID %d: %v", id, err)
		return false, err
	}
	var affected int64
	affected = commandTag.RowsAffected()
	var seqQuery string
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.component', 'id'), COALESCE(MAX("id"), 1)) FROM public.component`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Error("Failed to reset component sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}
