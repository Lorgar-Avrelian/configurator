package database

import (
	"context"
	"filler/internal/config"
	log "filler/internal/logger"
	"fmt"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

func Init() {
	cfg := config.Get()
	if cfg == nil {
		log.Fatalf("Ошибка подключения к БД: конфигурация не инициализирована")
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.URL,
		cfg.DB.DB,
	)

	log.Printf("Подключение к базе данных %s на %s...", cfg.DB.DB, cfg.DB.URL)

	var err error
	conn, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	log.Println("Соединение с базой данных успешно установлено через нативный pgx")
}

func Get() *pgx.Conn {
	return conn
}
