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
// @Description     Добавляет запись в таблицу public.component_param
// @Tags            3. Модельный каталог: Связи
// @Accept          json
// @Produce         json
// @Param           request body dto.ComponentParamLinkDto true "ID сущностей для связывания"
// @Success         200  {object}  map[string]string "Сообщение об успешном связывании"
// @Failure         400  {object}  map[string]string "Ошибка валидации JSON"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/relation [post]
func BindParam(c *gin.Context) {
	var input dto.ComponentParamLinkDto
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during parameter binding: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = service.BindParam(c.Request.Context(), input)
	if err != nil {
		logger.Error("Service error occurred while binding parameter: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Parameter successfully bound to component"})
}

// UnbindParam разрывает связь между компонентом и параметром
// @Summary         Удалить связь компонента с параметром
// @Description     Удаляет запись из таблицы public.component_param
// @Tags            3. Модельный каталог: Связи
// @Param           componentId path      int  true  "ID Компонента"
// @Param           paramId     path      int  true  "ID Параметра"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string "Неверный формат ID"
// @Failure         404  {object}  map[string]string "Связь не найдена"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/relation/{componentId}/{paramId} [delete]
func UnbindParam(c *gin.Context) {
	var componentID int64
	var paramID int64
	var errComp error
	var errParam error
	componentID, errComp = strconv.ParseInt(c.Param("componentId"), 10, 64)
	paramID, errParam = strconv.ParseInt(c.Param("paramId"), 10, 64)
	if errComp != nil || errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid component ID or parameter ID format"})
		return
	}
	var found bool
	var err error
	found, err = service.UnbindParam(c.Request.Context(), componentID, paramID)
	if err != nil {
		logger.Error("Service error occurred while removing parameter from component: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Relationship between entities not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetComponentsByParam возвращает список компонентов, к которым напрямую привязан параметр
// @Summary         Получить компоненты по ID параметра
// @Description     Возвращает список компонентов, для которых параметр с данным ID является их собственным, а не унаследованным
// @Tags            3. Модельный каталог: Связи
// @Produce         json
// @Param           id   path      int  true  "ID Параметра"
// @Success         200  {array}   dto.ComponentDto
// @Failure         400  {object}  map[string]string "Неверный формат ID"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/relation/search/{id} [get]
func GetComponentsByParam(c *gin.Context) {
	var id int64
	var err error
	id = 0
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID format"})
		return
	}
	var res []dto.ComponentDto
	res, err = service.GetComponentsByDirectParamID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while fetching components for direct param %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
