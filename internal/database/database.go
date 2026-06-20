package database

import (
	"configurator/internal/config"
	log "configurator/internal/logger"
	"context"
	"fmt"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func Init() {
	cfg := config.Get()
	if cfg == nil {
		log.Fatalf("DB connection error: configuration is not initialized")
	}
	maxConns := runtime.GOMAXPROCS(0)
	if maxConns < 10 {
		maxConns = 10
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.URL,
		cfg.DB.DB,
	)
	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Couldn't parse the database configuration: %v", err)
	}
	poolCfg.MaxConns = int32(maxConns)
	log.Printf("Connecting to the %s database on %s with pool size %d", cfg.DB.DB, cfg.DB.URL, maxConns)
	pool, err = pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		log.Fatalf("Couldn't connect to the database: %v", err)
	}
	log.Println("DB Connection has been established successfully")
}

func Get() *pgxpool.Pool {
	return pool
}
