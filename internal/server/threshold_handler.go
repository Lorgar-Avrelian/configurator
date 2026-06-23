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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ThresholdDto
	res, err = service.CreateThreshold(c.Request.Context(), input)
	if err != nil {
		logger.Error("Service error occurred while creating threshold: %v", err)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var res *dto.ThresholdDto
	res, err = service.GetThresholdByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving threshold %d: %v", id, err)
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
		logger.Error("Service error occurred while retrieving all thresholds: %v", err)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var input dto.ThresholdCreateDto
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ThresholdDto
	res, err = service.UpdateThreshold(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Service error occurred while updating threshold %d: %v", id, err)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	if 1 > id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threshold ID format"})
		return
	}
	var found bool
	found, err = service.DeleteThreshold(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while deleting threshold %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Threshold not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
