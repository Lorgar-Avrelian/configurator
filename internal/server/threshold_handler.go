package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	_ "configurator/internal/model"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateThreshold создаёт порог
// @Summary         Создать порог
// @Tags            12. Конфигурация: Пороги
// @Accept          json
// @Produce         json
// @Param           request body dto.ThresholdCreateDto true "Данные порога"
// @Success         201  {object}  dto.ThresholdDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold [post]
func CreateThreshold(c *gin.Context) {
	var input dto.ThresholdCreateDto
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Errorf("Failed to bind JSON for threshold creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ThresholdDto
	res, err = service.CreateThreshold(c.Request.Context(), input)
	if err != nil {
		logger.Errorf("Service error occurred while creating threshold: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetThreshold возвращает порог по ID
// @Summary         Получить порог по ID
// @Tags            12. Конфигурация: Пороги
// @Produce         json
// @Param           id   path      int  true  "ID порога"
// @Success         200  {object}  dto.ThresholdDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold/{id} [get]
func GetThreshold(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse threshold ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		logger.Errorf("Threshold ID validation failed for ID: %d", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var res *dto.ThresholdDto
	res, err = service.GetThresholdByID(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while retrieving threshold %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Threshold not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllThresholds возвращает все пороги
// @Summary         Получить все пороги
// @Tags            12. Конфигурация: Пороги
// @Produce         json
// @Success         200  {array}   dto.ThresholdDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/thresholds [get]
func GetAllThresholds(c *gin.Context) {
	var res []dto.ThresholdDto
	var err error
	res, err = service.GetAllThresholds(c.Request.Context())
	if err != nil {
		logger.Errorf("Service error occurred while retrieving all thresholds: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateThreshold обновляет порог по ID
// @Summary         Обновить порог по ID
// @Tags            12. Конфигурация: Пороги
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID порога"
// @Param           request body dto.ThresholdCreateDto true "Новые данные"
// @Success         200  {object}  dto.ThresholdDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold/{id} [put]
func UpdateThreshold(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse threshold ID for update: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		logger.Errorf("Threshold ID validation failed for update ID: %d", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var input dto.ThresholdCreateDto
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Errorf("Failed to bind JSON for threshold update: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ThresholdDto
	res, err = service.UpdateThreshold(c.Request.Context(), id, input)
	if err != nil {
		logger.Errorf("Service error occurred while updating threshold %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Threshold not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteThreshold удаляет порог по ID
// @Summary         Удалить порог по ID
// @Tags            12. Конфигурация: Пороги
// @Param           id   path      int  true  "ID порога"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold/{id} [delete]
func DeleteThreshold(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse threshold ID for deletion: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		logger.Errorf("Threshold ID validation failed for deletion ID: %d", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var found bool
	found, err = service.DeleteThreshold(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while deleting threshold %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Threshold not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// ChangeThresholdDataHandler изменяет ID порога
// @Summary         Изменить ID порога
// @Tags            12. Конфигурация: Пороги
// @Produce         json
// @Param           prevId path      int  true  "Предыдущий ID порога"
// @Param           newId  path      int  true  "Новый ID порога"
// @Success         200  "OK"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold/{prevId}/{newId} [patch]
func ChangeThresholdDataHandler(c *gin.Context) {
	var prevId int64
	var newId int64
	var err error
	prevId, err = strconv.ParseInt(c.Param("prevId"), 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse prevId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prevId format"})
		return
	}
	newId, err = strconv.ParseInt(c.Param("newId"), 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse newId: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid newId format"})
		return
	}
	var success bool
	success, err = service.ChangeThresholdData(c.Request.Context(), prevId, newId)
	if err != nil {
		logger.Errorf("Service error during changing threshold data from %d to %d: %v", prevId, newId, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation rejected by service"})
		return
	}
	c.Status(http.StatusOK)
}

// CreateThresholdFromString создаёт порог из эквивалентного строкового выражения
// @Summary         Создать порог из эквивалентной строки
// @Tags            12. Конфигурация: Пороги
// @Accept          plain
// @Produce         json
// @Param           request body string true "Эквивалентное строковое выражение порога"
// @Success         201  {object}  dto.ThresholdDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold/from-string [post]
func CreateThresholdFromString(c *gin.Context) {
	var bodyBytes []byte
	var err error
	bodyBytes, err = c.GetRawData()
	if err != nil {
		logger.Errorf("Failed to read text expression raw body bytes: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text request body content"})
		return
	}
	var res dto.ThresholdDto
	res, err = service.CreateThresholdFromString(c.Request.Context(), string(bodyBytes))
	if err != nil {
		logger.Errorf("Service error occurred while creating threshold from string: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetThresholdStringByID возвращает эквивалентную логическую строку порога по ID
// @Summary         Получить эквивалентную строку выражения порога по ID
// @Tags            12. Конфигурация: Пороги
// @Produce         plain
// @Param           id   path      int  true  "ID порога"
// @Success         200  {string}  string "Эквивалентное строковое выражение порога"
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold/{id}/from-string [get]
func GetThresholdStringByID(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse threshold string ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		logger.Errorf("Threshold string ID validation failed for ID: %d", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var res string
	res, err = service.GetThresholdStringByID(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while retrieving threshold expression %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.String(http.StatusOK, res)
}

// UpdateThresholdFromString обновляет порог по ID посредством эквивалентного выражения
// @Summary         Обновить порог по ID из строки
// @Tags            12. Конфигурация: Пороги
// @Accept          plain
// @Produce         json
// @Param           id      path      int  true  "ID порога"
// @Param           request body string true "Эквивалентное строковое выражение порога"
// @Success         200  {object}  dto.ThresholdDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/threshold/{id}/from-string [put]
func UpdateThresholdFromString(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		logger.Errorf("Failed to parse threshold ID for update from string: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		logger.Errorf("Threshold ID validation failed for string update ID: %d", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var bodyBytes []byte
	bodyBytes, err = c.GetRawData()
	if err != nil {
		logger.Errorf("Failed to read text expression update raw body bytes: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid text request body content"})
		return
	}
	var res dto.ThresholdDto
	res, err = service.UpdateThresholdFromString(c.Request.Context(), id, string(bodyBytes))
	if err != nil {
		logger.Errorf("Service error occurred while updating threshold %d from string: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
