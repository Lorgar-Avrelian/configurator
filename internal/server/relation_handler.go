package server

import (
	"configurator/internal/logger"
	"net/http"
	"strconv"

	"configurator/internal/dao"
	"configurator/internal/dto"

	"github.com/gin-gonic/gin"
)

// BindParam связывает компонент и параметр
// @Summary         Связать компонент с параметром
// @Description     Добавляет запись в таблицу public.component_param
// @Tags            3. Модельный каталог: Связи
// @Accept          json
// @Produce         json
// @Param           request body dto.BindParamRequest true "ID сущностей для связывания"
// @Success         200  {object}  map[string]string "Сообщение об успешном связывании"
// @Failure         400  {object}  map[string]string "Ошибка валидации JSON"
// @Failure         500  {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/relations [post]
func BindParam(c *gin.Context) {
	var input dto.BindParamRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при связывании сущностей: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := dao.BindParam(c.Request.Context(), input)
	if err != nil {
		logger.Error("Ошибка DAO при связывании компонента %d и параметра %d: %v", input.ComponentID, input.ParamID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Параметр успешно привязан к компоненту"})
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
// @Router          /api/v1/relations/{componentId}/{paramId} [delete]
func UnbindParam(c *gin.Context) {
	componentID, errComp := strconv.ParseInt(c.Param("componentId"), 10, 64)
	paramID, errParam := strconv.ParseInt(c.Param("paramId"), 10, 64)
	if errComp != nil || errParam != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные ID компонентов или параметров"})
		return
	}

	found, err := dao.UnbindParam(c.Request.Context(), componentID, paramID)
	if err != nil {
		logger.Error("Ошибка DAO при удалении связи между компонентом %d и параметром %d: %v", componentID, paramID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Связь между указанными сущностями не найдена"})
		return
	}
	c.Status(http.StatusNoContent)
}
