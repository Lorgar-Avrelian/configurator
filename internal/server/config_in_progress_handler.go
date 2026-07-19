package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllConfigInProcess возвращает все данные устройств, находящихся в процессе конфигурирования
// @Summary         Получить все данные устройств, находящихся в процессе конфигурирования
// @Tags            13. Просмотр: В процессе конфигурирования
// @Produce         json
// @Success         200  {array}   dto.ConfigInProcessDto
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/config/in-progress [get]
func GetAllConfigInProcess(c *gin.Context) {
	var res []dto.ConfigInProcessDto
	var err error
	res, err = service.GetAllConfigInProcess(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while fetching all in-progress configs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No configurations in progress found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// SearchConfigInProcess ищет данные устройств, находящихся в процессе конфигурирования
// @Summary         Поиск данных устройств, находящихся в процессе конфигурирования
// @Tags            13. Просмотр: В процессе конфигурирования
// @Produce         json
// @Param           host   query    string false "Host устройства"
// @Param           port   query    int    false "Порт устройства"
// @Success         200  {array}   dto.ConfigInProcessDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/config/in-progress/search [post]
func SearchConfigInProcess(c *gin.Context) {
	var hostRaw string
	var portRaw string
	hostRaw = c.Query("host")
	portRaw = c.Query("port")
	var hostPtr *string
	var portPtr *int32
	hostPtr = nil
	portPtr = nil
	if hostRaw != "" {
		hostPtr = &hostRaw
	}
	if portRaw != "" {
		var portVal int64
		var err error
		portVal, err = strconv.ParseInt(portRaw, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid port format"})
			return
		}
		if 1 > portVal {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Port must be positive"})
			return
		}
		var finalPort int32
		finalPort = int32(portVal)
		portPtr = &finalPort
	}
	var res []dto.ConfigInProcessDto
	var err error
	res, err = service.SearchConfigInProcess(c.Request.Context(), hostPtr, portPtr)
	if err != nil {
		logger.Error("Service error occurred while searching in-progress configs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No configurations found matching search criteria"})
		return
	}
	c.JSON(http.StatusOK, res)
}
