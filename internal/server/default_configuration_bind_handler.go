package server

import (
	"configurator/internal/dao"
	_ "configurator/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BindDefaultConfigThreshold связывает конфигурацию по умолчанию с порогом
// @Summary         Привязать порог к конфигурации по умолчанию
// @Tags            13. Конфигурация: Привязка порога к конфигурации по умолчанию
// @Accept          json
// @Produce         json
// @Param           request body dto.BindParamRequest true "ID конфигурации и ID порога"
// @Success         200  {object}  map[string]string
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configurations/bind [post]
func BindDefaultConfigThreshold(c *gin.Context) {
	var input struct {
		DefaultConfigurationID int64 `json:"default_configuration_id" binding:"required"`
		ThresholdID            int64 `json:"threshold_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.BindDefaultConfigThreshold(c.Request.Context(), input.DefaultConfigurationID, input.ThresholdID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Порог успешно привязан к дефолтной конфигурации"})
}

// UnbindDefaultConfigThreshold разрывает связь конфигурации по умолчанию с порогом
// @Summary         Удалить связь конфигурации по умолчанию с порогом
// @Tags            13. Конфигурация: Привязка порога к конфигурации по умолчанию
// @Param           defaultConfigurationId path      int  true  "ID Дефолтной конфигурации"
// @Param           thresholdId            path      int  true  "ID Порога"
// @Success         204  "No Content"
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/default-configurations/bind/{defaultConfigurationId}/{thresholdId} [delete]
func UnbindDefaultConfigThreshold(c *gin.Context) {
	defCfgID, _ := strconv.ParseInt(c.Param("defaultConfigurationId"), 10, 64)
	tID, _ := strconv.ParseInt(c.Param("thresholdId"), 10, 64)
	found, err := dao.UnbindDefaultConfigThreshold(c.Request.Context(), defCfgID, tID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Связь не найдена"})
		return
	}
	c.Status(http.StatusNoContent)
}
