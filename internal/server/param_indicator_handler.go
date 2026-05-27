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

// CreateParamIndicator создает индикатор параметров
// @Summary         Создать индикатор параметров
// @Tags            6. Конфигурация: Индикаторы параметров
// @Accept          json
// @Produce         json
// @Param           request body dto.ParamIndicatorCreate true "Данные"
// @Success         201  {object}  model.ParamIndicator
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param-indicators [post]
func CreateParamIndicator(c *gin.Context) {
	var input dto.ParamIndicatorCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при создании индикатора параметров: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.CreateParamIndicator(c.Request.Context(), input)
	if err != nil {
		logger.Error("Ошибка DAO при создании индикатора параметров: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetParamIndicator возвращает индикатор параметров по ID
// @Summary         Получить индикатор параметров по ID
// @Tags            6. Конфигурация: Индикаторы параметров
// @Produce         json
// @Param           id   path      int  true  "ID Индикатора"
// @Success         200  {object}  model.ParamIndicator
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param-indicators/{id} [get]
func GetParamIndicator(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID индикатора параметров"})
		return
	}
	res, err := dao.GetParamIndicatorByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при получении индикатора параметров %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Индикатор параметров не найден"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllParamIndicators возвращает все индикаторы параметров
// @Summary         Получить все индикаторы параметров
// @Tags            6. Конфигурация: Индикаторы параметров
// @Produce         json
// @Success         200  {array}   model.ParamIndicator
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param-indicators [get]
func GetAllParamIndicators(c *gin.Context) {
	res, err := dao.GetAllParamIndicators(c.Request.Context())
	if err != nil {
		logger.Error("Ошибка DAO при выгрузке всех индикаторов параметров: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateParamIndicator обновляет индикатор параметров по ID
// @Summary         Обновить индикатор параметров по ID
// @Tags            6. Конфигурация: Индикаторы параметров
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Индикатора"
// @Param           request body dto.ParamIndicatorUpdate true "Новые данные"
// @Success         200  {object}  model.ParamIndicator
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param-indicators/{id} [put]
func UpdateParamIndicator(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID индикатора"})
		return
	}
	var input dto.ParamIndicatorUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при обновлении индикатора параметров %d: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.UpdateParamIndicator(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Ошибка DAO при обновлении индикатора параметров %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Индикатор параметров не найден для обновления"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteParamIndicator удаляет индикатор параметров по ID
// @Summary         Удалить индикатор параметров по ID
// @Tags            6. Конфигурация: Индикаторы параметров
// @Param           id   path      int  true  "ID Индикатора"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/param-indicators/{id} [delete]
func DeleteParamIndicator(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID индикатора"})
		return
	}
	found, err := dao.DeleteParamIndicator(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при удалении индикатора параметров %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Индикатор параметров не найден в системе"})
		return
	}
	c.Status(http.StatusNoContent)
}
