package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetOids возвращает список OID по переданным фильтрам в формате JSON
// @Summary         Поиск OID по фильтрам
// @Description     Возвращает список OID на основе переданной нотации, MIB, названия производителя и параметров пагинации
// @Tags            4. Парсер: OID
// @Accept          json
// @Produce         json
// @Param           request  body     dto.OidRequestDto  true  "Фильтры для поиска OID"
// @Success         200      {array}  dto.OidDto
// @Failure         400      {object} map[string]string
// @Failure         500      {object} map[string]string
// @Router          /api/v1/oid [post]
func GetOids(c *gin.Context) {
	var req dto.OidRequestDto
	var err error
	var res []dto.OidDto
	err = c.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Failed to bind JSON request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
		return
	}
	res, err = service.GetOids(c.Request.Context(), req)
	if err != nil {
		logger.Errorf("Service error occurred while fetching OIDs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
