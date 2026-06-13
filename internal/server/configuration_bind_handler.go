package server

import (
	"configurator/internal/dao"
	_ "configurator/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BindConfigThreshold связывает рабочую конфигурацию с порогом
// @Summary         Привязать порог к рабочей конфигурации
// @Tags            14. Конфигурация: Привязка порога к рабочей конфигурации
// @Accept          json
// @Produce         json
// @Param           request body dto.BindParamRequest true "ID конфигурации и ID порога (используем BindParamRequest для совместимости структуры)"
// @Success         200  {object}  map[string]string
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/configurations/bind [post]
func BindConfigThreshold(c *gin.Context) {
	var input struct {
		ConfigurationID int64 `json:"configuration_id" binding:"required"`
		ThresholdID     int64 `json:"threshold_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := dao.BindConfigThreshold(c.Request.Context(), input.ConfigurationID, input.ThresholdID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Порог успешно привязан к рабочей конфигурации"})
}

// UnbindConfigThreshold разрывает связь рабочей конфигурации с порогом
// @Summary         Удалить связь рабочей конфигурации с порогом
// @Tags            14. Конфигурация: Привязка порога к рабочей конфигурации
// @Param           configurationId path      int  true  "ID Рабочей конфигурации"
// @Param           thresholdId     path      int  true  "ID Порога"
// @Success         204  "No Content"
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/configurations/bind/{configurationId}/{thresholdId} [delete]
func UnbindConfigThreshold(c *gin.Context) {
	cfgID, _ := strconv.ParseInt(c.Param("configurationId"), 10, 64)
	tID, _ := strconv.ParseInt(c.Param("thresholdId"), 10, 64)
	found, err := dao.UnbindConfigThreshold(c.Request.Context(), cfgID, tID)
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
