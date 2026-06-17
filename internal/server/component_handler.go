package server

import (
	"configurator/internal/logger"
	"net/http"
	"strconv"

	"configurator/internal/dao"
	"configurator/internal/dto"
	_ "configurator/internal/model"

	"github.com/gin-gonic/gin"
)

// CreateComponent создает новый компонент
// @Summary         Создать компонент
// @Description     Создает новую запись составной части устройства. Если base_component равен null, создается базовый компонент.
// @Tags            1. Модельный каталог: Компоненты
// @Accept          json
// @Produce         json
// @Param           request body dto.ComponentCreate true "Данные компонента"
// @Success         201  {object}  model.Component
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/components [post]
func CreateComponent(c *gin.Context) {
	var input dto.ComponentCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при создании компонента: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.CreateComponent(c.Request.Context(), input)
	if err != nil {
		logger.Error("Ошибка DAO при создании компонента: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetComponent возвращает компонент со всеми его параметрами (включая наследуемые)
// @Summary         Получить компонент по ID
// @Description     Возвращает компонент и иерархически объединенный список его параметров из БД
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Param           id   path      int  true  "ID Компонента"
// @Success         200  {object}  model.Component
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/components/{id} [get]
func GetComponent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компонента"})
		return
	}
	res, err := dao.GetComponentByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при получении компонента %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Компонент не найден"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllComponents возвращает все компоненты с их параметрами
// @Summary         Получить все компоненты вместе с параметрами
// @Description     Возвращает список всех составных частей устройств вместе со всеми собственными и унаследованными параметрами
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Success         200  {array}   model.Component
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/components [get]
func GetAllComponents(c *gin.Context) {
	res, err := dao.GetAllComponentsWithParams(c.Request.Context())
	if err != nil {
		logger.Error("Ошибка DAO при получении списка всех компонентов: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// SearchComponents осуществляет полнотекстовый поиск компонентов по совпадению строки
// @Summary         Поиск компонентов по подстроке
// @Description     Ищет компоненты, у которых строка совпадает с title, name_en, name_ru, description_en или description_ru
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Param           query query     string true  "Строка поиска"
// @Success         200  {array}   model.Component
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/components/search [get]
func SearchComponents(c *gin.Context) {
	queryText := c.Query("query")
	if queryText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр поиска 'query' не может быть пустым"})
		return
	}
	res, err := dao.SearchComponents(c.Request.Context(), queryText)
	if err != nil {
		logger.Error("Ошибка DAO при поиске компонентов по запросу '%s': %v", queryText, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateComponent обновляет существующий компонент
// @Summary         Обновить компонент
// @Tags            1. Модельный каталог: Компоненты
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Компонента"
// @Param           request body dto.ComponentUpdate true "Новые данные"
// @Success         200  {object}  model.Component
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/components/{id} [put]
func UpdateComponent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компонента"})
		return
	}
	var input dto.ComponentUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при обновлении компонента %d: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.UpdateComponent(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Ошибка DAO при обновлении компонента %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Компонент не найден для обновления"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteComponent удаляет компонент рекурсивно со всеми его наследниками
// @Summary         Удалить компонент и всех его наследников
// @Description     Удаляет компонент по ID, а также автоматически стирает всю иерархию дочерних элементов
// @Tags            1. Модельный каталог: Компоненты
// @Param           id   path      int  true  "ID Компонента"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/components/{id} [delete]
func DeleteComponent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компонента"})
		return
	}
	found, err := dao.DeleteComponent(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при удалении компонента %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Компонент не найден"})
		return
	}
	c.Status(http.StatusNoContent)
}

/*// ShiftComponent сдвигает компонент вверх или вниз по списку ID
// @Summary         Переместить компонент в списке компонентов
// @Description     Меняет ID указанного компонента с соседним в зависимости от направления (up/down) с каскадным обновлением связей. Проверяет условие base_component > id.
// @Tags            1. Модельный каталог: Компоненты
// @Produce         json
// @Param           id        path      int     true  "ID компонента"
// @Param           target_id query     int     true  "Целевой ID"
// @Success         200       {object}  map[string]string "Сообщение об успешном сдвиге"
// @Failure         400       {object}  map[string]string "Неверный ID, нарушение иерархии или параметров"
// @Failure         404       {object}  map[string]string "Компонент или целевой ID не найден"
// @Failure         500       {object}  map[string]string "Внутренняя ошибка сервера"
// @Router          /api/v1/components/{id}/shift [patch]
func ShiftComponent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компонента"})
		return
	}
	targetID, err := strconv.ParseInt(c.Query("target_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный target_id"})
		return
	}
	moved, err := dao.ShiftComponent(c.Request.Context(), id, targetID)
	if err != nil {
		logger.Error("Ошибка при сдвиге компонента %d к %d: %v", id, targetID, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !moved {
		c.JSON(http.StatusNotFound, gin.H{"error": "Компонент или целевая позиция не найдены"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Компонент успешно сдвинут на целевую позицию"})
}*/
