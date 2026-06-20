package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func BindParam(ctx context.Context, componentID int64, paramID int64) error {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `INSERT INTO public.component_param ("component_id", "param_id")
		     VALUES ($1, $2)
		     ON CONFLICT ("component_id", "param_id") DO NOTHING`
	var err error
	_, err = conn.Exec(ctx, query, componentID, paramID)
	if err != nil {
		logger.Error("Failed to bind parameter ID %d to component ID %d: %v", paramID, componentID, err)
		return err
	}
	return nil
}

func UnbindParam(ctx context.Context, componentID int64, paramID int64) (bool, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `DELETE FROM public.component_param
       		 WHERE "component_id" = $1 AND "param_id" = $2`
	var commandTag interface{ RowsAffected() int64 }
	var err error
	commandTag, err = conn.Exec(ctx, query, componentID, paramID)
	if err != nil {
		logger.Error("Failed to unbind parameter ID %d from component ID %d: %v", paramID, componentID, err)
		return false, err
	}
	var affected int64
	affected = commandTag.RowsAffected()
	return affected > 0, nil
}

func GetComponentsByDirectParamID(ctx context.Context, paramID int64) ([]ComponentDao, error) {
	var conn *pgxpool.Pool
	conn = database.Get()
	var query string
	query = `SELECT c."id", c."title", c."name_en", c."name_ru", c."plural_name_en", c."plural_name_ru", c."base_component", c."description_en", c."description_ru", c."access" FROM public.component c JOIN public.component_param cp ON c."id" = cp."component_id" WHERE cp."param_id" = $1 ORDER BY c."id" ASC`
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
