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

// CreateConfiguration создаёт рабочую конфигурацию
// @Summary         Создать рабочую конфигурацию
// @Tags            11. Конфигурация: Конфигурации устройств
// @Produce         json
// @Param           indicator          query    int  true  "ID индикатора устройства"
// @Param           device_component   query    int  false "ID компонента устройства"
// @Success         201  {object}  dto.ConfigurationDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/configuration [post]
func CreateConfiguration(c *gin.Context) {
	var indID int64
	var err error
	indID, err = strconv.ParseInt(c.Query("indicator"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing indicator ID parameter"})
		return
	}
	if 1 > indID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing indicator ID parameter"})
		return
	}
	var dcID *int64
	var dcRaw string
	dcID = nil
	dcRaw = c.Query("device_component")
	if dcRaw != "" {
		var parsedDc int64
		parsedDc, err = strconv.ParseInt(dcRaw, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device component ID format"})
			return
		}
		if 1 > parsedDc {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Device component ID must be positive"})
			return
		}
		dcID = &parsedDc
	}
	var res *dto.ConfigurationDto
	res, err = service.CreateConfiguration(c.Request.Context(), indID, dcID)
	if err != nil {
		logger.Errorf("Service error occurred while creating configuration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetConfiguration возвращает конфигурацию по ID
// @Summary         Получить рабочую конфигурацию по ID
// @Tags            11. Конфигурация: Конфигурации устройств
// @Produce         json
// @Param           id   path      int  true  "ID конфигурации"
// @Success         200  {object}  dto.ConfigurationDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/configuration/{id} [get]
func GetConfiguration(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid configuration ID format"})
		return
	}
	if 1 > id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid configuration ID format"})
		return
	}
	var res *dto.ConfigurationDto
	res, err = service.GetConfigurationByID(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while retrieving configuration %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Configuration not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllConfigurations возвращает все конфигурации
// @Summary         Получить все рабочие конфигурации
// @Tags            11. Конфигурация: Конфигурации устройств
// @Produce         json
// @Success         200  {array}   dto.ConfigurationDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/configurations [get]
func GetAllConfigurations(c *gin.Context) {
	var res []dto.ConfigurationDto
	var err error
	res, err = service.GetAllConfigurations(c.Request.Context())
	if err != nil {
		logger.Errorf("Service error occurred while retrieving all configurations: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateConfiguration обновляет конфигурацию по ID
// @Summary         Обновить рабочую конфигурацию по ID
// @Tags            11. Конфигурация: Конфигурации устройств
// @Produce         json
// @Param           id                 path     int  true  "ID конфигурации"
// @Param           indicator          query    int  true  "Новый ID индикатора устройства"
// @Param           device_component   query    int  false "Новый ID компонента устройства"
// @Success         200  {object}  dto.ConfigurationDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/configuration/{id} [put]
func UpdateConfiguration(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid configuration ID format"})
		return
	}
	if 1 > id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid configuration ID format"})
		return
	}
	var indID int64
	indID, err = strconv.ParseInt(c.Query("indicator"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing indicator ID parameter"})
		return
	}
	if 1 > indID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing indicator ID parameter"})
		return
	}
	var dcID *int64
	var dcRaw string
	dcID = nil
	dcRaw = c.Query("device_component")
	if dcRaw != "" {
		var parsedDc int64
		parsedDc, err = strconv.ParseInt(dcRaw, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid device component ID format"})
			return
		}
		if 1 > parsedDc {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Device component ID must be positive"})
			return
		}
		dcID = &parsedDc
	}
	var res *dto.ConfigurationDto
	res, err = service.UpdateConfiguration(c.Request.Context(), id, indID, dcID)
	if err != nil {
		logger.Errorf("Service error occurred while updating configuration %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Configuration not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteConfiguration удаляет конфигурацию по ID
// @Summary         Удалить рабочую конфигурацию по ID
// @Tags            11. Конфигурация: Конфигурации устройств
// @Param           id   path      int  true  "ID конфигурации"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/configuration/{id} [delete]
func DeleteConfiguration(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid configuration ID format"})
		return
	}
	if 1 > id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid configuration ID format"})
		return
	}
	var found bool
	found, err = service.DeleteConfiguration(c.Request.Context(), id)
	if err != nil {
		logger.Errorf("Service error occurred while deleting configuration %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Configuration not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
