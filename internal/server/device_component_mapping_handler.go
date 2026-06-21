package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BindDeviceMapping связывает составную часть устройства и сопоставление параметра
// @Summary         Связать составную часть устройства с сопоставлением параметра
// @Tags            9. Конфигурация: Связь компонентов устройства и сопоставлений параметров
// @Produce         json
// @Param           deviceComponentId path      int  true  "ID составной части"
// @Param           mappingId         path      int  true  "ID сопоставления"
// @Success         200  {object}  dto.DeviceComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/link/device-component-mapping/{deviceComponentId}/{mappingId} [post]
func BindDeviceMapping(c *gin.Context) {
	var dcID int64
	var mID int64
	var err error
	var res *dto.DeviceComponentDto
	dcID, _ = strconv.ParseInt(c.Param("deviceComponentId"), 10, 64)
	mID, _ = strconv.ParseInt(c.Param("mappingId"), 10, 64)
	if 1 > dcID || 1 > mID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IDs format"})
		return
	}
	res, err = service.BindDeviceMapping(c.Request.Context(), dcID, mID)
	if err != nil {
		logger.Error("Service error during binding: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to bind entities"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Device component not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UnbindDeviceMapping разрывает связь между составной частью устройства и сопоставлением параметра
// @Summary         Удалить связь составной части устройства с сопоставлением параметра
// @Tags            9. Конфигурация: Связь компонентов устройства и сопоставлений параметров
// @Produce         json
// @Param           deviceComponentId path      int  true  "ID составной части"
// @Param           mappingId         path      int  true  "ID сопоставления"
// @Success         200  {object}  dto.DeviceComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/link/device-component-mapping/{deviceComponentId}/{mappingId} [delete]
func UnbindDeviceMapping(c *gin.Context) {
	var dcID int64
	var mID int64
	var err error
	var res *dto.DeviceComponentDto
	dcID, _ = strconv.ParseInt(c.Param("deviceComponentId"), 10, 64)
	mID, _ = strconv.ParseInt(c.Param("mappingId"), 10, 64)
	if 1 > dcID || 1 > mID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid IDs format"})
		return
	}
	res, err = service.UnbindDeviceMapping(c.Request.Context(), dcID, mID)
	if err != nil {
		logger.Error("Service error during unbinding: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unbind entities"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relationship or device component not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}
