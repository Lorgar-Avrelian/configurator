package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllParamResults возвращает все сохранённые значения параметров
// @Summary         Получить все сохранённые значения параметров
// @Tags            15. Просмотр: Значения параметров
// @Produce         json
// @Success         200  {array}   dto.ParamResultDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param-results [get]
func GetAllParamResults(c *gin.Context) {
	var res []dto.ParamResultDto
	var err error
	res, err = service.GetAllParamResults(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while retrieving all param results: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetParamResultsByFilter возвращает сохранённые значения параметров, отфильтрованные по переданным критериям
// @Summary         Получить отфильтрованные сохранённые значения параметров
// @Tags            15. Просмотр: Значения параметров
// @Produce         json
// @Param           request body   dto.ParamResultGetDto  true  "Фильтры параметров"
// @Success         200  {array}   dto.ParamResultDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param-result [post]
func GetParamResultsByFilter(c *gin.Context) {
	var input dto.ParamResultGetDto
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during param results body filtering: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res []dto.ParamResultDto
	res, err = service.GetParamResultsByFilter(c.Request.Context(), input)
	if err != nil {
		logger.Error("Service error occurred while filtering param results: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
