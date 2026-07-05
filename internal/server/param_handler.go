package server

import (
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"configurator/internal/dto"
	_ "configurator/internal/model"

	"github.com/gin-gonic/gin"
)

// CreateParam создаёт новый параметр
// @Summary         Создать параметр
// @Description     Создает новый системный параметр в таблице public.param
// @Tags            2. Модельный каталог: Параметры
// @Accept          json
// @Produce         json
// @Param           request body dto.ParamCreateDto true "Данные параметра"
// @Success         201  {object}  dto.ParamDto
// @Failure         400  {object}  map[string]string "Ошибка валидации"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/param [post]
func CreateParam(c *gin.Context) {
	var input dto.ParamCreateDto
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during parameter creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ParamDto
	res, err = service.CreateParam(c.Request.Context(), input)
	if err != nil {
		logger.Error("Service error occurred while creating parameter: %v", err)
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
// @Success         200  {object}  dto.ParamDto
// @Failure         400  {object}  map[string]string "Неверный формат ID"
// @Failure         404  {object}  map[string]string "Параметр не найден"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/param/{id} [get]
func GetParam(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID format"})
		return
	}
	var res *dto.ParamDto
	res, err = service.GetParamByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving parameter %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parameter not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateParam обновляет параметр
// @Summary         Обновить параметр
// @Tags            2. Модельный каталог: Параметры
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Параметра"
// @Param           request body dto.ParamUpdateDto true "Новые данные"
// @Success         200  {object}  dto.ParamDto
// @Failure         400  {object}  map[string]string "Ошибка валидации или неверный ID"
// @Failure         404  {object}  map[string]string "Параметр не найден"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/param/{id} [put]
func UpdateParam(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID format"})
		return
	}
	var input dto.ParamUpdateDto
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during parameter update for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ParamDto
	res, err = service.UpdateParam(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Service error occurred while updating parameter %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parameter not found for update"})
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
// @Router          /api/v1/param/{id} [delete]
func DeleteParam(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID format"})
		return
	}
	var found bool
	found, err = service.DeleteParam(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while deleting parameter %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Parameter not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetAllParams возвращает все параметры модельного каталога
// @Summary         Получить все параметры
// @Tags            2. Модельный каталог: Параметры
// @Produce         json
// @Success         200  {array}   dto.ParamDto
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/params [get]
func GetAllParams(c *gin.Context) {
	var res []dto.ParamDto
	var err error
	res, err = service.GetAllParams(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while retrieving all parameters: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetUnattachedParams возвращает параметры без связей
// @Summary			Получить непривязанные параметры
// @Description		Возвращает список всех параметров, которые не привязаны ни к одной составной части устройства
// @Tags            2. Модельный каталог: Параметры
// @Produce			json
// @Success			200   {array}   dto.ParamDto
// @Failure			500  {object}  map[string]string "Ошибка базы данных"
// @Router			/api/v1/param/unattached [get]
func GetUnattachedParams(c *gin.Context) {
	var res []dto.ParamDto
	var err error
	res, err = service.GetUnattachedParams(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while retrieving unattached parameters: %v", err)
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
// @Success         200   {array}   dto.ParamDto
// @Failure         400   {object}  map[string]string "Пустой запрос"
// @Failure         500   {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/param/search [get]
func SearchParams(c *gin.Context) {
	var queryText string
	queryText = c.Query("query")
	if queryText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query parameter cannot be empty"})
		return
	}
	var res []dto.ParamDto
	var err error
	res, err = service.SearchParams(c.Request.Context(), queryText)
	if err != nil {
		logger.Error("Service error occurred while searching parameters for query '%s': %v", queryText, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetComponentsByParam возвращает список компонентов, к которым напрямую привязан параметр
// @Summary         Получить компоненты по ID параметра
// @Description     Возвращает список компонентов, для которых параметр с данным ID является их собственным, а не унаследованным
// @Tags            2. Модельный каталог: Параметры
// @Produce         json
// @Param           id   path      int  true  "ID Параметра"
// @Success         200  {array}   dto.ComponentDto
// @Failure         400  {object}  map[string]string "Неверный формат ID"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/param/search/{id} [get]
func GetComponentsByParam(c *gin.Context) {
	var id int64
	var err error
	id = 0
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID format"})
		return
	}
	var res []dto.ComponentDto
	res, err = service.GetComponentsByDirectParamID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while fetching components for direct param %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ChangeParamDataHandler изменяет ID параметра
// @Summary         Изменить ID параметра
// @Tags            2. Модельный каталог: Параметры
// @Produce         json
// @Param           prevId path      int  true  "Предыдущий ID параметра"
// @Param           newId  path      int  true  "Новый ID параметра"
// @Success         200  "OK"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param/{prevId}/{newId} [patch]
func ChangeParamDataHandler(c *gin.Context) {
	var prevId int64
	var newId int64
	var err error
	prevId, err = strconv.ParseInt(c.Param("prevId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prevId format"})
		return
	}
	newId, err = strconv.ParseInt(c.Param("newId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid newId format"})
		return
	}
	var success bool
	success, err = service.ChangeParamData(c.Request.Context(), prevId, newId)
	if err != nil {
		logger.Error("Service error during changing param data from %d to %d: %v", prevId, newId, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation rejected by service"})
		return
	}
	c.Status(http.StatusOK)
}
