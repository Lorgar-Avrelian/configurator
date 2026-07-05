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

// CreateDeviceComponent создаёт новую составную часть устройства
// @Summary         Создать составную часть устройства
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Accept          json
// @Produce         json
// @Param           request body dto.DeviceComponentCreateDto true "Данные составной части"
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

// GetDeviceComponent возвращает составную часть устройства по ID с указанием положения внутри структуры устройства
// @Summary         Получить составную часть устройства по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Produce         json
// @Param           id   path      int  true  "ID составной части устройства"
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

// GetAllDeviceComponents возвращает всю структуру составных частей устройств
// @Summary         Получить всю структуру составных частей устройств
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

// UpdateDeviceComponent обновляет метаданные составной части устройства по ID
// @Summary         Обновить составную часть устройства по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID составной части"
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

// DeleteDeviceComponent удаляет составную часть устройства
// @Summary         Удалить составную часть устройства по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Param           id   path      int  true  "ID составной части"
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

// GetDeviceComponentOwn возвращает составную часть устройства по ID без указания положения внутри структуры устройства
// @Summary         Получить изолированную составную часть устройства по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Produce         json
// @Param           id   path      int  true  "ID составной части"
// @Success         200  {object}  dto.DeviceComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/device-component/{id}/own [get]
func GetDeviceComponentOwn(c *gin.Context) {
	var id int64
	var err error
	var res *dto.DeviceComponentDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device component ID format"})
		return
	}
	res, err = service.GetDeviceComponentByIDOwn(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving single device component %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve device component"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device component not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ChangeDeviceComponentDataHandler изменяет ID составной части устройства
// @Summary         Изменить ID составной части устройства
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Produce         json
// @Param           prevId path      int  true  "Предыдущий ID составной части устройства"
// @Param           newId  path      int  true  "Новый ID составной части устройства"
// @Success         200  "OK"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/device-component/{prevId}/{newId} [patch]
func ChangeDeviceComponentDataHandler(c *gin.Context) {
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
	success, err = service.ChangeDeviceComponentData(c.Request.Context(), prevId, newId)
	if err != nil {
		logger.Error("Service error during changing device component data from %d to %d: %v", prevId, newId, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation rejected by service"})
		return
	}
	c.Status(http.StatusOK)
}
