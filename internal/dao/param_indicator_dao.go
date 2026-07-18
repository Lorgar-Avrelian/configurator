package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetParamIndicatorByFields(ctx context.Context, oidID uuid.UUID, notation string) (*ParamIndicator, error) {
	var conn *pgxpool.Pool
	var query string
	var row pgx.Row
	var p ParamIndicator
	var err error
	conn = database.Get()
	if oidID == uuid.Nil {
		query = `SELECT pi."id", pi."oid_id", pi."dotter_notation",
				        o."mib", m."path", m."name", m."vendor",
				        o."type", o."name", o."number", o."dotter_notation", o."object_descriptor",
				        o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
				 FROM public.param_indicator pi
				     LEFT JOIN public.oid o ON pi."oid_id" = o."id"
				     LEFT JOIN public.mib m ON o."mib" = m."id"
				 WHERE pi."oid_id" IS NULL AND pi."dotter_notation" = $1`
		row = conn.QueryRow(ctx, query, notation)
	} else {
		query = `SELECT pi."id", pi."oid_id", pi."dotter_notation",
				        o."mib", m."path", m."name", m."vendor",
				        o."type", o."name", o."number", o."dotter_notation", o."object_descriptor",
				        o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
				 FROM public.param_indicator pi
				     LEFT JOIN public.oid o ON pi."oid_id" = o."id"
				     LEFT JOIN public.mib m ON o."mib" = m."id"
				 WHERE pi."oid_id" = $1::uuid AND pi."dotter_notation" = $2`
		row = conn.QueryRow(ctx, query, oidID, notation)
	}
	err = row.Scan(&p.ID, &p.OidID, &p.DotterNotation, &p.OidMibID, &p.OidMibPath, &p.OidMibName, &p.OidMibVendor, &p.OidType, &p.OidName, &p.OidNumber, &p.OidDotterNotation, &p.OidObjectDescriptor, &p.OidSyntax, &p.OidEnum, &p.OidStatus, &p.OidAccess, &p.OidUnits, &p.OidDescription, &p.OidCategory)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Errorf("Failed to retrieve param indicator by fields: %v", err)
		return nil, err
	}
	return &p, nil
}

func CreateParamIndicator(ctx context.Context, d ParamIndicatorDao) (*ParamIndicator, error) {
	var conn *pgxpool.Pool
	var query string
	var insertedID int64
	var err error
	var res *ParamIndicator
	conn = database.Get()
	if d.OidID == uuid.Nil {
		query = `INSERT INTO public.param_indicator ("oid_id", "dotter_notation")
			 VALUES (NULL, $1)
			 RETURNING "id"`
		err = conn.QueryRow(ctx, query, d.DotterNotation).Scan(&insertedID)
	} else {
		query = `INSERT INTO public.param_indicator ("oid_id", "dotter_notation")
			 VALUES ($1::uuid, $2)
			 RETURNING "id"`
		err = conn.QueryRow(ctx, query, d.OidID, d.DotterNotation).Scan(&insertedID)
	}
	if err != nil {
		logger.Errorf("Failed to insert param indicator into DB: %v", err)
		return nil, err
	}
	res, err = GetParamIndicatorByID(ctx, insertedID)
	return res, err
}

func GetParamIndicatorByID(ctx context.Context, id int64) (*ParamIndicator, error) {
	var conn *pgxpool.Pool
	var query string
	var row pgx.Row
	var p ParamIndicator
	var err error
	conn = database.Get()
	query = `SELECT pi."id", pi."oid_id", pi."dotter_notation",
			        o."mib", m."path", m."name", m."vendor",
			        o."type", o."name", o."number", o."dotter_notation", o."object_descriptor",
			        o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
			 FROM public.param_indicator pi
			     LEFT JOIN public.oid o ON pi."oid_id" = o."id"
			     LEFT JOIN public.mib m ON o."mib" = m."id"
			 WHERE pi."id" = $1`
	row = conn.QueryRow(ctx, query, id)
	err = row.Scan(&p.ID, &p.OidID, &p.DotterNotation, &p.OidMibID, &p.OidMibPath, &p.OidMibName, &p.OidMibVendor, &p.OidType, &p.OidName, &p.OidNumber, &p.OidDotterNotation, &p.OidObjectDescriptor, &p.OidSyntax, &p.OidEnum, &p.OidStatus, &p.OidAccess, &p.OidUnits, &p.OidDescription, &p.OidCategory)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Errorf("Failed to retrieve param indicator ID %d: %v", id, err)
		return nil, err
	}
	return &p, nil
}

func GetAllParamIndicators(ctx context.Context) ([]ParamIndicator, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []ParamIndicator
	var p ParamIndicator
	conn = database.Get()
	query = `SELECT pi."id", pi."oid_id", pi."dotter_notation",
			        o."mib", m."path", m."name", m."vendor",
			        o."type", o."name", o."number", o."dotter_notation", o."object_descriptor",
			        o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
			 FROM public.param_indicator pi
			     LEFT JOIN public.oid o ON pi."oid_id" = o."id"
			     LEFT JOIN public.mib m ON o."mib" = m."id"
			 ORDER BY pi."id" ASC`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Errorf("Failed to fetch all param indicators from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []ParamIndicator{}
	for rows.Next() {
		err = rows.Scan(&p.ID, &p.OidID, &p.DotterNotation, &p.OidMibID, &p.OidMibPath, &p.OidMibName, &p.OidMibVendor, &p.OidType, &p.OidName, &p.OidNumber, &p.OidDotterNotation, &p.OidObjectDescriptor, &p.OidSyntax, &p.OidEnum, &p.OidStatus, &p.OidAccess, &p.OidUnits, &p.OidDescription, &p.OidCategory)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
	}
	return list, rows.Err()
}

func UpdateParamIndicator(ctx context.Context, id int64, d ParamIndicatorDao) (*ParamIndicator, error) {
	var conn *pgxpool.Pool
	var query string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	var res *ParamIndicator
	conn = database.Get()
	if d.OidID == uuid.Nil {
		query = `UPDATE public.param_indicator
			 SET "oid_id" = NULL, "dotter_notation" = $1
			 WHERE "id" = $2`
		commandTag, err = conn.Exec(ctx, query, d.DotterNotation, id)
	} else {
		query = `UPDATE public.param_indicator
			 SET "oid_id" = NULLIF($1, '00000000-0000-0000-0000-000000000000'::uuid), "dotter_notation" = $2
			 WHERE "id" = $3`
		commandTag, err = conn.Exec(ctx, query, d.OidID, d.DotterNotation, id)
	}
	if err != nil {
		logger.Errorf("Failed to update param indicator ID %d: %v", id, err)
		return nil, err
	}
	if commandTag.RowsAffected() == 0 {
		return nil, nil
	}
	res, err = GetParamIndicatorByID(ctx, id)
	return res, err
}

func DeleteParamIndicator(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var seqQuery string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	var affected int64
	conn = database.Get()
	query = `DELETE FROM public.param_indicator
			 WHERE "id" = $1`
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Errorf("Failed to delete param indicator ID %d: %v", id, err)
		return false, err
	}
	affected = commandTag.RowsAffected()
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.param_indicator', 'id'), COALESCE(MAX("id"), 0) + 1, false)
				FROM public.param_indicator`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Errorf("Failed to reset param indicator sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}
