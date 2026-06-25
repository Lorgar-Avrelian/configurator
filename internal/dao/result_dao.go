package dao

import (
	"configurator/internal/database"
	"configurator/internal/logger"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetAllParamResults(ctx context.Context) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 ORDER BY r."host" ASC, r."port" ASC, r."component" ASC, r."internal_order" ASC`
	rows, err = conn.Query(ctx, query)
	if err != nil {
		logger.Error("Failed to fetch all param results: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByParamTitle(ctx context.Context, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE p."title" = $1 
			 ORDER BY r."host" ASC, r."port" ASC, r."component" ASC, r."internal_order" ASC`
	rows, err = conn.Query(ctx, query, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByInternalOrder(ctx context.Context, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."internal_order" = $1 
			 ORDER BY r."host" ASC, r."port" ASC, r."component" ASC`
	rows, err = conn.Query(ctx, query, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByInternalOrderAndParamTitle(ctx context.Context, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."internal_order" = $1 
			   AND p."title" = $2 
			 ORDER BY r."host" ASC, r."port" ASC, r."component" ASC`
	rows, err = conn.Query(ctx, query, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByComponentTitle(ctx context.Context, compTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE c."title" = $1 
			 ORDER BY r."host" ASC, r."port" ASC, r."internal_order" ASC`
	rows, err = conn.Query(ctx, query, compTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by component title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByComponentTitleAndParamTitle(ctx context.Context, compTitle string, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE c."title" = $1 
			   AND p."title" = $2 
			 ORDER BY r."host" ASC, r."port" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, compTitle, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by component title and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByComponentTitleAndInternalOrder(ctx context.Context, compTitle string, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE c."title" = $1 
			   AND r."internal_order" = $2 
			 ORDER BY r."host" ASC, r."port" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, compTitle, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by component title and internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByComponentTitleAndInternalOrderAndParamTitle(ctx context.Context, compTitle string, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE c."title" = $1 
			   AND r."internal_order" = $2 
			   AND p."title" = $3 
			 ORDER BY r."host" ASC, r."port" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, compTitle, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by component title, internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPort(ctx context.Context, port int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			 ORDER BY r."host" ASC, r."component" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, port)
	if err != nil {
		logger.Error("Failed to fetch param results by port: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPortAndParamTitle(ctx context.Context, port int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			   AND p."title" = $2 
			 ORDER BY r."host" ASC, r."component" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, port, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by port and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPortAndInternalOrder(ctx context.Context, port int32, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			   AND r."internal_order" = $2 
			 ORDER BY r."host" ASC, r."component" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, port, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by port and internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPortAndInternalOrderAndParamTitle(ctx context.Context, port int32, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			   AND r."internal_order" = $2 
			   AND p."title" = $3 
			 ORDER BY r."host" ASC, r."component" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, port, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by port, internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPortAndComponentTitle(ctx context.Context, port int32, compTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			   AND c."title" = $2 
			 ORDER BY r."host" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, port, compTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by port and component title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPortAndComponentTitleAndParamTitle(ctx context.Context, port int32, compTitle string, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			   AND c."title" = $2 
			   AND p."title" = $3 
			 ORDER BY r."host" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, port, compTitle, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by port, component title and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPortAndComponentTitleAndInternalOrder(ctx context.Context, port int32, compTitle string, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			   AND c."title" = $2 
			   AND r."internal_order" = $3 
			 ORDER BY r."host" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, port, compTitle, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by port, component title and internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByPortAndComponentTitleAndInternalOrderAndParamTitle(ctx context.Context, port int32, compTitle string, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."port" = $1 
			   AND c."title" = $2 
			   AND r."internal_order" = $3 
			   AND p."title" = $4`
	rows = nil
	rows, err = conn.Query(ctx, query, port, compTitle, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by port, component title, internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHost(ctx context.Context, host string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			 ORDER BY r."port" ASC, r."component" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host)
	if err != nil {
		logger.Error("Failed to fetch param results by host: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndParamTitle(ctx context.Context, host string, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND p."title" = $2 
			 ORDER BY r."port" ASC, r."component" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndInternalOrder(ctx context.Context, host string, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."internal_order" = $2 
			 ORDER BY r."port" ASC, r."component" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by host and internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndInternalOrderAndParamTitle(ctx context.Context, host string, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."internal_order" = $2 
			   AND p."title" = $3 
			 ORDER BY r."port" ASC, r."component" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndComponentTitle(ctx context.Context, host string, compTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND c."title" = $2 
			 ORDER BY r."port" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, compTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host and component title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndComponentTitleAndParamTitle(ctx context.Context, host string, compTitle string, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND c."title" = $2 
			   AND p."title" = $3 
			 ORDER BY r."port" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, compTitle, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, component title and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndComponentTitleAndInternalOrder(ctx context.Context, host string, compTitle string, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND c."title" = $2 
			   AND r."internal_order" = $3 
			 ORDER BY r."port" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, compTitle, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by host, component title and internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndComponentTitleAndInternalOrderAndParamTitle(ctx context.Context, host string, compTitle string, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND c."title" = $2 
			   AND r."internal_order" = $3 
			   AND p."title" = $4`
	rows = nil
	rows, err = conn.Query(ctx, query, host, compTitle, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, component title, internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPort(ctx context.Context, host string, port int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			 ORDER BY r."component" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port)
	if err != nil {
		logger.Error("Failed to fetch param results by host and port: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPortAndParamTitle(ctx context.Context, host string, port int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			   AND p."title" = $3 
			 ORDER BY r."component" ASC, r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, port and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPortAndInternalOrder(ctx context.Context, host string, port int32, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			   AND r."internal_order" = $3 
			 ORDER BY r."component" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by host, port and internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPortAndInternalOrderAndParamTitle(ctx context.Context, host string, port int32, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			   AND r."internal_order" = $3 
			   AND p."title" = $4`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, port, internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPortAndComponentTitle(ctx context.Context, host string, port int32, compTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			   AND c."title" = $3 
			 ORDER BY r."internal_order" ASC`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port, compTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, port and component title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPortAndComponentTitleAndParamTitle(ctx context.Context, host string, port int32, compTitle string, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			   AND c."title" = $3 
			   AND p."title" = $4`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port, compTitle, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, port, component title and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPortAndComponentTitleAndInternalOrder(ctx context.Context, host string, port int32, compTitle string, internalOrder int32) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			   AND c."title" = $3 
			   AND r."internal_order" = $4`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port, compTitle, internalOrder)
	if err != nil {
		logger.Error("Failed to fetch param results by host, port, component title and internal order: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}

func GetParamResultsByHostAndPortAndComponentTitleAndInternalOrderAndParamTitle(ctx context.Context, host string, port int32, compTitle string, internalOrder int32, paramTitle string) ([]Result, error) {
	var conn *pgxpool.Pool
	var query string
	var rows pgx.Rows
	var err error
	conn = database.Get()
	query = `SELECT 
			    r."host" AS "host", 
			    r."port" AS "port", 
			    c."title" AS "component_title", 
			    r."internal_order" AS "internal_order", 
			    p."title" AS "param_title", 
			    r."value" AS "value" 
			 FROM public.result r 
			 JOIN public.component c 
			     ON r."component" = c."id" 
			 JOIN public.param p 
			     ON r."param" = p."id" 
			 WHERE r."host" = $1 
			   AND r."port" = $2 
			   AND c."title" = $3 
			   AND r."internal_order" = $4 
			   AND p."title" = $5`
	rows = nil
	rows, err = conn.Query(ctx, query, host, port, compTitle, internalOrder, paramTitle)
	if err != nil {
		logger.Error("Failed to fetch param results by host, port, component title, internal order and param title: %v", err)
		return nil, err
	}
	defer rows.Close()
	var list []Result
	list = []Result{}
	var item Result
	for rows.Next() {
		err = rows.Scan(&item.Host, &item.Port, &item.ComponentTitle, &item.InternalOrder, &item.ParamTitle, &item.Value)
		if err != nil {
			return nil, err
		}
		list = append(list, item)
	}
	return list, rows.Err()
}
