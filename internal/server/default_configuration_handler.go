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

// CreateDefaultConfiguration создаёт конфигурацию по умолчанию
// @Summary         Создать конфигурацию по умолчанию
// @Tags            10. Конфигурация: Конфигурации по-умолчанию
// @Produce         json
// @Param           indicator          query    int  true  "ID индикатора устройства"
// @Param           device_component   query    int  false "ID компонента устройства"
// @Success         201  {object}  dto.DefaultConfigurationDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configuration [post]
func CreateDefaultConfiguration(c *gin.Context) {
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
	var res *dto.DefaultConfigurationDto
	res, err = service.CreateDefaultConfiguration(c.Request.Context(), indID, dcID)
	if err != nil {
		logger.Error("Service error occurred while creating default configuration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetDefaultConfiguration возвращает конфигурацию по умолчанию по ID
// @Summary         Получить конфигурацию по умолчанию по ID
// @Tags            10. Конфигурация: Конфигурации по-умолчанию
// @Produce         json
// @Param           id   path      int  true  "ID конфигурации"
// @Success         200  {object}  dto.DefaultConfigurationDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configuration/{id} [get]
func GetDefaultConfiguration(c *gin.Context) {
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
	var res *dto.DefaultConfigurationDto
	res, err = service.GetDefaultConfigurationByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving default configuration %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Default configuration not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllDefaultConfigurations возвращает все конфигурации по умолчанию
// @Summary         Получить все конфигурации по умолчанию
// @Tags            10. Конфигурация: Конфигурации по-умолчанию
// @Produce         json
// @Success         200  {array}   dto.DefaultConfigurationDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configurations [get]
func GetAllDefaultConfigurations(c *gin.Context) {
	var res []dto.DefaultConfigurationDto
	var err error
	res, err = service.GetAllDefaultConfigurations(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while retrieving all default configurations: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateDefaultConfiguration обновляет конфигурацию по умолчанию по ID
// @Summary         Обновить конфигурацию по умолчанию по ID
// @Tags            10. Конфигурация: Конфигурации по-умолчанию
// @Produce         json
// @Param           id                 path     int  true  "ID конфигурации"
// @Param           indicator          query    int  true  "Новый ID индикатора устройства"
// @Param           device_component   query    int  false "Новый ID компонента устройства"
// @Success         200  {object}  dto.DefaultConfigurationDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configuration/{id} [put]
func UpdateDefaultConfiguration(c *gin.Context) {
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
	var res *dto.DefaultConfigurationDto
	res, err = service.UpdateDefaultConfiguration(c.Request.Context(), id, indID, dcID)
	if err != nil {
		logger.Error("Service error occurred while updating default configuration %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Default configuration not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteDefaultConfiguration удаляет конфигурацию по умолчанию по ID
// @Summary         Удалить конфигурацию по умолчанию по ID
// @Tags            10. Конфигурация: Конфигурации по-умолчанию
// @Param           id   path      int  true  "ID конфигурации"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configuration/{id} [delete]
func DeleteDefaultConfiguration(c *gin.Context) {
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
	found, err = service.DeleteDefaultConfiguration(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while deleting default configuration %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Default configuration not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// ChangeDefaultConfigurationDataHandler изменяет ID конфигурации по умолчанию
// @Summary         Изменить ID конфигурации по умолчанию
// @Tags            10. Конфигурация: Конфигурации по-умолчанию
// @Produce         json
// @Param           prevId path      int  true  "Предыдущий ID конфигурации по умолчанию"
// @Param           newId  path      int  true  "Новый ID конфигурации по умолчанию"
// @Success         200  "OK"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configuration/{prevId}/{newId} [patch]
func ChangeDefaultConfigurationDataHandler(c *gin.Context) {
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
	success, err = service.ChangeDefaultConfigurationData(c.Request.Context(), prevId, newId)
	if err != nil {
		logger.Error("Service error during changing threshold data from %d to %d: %v", prevId, newId, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation rejected by service"})
		return
	}
	c.Status(http.StatusOK)
}
