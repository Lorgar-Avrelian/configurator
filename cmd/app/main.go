package main

import (
	"context"
	"filler/internal/dao"
	"os/exec"

	"filler/internal/config"
	"filler/internal/database"
	"filler/internal/logger"
	"filler/internal/server"
)

// @title           DB Filler API
// @version         1.0
// @description     HTTP-сервер на Go с автоматической генерацией Swagger при старте
// @contact.name    Victor Tokovenko
// @contact.url     https://opk-bulat.ru/
// @contact.email   tokovenko@opk-bulat.ru
// @license.name    Apache 2.0
// @license.url     http://apache.org
// @host            localhost:8082
func main() {
	configPath := "cmd/config.yml"
	config.Init(configPath)
	logger.Init(config.Get().Logger.Level)
	generateSwagger()
	database.Init()
	if err := dao.LoadEnumsFromDB(context.Background()); err != nil {
		logger.Fatalf("Критическая ошибка при загрузке справочников энумов из БД: %v", err)
	}
	srv := server.NewServer()
	logger.Info("Запуск HTTP-сервера на порту %d...", config.Get().Server.Port)
	if err := srv.Run(); err != nil {
		logger.Fatalf("Ошибка работы HTTP-сервера: %v", err)
	}
}

func generateSwagger() {
	logger.Info("Автогенерация файлов Swagger...")

	cmd := exec.Command("swag", "init",
		"-g", "cmd/app/main.go",
		"-d", "./,internal/server,internal/dto,internal/model",
		"--parseDependency",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Warn("Не удалось автоматически обновить Swagger через код (возможно, утилита swag не добавлена в PATH системы).")
		logger.Debug("Детали ошибки генератора:\n%s", string(output))
		return
	}

	logger.Info("Файлы Swagger успешно обновлены на лету!")
}
