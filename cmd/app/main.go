package main

import (
	"configurator/internal/dao"
	"context"
	"os/exec"

	"configurator/internal/config"
	"configurator/internal/database"
	"configurator/internal/logger"
	"configurator/internal/server"
)

// @title           Configurator
// @version         1.0
// @description     HTTP-сервер на Go с автоматической генерацией Swagger при старте
// @tag.name        1. Модельный каталог: Компоненты
// @tag.name        2. Модельный каталог: Параметры
// @tag.name        3. Модельный каталог: Связи
// @tag.name        4. Парсер: OID
// @tag.name        5. Конфигурация: Индикаторы устройств
// @tag.name        6. Конфигурация: Индикаторы параметров
// @tag.name        7. Конфигурация: Сопоставления параметров
// @tag.name        8. Конфигурация: Структура компонентов устройства
// @tag.name        9. Конфигурация: Связь компонентов устройства и сопоставлений параметров
// @tag.name        10. Конфигурация: Конфигурации по-умолчанию
// @tag.name        11. Конфигурация: Конфигурации устройств
// @tag.name        12. Конфигурация: Пороги
// @tag.name        13. Конфигурация: Привязка порога к конфигурации по умолчанию
// @tag.name        14. Конфигурация: Привязка порога к рабочей конфигурации
// @tag.name        15. Результат: Экспортировать БД в SQL скрипт
// @contact.name    Lorgar Avrelian
// @contact.url     https://github.com/Lorgar-Avrelian
// @contact.email   victor-14-244@mail.ru
// @license.name    Apache 2.0
// @license.url     http://apache.org
// @host            localhost:8082
func main() {
	configPath := "cmd/config.yml"
	config.Init(configPath)
	logger.Init(config.Get().Logger.Level)
	//generateSwagger()
	database.Init()
	if err := dao.LoadEnumsFromDB(context.Background()); err != nil {
		logger.Fatalf("Критическая ошибка при загрузке справочников enum из БД: %v", err)
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
		logger.Warn("Не удалось автоматически обновить Swagger через код.")
		logger.Error("Детали ошибки генератора:\n%s", string(output))
		return
	}

	logger.Info("Файлы Swagger успешно обновлены")
}
