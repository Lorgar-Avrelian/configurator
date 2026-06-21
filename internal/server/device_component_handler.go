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

// CreateDeviceComponent создает новый узел состава устройства
// @Summary         Создать узел состава устройства
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Accept          json
// @Produce         json
// @Param           request body dto.DeviceComponentCreateDto true "Данные узла"
// @Success         201  {object}  dto.DeviceComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/device-component [post]
func CreateDeviceComponent(c *gin.Context) {
	var input dto.DeviceComponentCreateDto
	var err error
	var res *dto.DeviceComponentDto
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during device component creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	res, err = service.CreateDeviceComponent(c.Request.Context(), input)
	if err != nil {
		logger.Error("Service error occurred while creating device component: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create device component"})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetDeviceComponent возвращает узел состава по ID вместе с иерархией подчиненности
// @Summary         Получить узел состава по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Produce         json
// @Param           id   path      int  true  "ID Узла"
// @Success         200  {object}  dto.DeviceComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/device-component/{id} [get]
func GetDeviceComponent(c *gin.Context) {
	var id int64
	var err error
	var res *dto.DeviceComponentDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device component ID format"})
		return
	}
	res, err = service.GetDeviceComponentByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving device component %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve device component"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device component not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllDeviceComponents возвращает все узлы с их иерархией подчиненности
// @Summary         Получить всю структуру подчиненности устройств
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Produce         json
// @Success         200  {array}   dto.DeviceComponentDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/device-components [get]
func GetAllDeviceComponents(c *gin.Context) {
	var res []dto.DeviceComponentDto
	var err error
	res, err = service.GetAllDeviceComponents(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while retrieving all device components: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve all device components"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateDeviceComponent обновляет метаданные узла по ID
// @Summary         Обновить узел состава по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Узла"
// @Param           request body dto.DeviceComponentCreateDto true "Новые данные"
// @Success         200  {object}  dto.DeviceComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/device-component/{id} [put]
func UpdateDeviceComponent(c *gin.Context) {
	var id int64
	var err error
	var input dto.DeviceComponentCreateDto
	var res *dto.DeviceComponentDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device component ID format"})
		return
	}
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during device component update for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	res, err = service.UpdateDeviceComponent(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Service error occurred while updating device component %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update device component"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device component not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteDeviceComponent удаляет узел из дерева подчиненности
// @Summary         Удалить узел состава по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Param           id   path      int  true  "ID Узла"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/device-component/{id} [delete]
func DeleteDeviceComponent(c *gin.Context) {
	var id int64
	var err error
	var found bool
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device component ID format"})
		return
	}
	found, err = service.DeleteDeviceComponent(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while deleting device component %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete device component"})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device component not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
