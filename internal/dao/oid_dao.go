package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetOidsByPrefixMibVendorPaged(ctx context.Context, notation string, mib string, vendor string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var notationPattern string
	var vendorPattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	notationPattern = notation + "%"
	vendorPattern = "%" + vendor + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" LIKE $1 AND m."name" = $2 AND v."name" LIKE $3
			     UNION
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" LIKE $1 AND m."name" = $2 AND m."path" NOT LIKE '%/%'
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $4 OFFSET (($5 - 1) * $4)`
	rows, err = conn.Query(ctx, query, notationPattern, mib, vendorPattern, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by prefix, mib and vendor: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByExactMibVendorPaged(ctx context.Context, notation string, mib string, vendor string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var vendorPattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	vendorPattern = "%" + vendor + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" = $1 AND m."name" = $2 AND v."name" LIKE $3
			     UNION
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" = $1 AND m."name" = $2 AND m."path" NOT LIKE '%/%'
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $4 OFFSET (($5 - 1) * $4)`
	rows, err = conn.Query(ctx, query, notation, mib, vendorPattern, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by exact notation, mib and vendor: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByPrefixMibPaged(ctx context.Context, notation string, mib string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var pattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	pattern = notation + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         LEFT JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" LIKE $1 AND m."name" = $2
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $3 OFFSET (($4 - 1) * $3)`
	rows, err = conn.Query(ctx, query, pattern, mib, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by prefix and mib: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByExactMibPaged(ctx context.Context, notation string, mib string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         LEFT JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" = $1 AND m."name" = $2
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $3 OFFSET (($4 - 1) * $3)`
	rows, err = conn.Query(ctx, query, notation, mib, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by exact notation and mib: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByPrefixVendorPaged(ctx context.Context, notation string, vendor string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var notationPattern string
	var vendorPattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	notationPattern = notation + "%"
	vendorPattern = "%" + vendor + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" LIKE $1 AND v."name" LIKE $2
			     UNION
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" LIKE $1 AND m."path" NOT LIKE '%/%'
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $3 OFFSET (($4 - 1) * $3)`
	rows, err = conn.Query(ctx, query, notationPattern, vendorPattern, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by prefix and vendor: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByExactVendorPaged(ctx context.Context, notation string, vendor string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var vendorPattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	vendorPattern = "%" + vendor + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" = $1 AND v."name" LIKE $2
			     UNION
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" = $1 AND m."path" NOT LIKE '%/%'
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $3 OFFSET (($4 - 1) * $3)`
	rows, err = conn.Query(ctx, query, notation, vendorPattern, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by exact notation and vendor: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByPrefixOnlyPaged(ctx context.Context, notation string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var pattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	pattern = notation + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         LEFT JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" LIKE $1
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $2 OFFSET (($3 - 1) * $2)`
	rows, err = conn.Query(ctx, query, pattern, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by prefix only: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByExactOnlyPaged(ctx context.Context, notation string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         LEFT JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE o."dotter_notation" = $1
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $2 OFFSET (($3 - 1) * $2)`
	rows, err = conn.Query(ctx, query, notation, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by exact notation only: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByMibVendorPaged(ctx context.Context, mib string, vendor string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var vendorPattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	vendorPattern = "%" + vendor + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE m."name" = $1 AND v."name" LIKE $2
			     UNION
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE m."name" = $1 AND m."path" NOT LIKE '%/%'
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $3 OFFSET (($4 - 1) * $3)`
	rows, err = conn.Query(ctx, query, mib, vendorPattern, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by mib and vendor: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByMibOnlyPaged(ctx context.Context, mib string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         LEFT JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE m."name" = $1
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $2 OFFSET (($3 - 1) * $2)`
	rows, err = conn.Query(ctx, query, mib, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by mib only: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetOidsByVendorOnlyPaged(ctx context.Context, vendor string, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var vendorPattern string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	vendorPattern = "%" + vendor + "%"
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE v."name" LIKE $1
			     UNION
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			     WHERE m."path" NOT LIKE '%/%'
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $2 OFFSET (($3 - 1) * $2)`
	rows, err = conn.Query(ctx, query, vendorPattern, size, page)
	if err != nil {
		logger.Error("Failed to query OIDs by vendor only: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
		err = rows.Scan(&o.ID, &o.MibID, &o.MibPath, &o.MibName, &o.MibVendor, &o.Type, &o.Name, &o.Number, &o.DotterNotation, &o.ObjectDescriptor, &o.Syntax, &o.Enum, &o.Status, &o.Access, &o.Units, &o.Description, &o.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, o)
	}
	return list, rows.Err()
}

func GetAllOidsPaged(ctx context.Context, page int, size int) ([]Oid, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	var list []Oid
	var o Oid
	conn = database.Get()
	query = `SELECT t."id", t."mib", t."path", t."mib_name", t."vendor", t."type", t."name_oid", t."number", t."dotter_notation", t."object_descriptor", t."syntax", t."enum", t."status", t."access", t."units", t."description", t."category"
			 FROM (
			     SELECT o."id", o."mib", m."path", m."name" AS "mib_name", m."vendor", o."type", o."name" AS "name_oid", o."number", o."dotter_notation", o."object_descriptor", o."syntax", o."enum", o."status", o."access", o."units", o."description", o."category", v."name" AS "vendor_name"
			     FROM public.oid o
			         LEFT JOIN public.mib m ON o."mib" = m."id"
			         LEFT JOIN public.vendor v ON m."vendor" = v."id"
			 ) t
			 ORDER BY t."vendor" ASC NULLS FIRST, t."vendor_name" ASC NULLS FIRST, t."mib_name" ASC, string_to_array(trim(leading '.' from t."dotter_notation"), '.')::int[], t."id"
			 LIMIT $1 OFFSET (($2 - 1) * $1)`
	rows, err = conn.Query(ctx, query, size, page)
	if err != nil {
		logger.Error("Failed to query all OIDs paged: %v", err)
		return nil, err
	}
	defer rows.Close()
	list = []Oid{}
	for rows.Next() {
		o = Oid{}
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
		logger.Errorf("Failed to fetch dotter notation by OID ID: %v", err)
		return "", err
	}
	return notation, nil
}
