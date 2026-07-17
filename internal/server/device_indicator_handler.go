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

// CreateIndicator создаёт новую запись индикатора
// @Summary         Создать индикатор устройства
// @Tags            5. Конфигурация: Индикаторы устройств
// @Accept          json
// @Produce         json
// @Param           request body dto.DeviceIndicatorCreateDto true "Данные индикатора"
// @Success         201  {object}  dto.DeviceIndicatorDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/device [post]
func CreateIndicator(c *gin.Context) {
	var input dto.DeviceIndicatorCreateDto
	var err error
	var res *dto.DeviceIndicatorDto
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warnf("Validation failed during device indicator creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	res, err = service.CreateDeviceIndicator(c.Request.Context(), input)
	if err != nil {
		logger.Errorf("Service error occurred while creating device indicator: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create device indicator"})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetIndicator возвращает индикатор по ID
// @Summary         Получить индикатор по ID
// @Tags            5. Конфигурация: Индикаторы устройств
// @Produce         json
// @Param           id   path      int  true  "ID Индикатора"
// @Success         200  {object}  dto.DeviceIndicatorDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/device/{id} [get]
func GetIndicator(c *gin.Context) {
	var id int64
	var err error
	var res *dto.DeviceIndicatorDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device indicator ID format"})
		return
	}
	res, err = service.GetDeviceIndicatorByID(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while retrieving device indicator %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve device indicator"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device indicator not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllIndicators возвращает полный список записей таблицы
// @Summary         Получить все индикаторы устройств
// @Tags            5. Конфигурация: Индикаторы устройств
// @Produce         json
// @Success         200  {array}   dto.DeviceIndicatorDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/devices [get]
func GetAllIndicators(c *gin.Context) {
	var res []dto.DeviceIndicatorDto
	var err error
	res, err = service.GetAllDeviceIndicators(c.Request.Context())
	if err != nil {
		logger.Errorf("Service error occurred while retrieving all device indicators: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve all device indicators"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateIndicator изменяет запись индикатора по ID
// @Summary         Обновить индикатор по ID
// @Tags            5. Конфигурация: Индикаторы устройств
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Индикатора"
// @Param           request body dto.DeviceIndicatorCreateDto true "Новые данные"
// @Success         200  {object}  dto.DeviceIndicatorDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/device/{id} [put]
func UpdateIndicator(c *gin.Context) {
	var id int64
	var err error
	var input dto.DeviceIndicatorCreateDto
	var res *dto.DeviceIndicatorDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if id == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device indicator ID value"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device indicator ID format"})
		return
	}
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warnf("Validation failed during device indicator update for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	res, err = service.UpdateDeviceIndicator(c.Request.Context(), id, input)
	if err != nil {
		logger.Errorf("Service error occurred while updating device indicator %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update device indicator"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device indicator not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteIndicator удаляет запись индикатора по ID
// @Summary         Удалить индикатор по ID
// @Tags            5. Конфигурация: Индикаторы устройств
// @Param           id   path      int  true  "ID Индикатора"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicator/device/{id} [delete]
func DeleteIndicator(c *gin.Context) {
	var id int64
	var err error
	var found bool
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device indicator ID format"})
		return
	}
	found, err = service.DeleteDeviceIndicator(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while deleting device indicator %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete device indicator"})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device indicator not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
