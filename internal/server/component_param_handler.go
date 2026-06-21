package server

import (
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"configurator/internal/dto"

	"github.com/gin-gonic/gin"
)

// BindParam связывает компонент и параметр
// @Summary         Связать компонент с параметром
// @Description     Добавляет запись в таблицу public.component_param и возвращает состояние компонента
// @Tags            3. Модельный каталог: Связи
// @Produce         json
// @Param           componentId path      int  true  "ID Компонента"
// @Param           paramId     path      int  true  "ID Параметра"
// @Success         200  {object}  dto.ComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/link/component-param/{componentId}/{paramId} [post]
func BindParam(c *gin.Context) {
	var compID int64
	var paramID int64
	var err error
	var res *dto.ComponentDto
	compID, _ = strconv.ParseInt(c.Param("componentId"), 10, 64)
	paramID, _ = strconv.ParseInt(c.Param("paramId"), 10, 64)
	if 1 > compID || 1 > paramID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid component ID or parameter ID format"})
		return
	}
	res, err = service.BindParam(c.Request.Context(), compID, paramID)
	if err != nil {
		logger.Error("Service error occurred while binding parameter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to bind parameter to component"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Component not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UnbindParam разрывает связь между компонентом и параметром
// @Summary         Удалить связь компонента с параметром
// @Description     Удаляет запись из таблицы public.component_param и возвращает обновленное состояние компонента
// @Tags            3. Модельный каталог: Связи
// @Produce         json
// @Param           componentId path      int  true  "ID Компонента"
// @Param           paramId     path      int  true  "ID Параметра"
// @Success         200  {object}  dto.ComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/link/component-param/{componentId}/{paramId} [delete]
func UnbindParam(c *gin.Context) {
	var compID int64
	var paramID int64
	var err error
	var res *dto.ComponentDto
	compID, _ = strconv.ParseInt(c.Param("componentId"), 10, 64)
	paramID, _ = strconv.ParseInt(c.Param("paramId"), 10, 64)
	if 1 > compID || 1 > paramID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid component ID or parameter ID format"})
		return
	}
	res, err = service.UnbindParam(c.Request.Context(), compID, paramID)
	if err != nil {
		logger.Error("Service error occurred while removing parameter from component: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unbind parameter from component"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relationship or component not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}
