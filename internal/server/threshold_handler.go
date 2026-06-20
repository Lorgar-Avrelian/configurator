package server

import (
	_ "configurator/internal/model"
)

/*// CreateThreshold создает порог
// @Summary         Создать порог
// @Tags            12. Конфигурация: Пороги
// @Accept          json
// @Produce         json
// @Param           request body dto.ThresholdCreate true "Данные порога"
// @Success         201  {object}  model.Threshold "Возвращает раскрытую цепочку условий"
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/thresholds [post]
func CreateThreshold(c *gin.Context) {
	var input dto.ThresholdCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := dao.CreateThreshold(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res, _ := dao.GetDetailedThresholdByID(c.Request.Context(), id)
	c.JSON(http.StatusCreated, res)
}

// GetThreshold возвращает порог и всю цепочку связей
// @Summary         Получить порог по ID вместе со всей цепочкой родителей
// @Tags            12. Конфигурация: Пороги
// @Produce         json
// @Param           id   path      int  true  "ID Порога"
// @Success         200  {object}  model.Threshold
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/thresholds/{id} [get]
func GetThreshold(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID порога"})
		return
	}
	res, err := dao.GetDetailedThresholdByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Порог не найден"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllThresholds возвращает все пороги с их цепочками
// @Summary         Получить все пороги
// @Tags            12. Конфигурация: Пороги
// @Produce         json
// @Success         200  {array}   model.Threshold
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/thresholds [get]
func GetAllThresholds(c *gin.Context) {
	res, err := dao.GetAllThresholdsExpanded(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateThreshold обновляет порог по ID
// @Summary         Обновить порог по ID
// @Tags            12. Конфигурация: Пороги
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Порога"
// @Param           request body dto.ThresholdUpdate true "Новые данные"
// @Success         200  {object}  model.Threshold
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/thresholds/{id} [put]
func UpdateThreshold(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID порога"})
		return
	}
	var input dto.ThresholdUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedID, err := dao.UpdateThreshold(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if updatedID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Порог не найден для обновления"})
		return
	}
	res, _ := dao.GetDetailedThresholdByID(c.Request.Context(), updatedID)
	c.JSON(http.StatusOK, res)
}

// DeleteThreshold удаляет порог по ID
// @Summary         Удалить порог по ID
// @Tags            12. Конфигурация: Пороги
// @Param           id   path      int  true  "ID Порога"
// @Success         204  "No Content"
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/thresholds/{id} [delete]
func DeleteThreshold(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	found, err := dao.DeleteThreshold(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Порог не найден"})
		return
	}
	c.Status(http.StatusNoContent)
}
*/
