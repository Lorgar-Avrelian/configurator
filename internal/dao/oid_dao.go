package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetOidsByExactDotter(ctx context.Context, notation string) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	query = `SELECT o."id", o."mib", m."path", m."name", m."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
			 FROM public.oid o
			     LEFT JOIN public.mib m ON o."mib" = m."id" WHERE o."dotter_notation" = $1
			 ORDER BY string_to_array(trim(leading '.' from o."dotter_notation"), '.')::int[]`
	rows, err = conn.Query(ctx, query, notation)
	if err != nil {
		logger.Error("Failed to query exact OIDs: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByDotterPrefixPaged(ctx context.Context, prefix string, page int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var pattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	pattern = prefix + "%"
	query = `SELECT o."id", o."mib", m."path", m."name", m."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
			 FROM public.oid o
			     LEFT JOIN public.mib m ON o."mib" = m."id"
			 WHERE o."dotter_notation" LIKE $1
			 ORDER BY string_to_array(trim(leading '.' from o."dotter_notation"), '.')::int[] LIMIT 100 OFFSET (($2 - 1) * 100)`
	rows, err = conn.Query(ctx, query, pattern, page)
	if err != nil {
		logger.Error("Failed to query paged OIDs by prefix: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByMibName(ctx context.Context, name string) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	query = `SELECT o."id", o."mib", m."path", m."name", m."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
			 FROM public.oid o
			     JOIN public.mib m ON o."mib" = m."id"
			 WHERE m."name" = $1
			 ORDER BY string_to_array(trim(leading '.' from o."dotter_notation"), '.')::int[]`
	rows, err = conn.Query(ctx, query, name)
	if err != nil {
		logger.Error("Failed to query OIDs by MIB name: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByVendorIDPaged(ctx context.Context, vendorID int64, filterByVendor bool, page int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	if filterByVendor {
		query = `SELECT o."id", o."mib", m."path", m."name", m."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
				 FROM public.oid o
				     JOIN public.mib m ON o."mib" = m."id"
				 WHERE m."vendor" = $1
				 ORDER BY string_to_array(trim(leading '.' from o."dotter_notation"), '.')::int[] LIMIT 100 OFFSET (($2 - 1) * 100)`
		rows, err = conn.Query(ctx, query, vendorID, page)
	} else {
		query = `SELECT o."id", o."mib", m."path", m."name", m."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
				 FROM public.oid o
				     LEFT JOIN public.mib m ON o."mib" = m."id"
				 ORDER BY string_to_array(trim(leading '.' from o."dotter_notation"), '.')::int[] LIMIT 100 OFFSET (($1 - 1) * 100)`
		rows, err = conn.Query(ctx, query, page)
	}
	if err != nil {
		logger.Error("Failed to query paged OIDs by vendor: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByDotterMibAndVendor(ctx context.Context, notation string, mibName string, vendorPtr *string) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	if vendorPtr != nil {
		query = `SELECT o."id", o."mib", m."path", m."name", m."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
				 FROM public.oid o
				     JOIN public.mib m ON o."mib" = m."id"
				     JOIN public.vendor v ON m."vendor" = v."id"
				 WHERE o."dotter_notation" = $1 AND m."name" = $2 AND (v."name" = $3 OR v."directory" = $3)
				 ORDER BY string_to_array(trim(leading '.' from o."dotter_notation"), '.')::int[]`
		rows, err = conn.Query(ctx, query, notation, mibName, *vendorPtr)
	} else {
		query = `SELECT o."id", o."mib", m."path", m."name", m."vendor", o."type", o."name", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category"
				 FROM public.oid o
				     JOIN public.mib m ON o."mib" = m."id"
				 WHERE o."dotter_notation" = $1 AND m."name" = $2 AND m."vendor" IS NULL
				 ORDER BY string_to_array(trim(leading '.' from o."dotter_notation"), '.')::int[]`
		rows, err = conn.Query(ctx, query, notation, mibName)
	}
	if err != nil {
		logger.Error("Failed to query OIDs by dotter, MIB and vendor: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetDotterNotationByOidID(ctx context.Context, id uuid.UUID) (string, error) {
	var conn *pgxpool.Pool
	var query string
	var notation string
	var err error
	conn = database.Get()
	query = `SELECT "dotter_notation"
			 FROM public.oid
			 WHERE "id" = $1`
	err = conn.QueryRow(ctx, query, id).Scan(&notation)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", nil
		}
		logger.Error("Failed to fetch dotter notation by OID ID: %v", err)
		return "", err
	}
	return notation, nil
}
