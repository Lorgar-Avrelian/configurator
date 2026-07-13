package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateMappingRaw(ctx context.Context, d MappingDao) (int64, error) {
	var conn *pgxpool.Pool
	var query string
	var insertedID int64
	var err error
	conn = database.Get()
	query = `INSERT INTO public.mapping ("indicator", "param", "frequency", "value", "coefficient", "enum", "position", "from", "position_type")
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			 RETURNING "id"`
	err = conn.QueryRow(ctx, query, d.Indicator, d.Param, d.Frequency, d.Value, d.Coefficient, d.Enum, d.Position, d.From, d.PositionType).Scan(&insertedID)
	if err != nil {
		logger.Error("Failed to raw insert mapping into DB: %v", err)
		return 0, err
	}
	return insertedID, nil
}

func GetMappingByID(ctx context.Context, id int64) ([]Mapping, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Mapping
	var m Mapping
	conn = database.Get()
	query = `WITH RECURSIVE root_search AS (
			     SELECT "id", "from"
			     FROM public.mapping
			     WHERE "id" = $1
			     UNION ALL
			     SELECT m."id", m."from"
			     FROM public.mapping m
			         JOIN root_search rs ON m."id" = rs."from"
			 ),
			 target_root AS (
			     SELECT "id"
			     FROM root_search
			     WHERE "from" IS NULL
			     LIMIT 1
			 ),
			 mapping_branch AS (
			     SELECT "id", "indicator", "param", "frequency", "value", "coefficient", "enum", "position", "from", "position_type"
			     FROM public.mapping
			     WHERE "id" = (SELECT "id" FROM target_root)
			     UNION ALL
			     SELECT m."id", m."indicator", m."param", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type"
			     FROM public.mapping m
			         JOIN mapping_branch mb ON m."from" = mb."id"
			 )
			 SELECT mb."id", mb."indicator", mb."param", mb."frequency", mb."value", mb."coefficient", mb."enum", mb."position", mb."from", mb."position_type",
			        pi."oid_id", pi."dotter_notation", o."mib", mib."path", mib."name", mib."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category",
			        p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
			 FROM mapping_branch mb
			     LEFT JOIN public.param_indicator pi ON mb."indicator" = pi."id"
			     LEFT JOIN public.oid o ON pi."oid_id" = o."id"
			     LEFT JOIN public.mib mib ON o."mib" = mib."id"
			     LEFT JOIN public.param p ON mb."param" = p."id"`
	rows, err = conn.Query(ctx, query, id)
	if err != nil {
		logger.Error("Failed to fetch recursive full tree branch from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Mapping{}
	for rows.Next() {
		err = rows.Scan(&m.ID, &m.IndicatorID, &m.ParamID, &m.Frequency, &m.Value, &m.Coefficient, &m.Enum, &m.Position, &m.From, &m.PositionType, &m.IndOidID, &m.IndDotterNotation, &m.OidMibID, &m.OidMibPath, &m.OidMibName, &m.OidMibVendor, &m.OidType, &m.OidName, &m.OidNumber, &m.OidDotterNotation, &m.OidObjectDescriptor, &m.OidSyntax, &m.OidEnum, &m.OidStatus, &m.OidAccess, &m.OidUnits, &m.OidDescription, &m.OidCategory, &m.ParamTitle, &m.ParamNameEn, &m.ParamNameRu, &m.ParamType, &m.ParamValue, &m.ParamDescriptionEn, &m.ParamDescriptionRu, &m.ParamUnitsEn, &m.ParamUnitsRu, &m.ParamAccess, &m.ParamSaved, &m.ParamVisible)
		if err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, rows.Err()
}

func GetAllMappings(ctx context.Context) ([]Mapping, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Mapping
	var m Mapping
	conn = database.Get()
	query = `WITH RECURSIVE mapping_tree AS (
			     SELECT "id", "indicator", "param", "frequency", "value", "coefficient", "enum", "position", "from", "position_type"
			     FROM public.mapping
			     WHERE "from" IS NULL
			     UNION ALL
			     SELECT m."id", m."indicator", m."param", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type"
			     FROM public.mapping m
			         JOIN mapping_tree mt ON m."from" = mt."id"
			 )
			 SELECT mt."id", mt."indicator", mt."param", mt."frequency", mt."value", mt."coefficient", mt."enum", mt."position", mt."from", mt."position_type",
			        pi."oid_id", pi."dotter_notation", o."mib", mib."path", mib."name", mib."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category",
			        p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
			 FROM mapping_tree mt
			     LEFT JOIN public.param_indicator pi ON mt."indicator" = pi."id"
			     LEFT JOIN public.oid o ON pi."oid_id" = o."id"
			     LEFT JOIN public.mib mib ON o."mib" = mib."id"
			     LEFT JOIN public.param p ON mt."param" = p."id"`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch recursive mappings from DB: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Mapping{}
	for rows.Next() {
		err = rows.Scan(&m.ID, &m.IndicatorID, &m.ParamID, &m.Frequency, &m.Value, &m.Coefficient, &m.Enum, &m.Position, &m.From, &m.PositionType, &m.IndOidID, &m.IndDotterNotation, &m.OidMibID, &m.OidMibPath, &m.OidMibName, &m.OidMibVendor, &m.OidType, &m.OidName, &m.OidNumber, &m.OidDotterNotation, &m.OidObjectDescriptor, &m.OidSyntax, &m.OidEnum, &m.OidStatus, &m.OidAccess, &m.OidUnits, &m.OidDescription, &m.OidCategory, &m.ParamTitle, &m.ParamNameEn, &m.ParamNameRu, &m.ParamType, &m.ParamValue, &m.ParamDescriptionEn, &m.ParamDescriptionRu, &m.ParamUnitsEn, &m.ParamUnitsRu, &m.ParamAccess, &m.ParamSaved, &m.ParamVisible)
		if err != nil {
			return nil, err
		}
		list = append(list, m)
	}
	return list, rows.Err()
}

func UpdateMappingRaw(ctx context.Context, id int64, d MappingDao) error {
	var conn *pgxpool.Pool
	var query string
	var err error
	conn = database.Get()
	query = `UPDATE public.mapping
			 SET "indicator" = $1, "param" = $2, "frequency" = $3, "value" = $4, "coefficient" = $5, "enum" = $6, "position" = $7, "from" = $8, "position_type" = $9
			 WHERE "id" = $10`
	_, err = conn.Exec(ctx, query, d.Indicator, d.Param, d.Frequency, d.Value, d.Coefficient, d.Enum, d.Position, d.From, d.PositionType, id)
	if err != nil {
		logger.Error("Failed to raw update mapping ID %d: %v", id, err)
		return err
	}
	return nil
}

func DeleteMapping(ctx context.Context, id int64) (bool, error) {
	var conn *pgxpool.Pool
	var query string
	var seqQuery string
	var commandTag interface{ RowsAffected() int64 }
	var err error
	var affected int64
	conn = database.Get()
	query = `DELETE FROM public.mapping WHERE "id" = $1`
	commandTag, err = conn.Exec(ctx, query, id)
	if err != nil {
		logger.Error("Failed to delete mapping ID %d: %v", id, err)
		return false, err
	}
	affected = commandTag.RowsAffected()
	seqQuery = `SELECT SETVAL(PG_GET_SERIAL_SEQUENCE('public.mapping', 'id'), COALESCE(MAX("id"), 0) + 1, false) FROM public.mapping`
	_, err = conn.Exec(ctx, seqQuery)
	if err != nil {
		logger.Error("Failed to reset mapping sequence: %v", err)
		return true, err
	}
	return affected > 0, nil
}

func GetMappingByIDOwn(ctx context.Context, id int64) (*Mapping, error) {
	var conn *pgxpool.Pool
	var query string
	var row pgx.Row
	var m Mapping
	var err error
	conn = database.Get()
	query = `SELECT m."id", m."indicator", m."param", m."frequency", m."value", m."coefficient", m."enum", m."position", m."from", m."position_type",
			        pi."oid_id", pi."dotter_notation", o."mib", mib."path", mib."name", mib."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category",
			        p."title", p."name_en", p."name_ru", p."type", p."value", p."description_en", p."description_ru", p."units_en", p."units_ru", p."access", p."saved", p."visible"
			 FROM public.mapping m
			     LEFT JOIN public.param_indicator pi ON m."indicator" = pi."id"
			     LEFT JOIN public.oid o ON pi."oid_id" = o."id"
			     LEFT JOIN public.mib mib ON o."mib" = mib."id"
			     LEFT JOIN public.param p ON m."param" = p."id"
			 WHERE m."id" = $1`
	row = conn.QueryRow(ctx, query, id)
	err = row.Scan(&m.ID, &m.IndicatorID, &m.ParamID, &m.Frequency, &m.Value, &m.Coefficient, &m.Enum, &m.Position, &m.From, &m.PositionType, &m.IndOidID, &m.IndDotterNotation, &m.OidMibID, &m.OidMibPath, &m.OidMibName, &m.OidMibVendor, &m.OidType, &m.OidName, &m.OidNumber, &m.OidDotterNotation, &m.OidObjectDescriptor, &m.OidSyntax, &m.OidEnum, &m.OidStatus, &m.OidAccess, &m.OidUnits, &m.OidDescription, &m.OidCategory, &m.ParamTitle, &m.ParamNameEn, &m.ParamNameRu, &m.ParamType, &m.ParamValue, &m.ParamDescriptionEn, &m.ParamDescriptionRu, &m.ParamUnitsEn, &m.ParamUnitsRu, &m.ParamAccess, &m.ParamSaved, &m.ParamVisible)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		logger.Error("Failed to retrieve single mapping own ID %d: %v", id, err)
		return nil, err
	}
	return &m, nil
}
