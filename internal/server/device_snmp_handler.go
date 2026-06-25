package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetWorkingConfiguration возвращает объединенную рабочую конфигурацию устройства
// @Summary         Получить рабочую конфигурацию устройства
// @Tags            14. Просмотр: Рабочая конфигурация
// @Produce         json
// @Param           host   query     string  true  "IP-адрес или хост устройства"
// @Param           port   query     int     true  "Порт устройства"
// @Success         200    {object}  dto.DeviceSnmpDto
// @Failure         400    {object}  map[string]string
// @Failure         404    {object}  map[string]string
// @Failure         500    {object}  map[string]string
// @Router          /api/v1/config/working [get]
func GetWorkingConfiguration(c *gin.Context) {
	var host string
	host = c.Query("host")
	if host == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Host parameter is required"})
		return
	}
	var portStr string
	portStr = c.Query("port")
	if portStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Port parameter is required"})
		return
	}
	var parsedInt int64
	var err error
	parsedInt, err = strconv.ParseInt(portStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid port format"})
		return
	}
	var portInt int32
	portInt = int32(parsedInt)
	if 1 > portInt {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Port must be positive"})
		return
	}
	var res *dto.DeviceSnmpDto
	res, err = service.GetWorkingConfiguration(c.Request.Context(), host, portInt)
	if err != nil {
		logger.Error("Service error occurred while retrieving working configuration: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Working configuration not found for specified host and port"})
		return
	}
	c.JSON(http.StatusOK, res)
}
