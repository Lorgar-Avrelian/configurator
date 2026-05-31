package server

import (
	"filler/internal/dao"
	"filler/internal/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SaveResult выгружает содержание основных таблиц в файлы формата SQL
// @Summary         Экспортировать результаты в SQL скрипт
// @Description     Выгружает данные основных таблиц в структурированные .sql файлы в директорию 'sql' в корне проекта
// @Tags            12. Результат: Экспортировать БД в SQL скрипт
// @Accept          json
// @Produce         json
// @Success         200  {object}  map[string]string "Сообщение об успешном создании файла"
// @Failure         400  {object}  map[string]string "Ошибка валидации данных"
// @Failure         500  {object}  map[string]string "Внутренняя ошибка при записи файла или чтении БД"
// @Router          /api/v1/save-result [post]
func SaveResult(c *gin.Context) {
	err := dao.ExportDatabaseToSqlFiles(c.Request.Context())
	if err != nil {
		logger.Error("Критическая ошибка при автономном экспорте таблиц в SQL: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	logger.Info("Экспорт всех таблиц базы данных в директорию sql/ успешно завершен")
	c.JSON(http.StatusOK, gin.H{
		"message":   "Все таблицы базы данных успешно выгружены в виде чистых SQL дампов",
		"directory": "sql/",
	})
}
