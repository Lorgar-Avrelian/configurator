package server

import (
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"configurator/internal/dto"
	_ "configurator/internal/model"

	"github.com/gin-gonic/gin"
)

// CreateComponent создаёт новый компонент устройства
// @Summary         Создать компонент устройства
// @Tags            1. Модельный каталог: Компоненты
// @Accept          json
// @Produce         json
// @Param           request body dto.ComponentCreateDto true "Данные компонента"
// @Success         201  {object}  dto.ComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/component [post]
func CreateComponent(c *gin.Context) {
	var input dto.ComponentCreateDto
	var err error
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during component creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ComponentDto
	res, err = service.CreateComponent(c.Request.Context(), input)
	if err != nil {
		logger.Error("Service error occurred while creating component: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetComponent возвращает компонент по ID вместе с маппингами
// @Summary         Получить компонент по ID
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Param           id   path      int  true  "ID компонента"
// @Success         200  {object}  dto.ComponentDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/component/{id} [get]
func GetComponent(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid component ID format"})
		return
	}
	var res *dto.ComponentDto
	res, err = service.GetComponentByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving component %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Component not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateComponent обновляет метаданные компонент
// @Summary         Обновить компонент по ID
// @Tags            1. Модельный каталог: Компоненты
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID компонента"
// @Param           request body dto.ComponentUpdateDto true "Новые данные"
// @Success         200  {object}  dto.ComponentDto
// @Router          /api/v1/component/{id} [put]
func UpdateComponent(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid component ID format"})
		return
	}
	var input dto.ComponentUpdateDto
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during component update for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var res *dto.ComponentDto
	res, err = service.UpdateComponent(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Service error occurred while updating component %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Component not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteComponent удаляет компонент из дерева подчиненности
// @Summary         Удалить компонент по ID
// @Tags            1. Модельный каталог: Компоненты
// @Param           id   path      int  true  "ID компонента"
// @Success         204  "No Content"
// @Router          /api/v1/component/{id} [delete]
func DeleteComponent(c *gin.Context) {
	var id int64
	var err error
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid component ID format"})
		return
	}
	var found bool
	found, err = service.DeleteComponent(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while deleting component %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Component not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetAllComponents возвращает все компоненты с их параметрами
// @Summary         Получить всю структуру подчиненности устройств
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Success         200  {array}   dto.ComponentDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/components [get]
func GetAllComponents(c *gin.Context) {
	var res []dto.ComponentDto
	var err error
	res, err = service.GetAllComponents(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while retrieving all components: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// SearchComponents выполняет полнотекстовый поиск компонента
// @Summary         Поиск компонентов по строке
// @Description     Ищет компоненты, у которых title, name_en, name_ru, description_en или description_ru частично совпадают с переданной строкой
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Param           query query     string  true  "Строка поиска"
// @Success         200   {array}   dto.ComponentDto
// @Failure         400   {object}  map[string]string "Пустой запрос"
// @Failure         500   {object}  map[string]string "Ошибка базы данных"
// @Router          /api/v1/component/search [get]
func SearchComponents(c *gin.Context) {
	var queryText string
	queryText = c.Query("query")
	if queryText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query parameter cannot be empty"})
		return
	}
	var res []dto.ComponentDto
	var err error
	res, err = service.SearchComponents(c.Request.Context(), queryText)
	if err != nil {
		logger.Error("Service error occurred while searching components for query '%s': %v", queryText, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ChangeComponentDataHandler изменяет ID компонента
// @Summary         Изменить ID компонента
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Param           prevId path      int  true  "Предыдущий ID компонента"
// @Param           newId  path      int  true  "Новый ID компонента"
// @Success         200  "OK"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/component/{prevId}/{newId} [patch]
func ChangeComponentDataHandler(c *gin.Context) {
	var prevId int64
	var newId int64
	var err error
	prevId, err = strconv.ParseInt(c.Param("prevId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prevId format"})
		return
	}
	newId, err = strconv.ParseInt(c.Param("newId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid newId format"})
		return
	}
	var success bool
	success, err = service.ChangeComponentData(c.Request.Context(), prevId, newId)
	if err != nil {
		logger.Error("Service error during changing component data from %d to %d: %v", prevId, newId, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operation rejected by service"})
		return
	}
	c.Status(http.StatusOK)
}
