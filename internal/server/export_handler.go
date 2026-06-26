package server

import (
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SaveResult выгружает содержание основных таблиц в файлы формата SQL
// @Summary         Экспортировать результаты в SQL скрипт
// @Description     Выгружает данные основных таблиц в структурированные .sql файлы в директорию 'sql' в корне проекта
// @Tags            16. Результат: Экспортировать БД в SQL скрипт
// @Accept          json
// @Produce         json
// @Success         200  {object}  map[string]string "Сообщение об успешном создании файла"
// @Failure         400  {object}  map[string]string "Ошибка валидации данных"
// @Failure         500  {object}  map[string]string "Внутренняя ошибка при записи файла или чтении БД"
// @Router          /api/v1/save-result [post]
func SaveResult(c *gin.Context) {
	var err error
	err = service.ExportDatabase(c.Request.Context())
	if err != nil {
		logger.Error("Critical error during autonomous export of tables to SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info("Export of all database tables to sql/ directory completed successfully")
	c.JSON(http.StatusOK, gin.H{
		"message":   "All database tables successfully dumped as clean SQL files",
		"directory": "sql/",
	})
}
