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

// CreateParamIndicator создает индикатор параметров
// @Summary         Создать индикатор параметров
// @Tags            6. Конфигурация: Индикаторы параметров
// @Accept          json
// @Produce         json
// @Param           request body dto.ParamIndicatorCreateDto true "Данные индикатора параметров"
// @Success         201  {object}  dto.ParamIndicatorDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/param [post]
func CreateParamIndicator(c *gin.Context) {
	var input dto.ParamIndicatorCreateDto
	var err error
	var res *dto.ParamIndicatorDto
	err = c.ShouldBindJSON(&input)
	if err != nil {
		logger.Warnf("Validation failed during param indicator creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	if input.DotterNotation == nil && input.OidID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body value"})
		return
	}
	res, err = service.CreateParamIndicator(c.Request.Context(), input)
	if err != nil {
		logger.Errorf("Service error occurred while creating param indicator: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetParamIndicator возвращает индикатор параметров по ID
// @Summary         Получить индикатор параметров по ID
// @Tags            6. Конфигурация: Индикаторы параметров
// @Produce         json
// @Param           id   path      int  true  "ID Индикатора"
// @Success         200  {object}  dto.ParamIndicatorDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/param/{id} [get]
func GetParamIndicator(c *gin.Context) {
	var id int64
	var err error
	var res *dto.ParamIndicatorDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param indicator ID format"})
		return
	}
	res, err = service.GetParamIndicatorByID(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while retrieving param indicator %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve param indicator"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Param indicator not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllParamIndicators возвращает все индикаторы параметров
// @Summary         Получить все индикаторы параметров
// @Tags            6. Конфигурация: Индикаторы параметров
// @Produce         json
// @Success         200  {array}   dto.ParamIndicatorDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/params [get]
func GetAllParamIndicators(c *gin.Context) {
	var res []dto.ParamIndicatorDto
	var err error
	res, err = service.GetAllParamIndicators(c.Request.Context())
	if err != nil {
		logger.Errorf("Service error occurred while retrieving all param indicators: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve all param indicators"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateParamIndicator обновляет индикатор параметров по ID
// @Summary         Обновить индикатор параметров по ID
// @Tags            6. Конфигурация: Индикаторы параметров
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Индикатора"
// @Param           request body dto.ParamIndicatorCreateDto true "Новые данные"
// @Success         200  {object}  dto.ParamIndicatorDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/param/{id} [put]
func UpdateParamIndicator(c *gin.Context) {
	var id int64
	var err error
	var input dto.ParamIndicatorCreateDto
	var res *dto.ParamIndicatorDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param indicator ID format"})
		return
	}
	err = c.ShouldBindJSON(&input)
	if err != nil {
		logger.Warnf("Validation failed during param indicator update for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	if input.DotterNotation == nil && input.OidID == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body value"})
		return
	}
	res, err = service.UpdateParamIndicator(c.Request.Context(), id, input)
	if err != nil {
		logger.Errorf("Service error occurred while updating param indicator %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Param indicator not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteParamIndicator удаляет индикатор параметров по ID
// @Summary         Удалить индикатор параметров по ID
// @Tags            6. Конфигурация: Индикаторы параметров
// @Param           id   path      int  true  "ID Индикатора"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/param/{id} [delete]
func DeleteParamIndicator(c *gin.Context) {
	var id int64
	var err error
	var found bool
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid param indicator ID format"})
		return
	}
	found, err = service.DeleteParamIndicator(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while deleting param indicator %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete param indicator"})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Param indicator not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
