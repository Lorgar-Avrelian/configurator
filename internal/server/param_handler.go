package server

import (
	"filler/internal/logger"
	"net/http"
	"strconv"

	"filler/internal/dao"
	"filler/internal/dto"
	_ "filler/internal/model"

	"github.com/gin-gonic/gin"
)

// CreateParam создает новый параметр
// @Summary         Создать параметр
// @Description     Создает новый системный параметр в таблице public.param
// @Tags            2. Модельный каталог: Параметры
// @Accept          json
// @Produce         json
// @Param           request body dto.ParamCreate true "Данные параметра"
// @Success         201  {object}  model.Param
// @Failure         400  {object}  map[string]string "Ошибка валидации"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/params [post]
func CreateParam(c *gin.Context) {
	var input dto.ParamCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при создании параметра: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.CreateParam(c.Request.Context(), input)
	if err != nil {
		logger.Error("Ошибка DAO при создании параметра: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetParam возвращает параметр по ID
// @Summary         Получить параметр по ID
// @Tags            2. Модельный каталог: Параметры
// @Produce         json
// @Param           id   path      int  true  "ID Параметра"
// @Success         200  {object}  model.Param
// @Failure         400  {object}  map[string]string "Неверный формат ID"
// @Failure         404  {object}  map[string]string "Параметр не найден"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/params/{id} [get]
func GetParam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID параметра"})
		return
	}
	res, err := dao.GetParamByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при получении параметра %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Параметр не найден"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateParam обновляет параметры
// @Summary         Обновить параметр
// @Tags            2. Модельный каталог: Параметры
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Параметра"
// @Param           request body dto.ParamUpdate true "Новые данные"
// @Success         200  {object}  model.Param
// @Failure         400  {object}  map[string]string "Ошибка валидации или неверный ID"
// @Failure         404  {object}  map[string]string "Параметр не найден"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/params/{id} [put]
func UpdateParam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID параметра"})
		return
	}
	var input dto.ParamUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при обновлении параметра %d: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.UpdateParam(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Ошибка DAO при обновлении параметра %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Параметр не найден для обновления"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteParam удаляет параметр
// @Summary         Удалить параметр
// @Tags            2. Модельный каталог: Параметры
// @Param           id   path      int  true  "ID Параметра"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string "Неверный формат ID"
// @Failure         404  {object}  map[string]string "Параметр не найден"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/params/{id} [delete]
func DeleteParam(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID параметра"})
		return
	}
	found, err := dao.DeleteParam(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при удалении параметра %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Параметр не найден"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetUnattachedParams возвращает параметры без связей
// @Summary			Получить непривязанные параметры
// @Description		Возвращает список всех параметров, которые не привязаны ни к одной составной части устройства
// @Tags            2. Модельный каталог: Параметры
// @Produce			json
// @Success			200   {array}   model.Param
// @Failure			500  {object}  map[string]string "Ошибка базы данных"
// @Router			/api/v1/params/unattached [get]
func GetUnattachedParams(c *gin.Context) {
	res, err := dao.GetUnattachedParams(c.Request.Context())
	if err != nil {
		logger.Error("Ошибка DAO при получении непривязанных параметров: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// SearchParams выполняет полнотекстовый поиск параметров
// @Summary         Поиск параметров по строке
// @Description     Ищет параметры, у которых title, name_en, name_ru, description_en или description_ru частично совпадают с переданной строкой query
// @Tags            2. Модельный каталог: Параметры
// @Produce         json
// @Param           query query     string  true  "Строка поиска"
// @Success         200   {array}   model.Param
// @Failure         400   {object}  map[string]string "Пустой запрос"
// @Failure         500   {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/params/search [get]
func SearchParams(c *gin.Context) {
	q := c.Query("query")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'query' не должен быть пустым"})
		return
	}
	res, err := dao.SearchParams(c.Request.Context(), q)
	if err != nil {
		logger.Error("Ошибка DAO при поиске параметров по запросу '%s': %v", q, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
