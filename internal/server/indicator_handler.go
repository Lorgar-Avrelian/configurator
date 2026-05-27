package server

import (
	"filler/internal/dao"
	"filler/internal/dto"
	"filler/internal/logger"
	"net/http"
	"strconv"

	_ "filler/internal/model"

	"github.com/gin-gonic/gin"
)

// CreateIndicator создает новую запись индикатора
// @Summary         Создать индикатор устройства
// @Tags            5. Конфигурация: Индикаторы устройств
// @Accept          json
// @Produce         json
// @Param           request body dto.DeviceIndicatorCreate true "Данные индикатора"
// @Success         201  {object}  model.DeviceIndicator
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicators [post]
func CreateIndicator(c *gin.Context) {
	var input dto.DeviceIndicatorCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при создании индикатора: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.CreateIndicator(c.Request.Context(), input)
	if err != nil {
		logger.Error("Ошибка DAO при создании индикатора: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetIndicator возвращает индикатор по ID
// @Summary         Получить индикатор по ID
// @Tags            5. Конфигурация: Индикаторы устройств
// @Produce         json
// @Param           id   path      int  true  "ID Индикатора"
// @Success         200  {object}  model.DeviceIndicator
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicators/{id} [get]
func GetIndicator(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID индикатора"})
		return
	}
	res, err := dao.GetIndicatorByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при получении индикатора %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Индикатор не найден"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllIndicators возвращает полный список записей таблицы
// @Summary         Получить все индикаторы устройств
// @Tags            5. Конфигурация: Индикаторы устройств
// @Produce         json
// @Success         200  {array}   model.DeviceIndicator
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicators [get]
func GetAllIndicators(c *gin.Context) {
	res, err := dao.GetAllIndicators(c.Request.Context())
	if err != nil {
		logger.Error("Ошибка DAO при выгрузке всех индикаторов: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateIndicator изменяет запись индикатора по ID
// @Summary         Обновить индикатор по ID
// @Tags            5. Конфигурация: Индикаторы устройств
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Индикатора"
// @Param           request body dto.DeviceIndicatorUpdate true "Новые данные"
// @Success         200  {object}  model.DeviceIndicator
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicators/{id} [put]
func UpdateIndicator(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID индикатора"})
		return
	}
	var input dto.DeviceIndicatorUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при обновлении индикатора %d: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.UpdateIndicator(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Ошибка DAO при обновлении индикатора %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Индикатор не найден для обновления"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteIndicator удаляет запись индикатора по ID
// @Summary         Удалить индикатор по ID
// @Tags            5. Конфигурация: Индикаторы устройств
// @Param           id   path      int  true  "ID Индикатора"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/indicators/{id} [delete]
func DeleteIndicator(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID индикатора"})
		return
	}
	found, err := dao.DeleteIndicator(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при удалении индикатора %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Индикатор не найден в системе"})
		return
	}
	c.Status(http.StatusNoContent)
}
