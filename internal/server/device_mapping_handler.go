package server

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BindDeviceMapping связывает узловой компонент и сопоставление параметра
// @Summary         Связать узел устройства с сопоставлением параметра
// @Tags            9. Конфигурация: Связь компонентов устройства и сопоставлений параметров
// @Accept          json
// @Produce         json
// @Param           request body dto.BindDeviceMappingRequest true "Данные связывания"
// @Success         200  {object}  map[string]string
// @Router          /api/v1/device-components/bind [post]
func BindDeviceMapping(c *gin.Context) {
	var input dto.BindDeviceMappingRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при связывании маппинга: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := dao.BindDeviceMapping(c.Request.Context(), input)
	if err != nil {
		logger.Error("Ошибка DAO при связывании узла %d и маппинга %d: %v", input.DeviceComponentID, input.MappingID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Маппинг параметров успешно привязан к физическому узлу устройства"})
}

// UnbindDeviceMapping разрывает связь между узловым компонентом и сопоставлением параметра
// @Summary         Удалить связь узла устройства с сопоставлением параметра
// @Tags            9. Конфигурация: Связь компонентов устройства и сопоставлений параметров
// @Param           deviceComponentId path      int  true  "ID Узла"
// @Param           mappingId         path      int  true  "ID Маппинга"
// @Success         204  "No Content"
// @Router          /api/v1/device-components/bind/{deviceComponentId}/{mappingId} [delete]
func UnbindDeviceMapping(c *gin.Context) {
	dcID, errDc := strconv.ParseInt(c.Param("deviceComponentId"), 10, 64)
	mID, errM := strconv.ParseInt(c.Param("mappingId"), 10, 64)
	if errDc != nil || errM != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные ID узлов или маппингов"})
		return
	}
	found, err := dao.UnbindDeviceMapping(c.Request.Context(), dcID, mID)
	if err != nil {
		logger.Error("Ошибка DAO при разрыве связи узла %d и маппинга %d: %v", dcID, mID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Связь между указанными сущностями не найдена"})
		return
	}
	c.Status(http.StatusNoContent)
}
