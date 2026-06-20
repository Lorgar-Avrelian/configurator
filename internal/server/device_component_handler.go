package server

import (
	_ "configurator/internal/model"
)

/*// CreateDeviceComponent создает новый узел состава устройства
// @Summary         Создать узел состава устройства
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Accept          json
// @Produce         json
// @Param           request body dto.DeviceComponentCreate true "Данные узла"
// @Success         201  {object}  model.DeviceComponent
// @Router          /api/v1/device-components [post]
func CreateDeviceComponent(c *gin.Context) {
	var input dto.DeviceComponentCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при создании узла устройства: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.CreateDeviceComponent(c.Request.Context(), input)
	if err != nil {
		logger.Error("Ошибка DAO при создании узла устройства: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetDeviceComponent возвращает узел состава по ID вместе с маппингами
// @Summary         Получить узел состава по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Produce         json
// @Param           id   path      int  true  "ID Узла"
// @Success         200  {object}  model.DeviceComponent
// @Router          /api/v1/device-components/{id} [get]
func GetDeviceComponent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID узла"})
		return
	}
	res, err := dao.GetDeviceComponentByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при получении узла %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Узел устройства не найден"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllDeviceComponents возвращает все узлы с их маппингами параметров
// @Summary         Получить всю структуру подчиненности устройств
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Produce         json
// @Success         200  {array}   model.DeviceComponent
// @Router          /api/v1/device-components [get]
func GetAllDeviceComponents(c *gin.Context) {
	res, err := dao.GetAllDeviceComponents(c.Request.Context())
	if err != nil {
		logger.Error("Ошибка DAO при выгрузке структуры подчиненности устройств: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateDeviceComponent обновляет метаданные узла
// @Summary         Обновить узел состава по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID Узла"
// @Param           request body dto.DeviceComponentUpdate true "Новые данные"
// @Success         200  {object}  model.DeviceComponent
// @Router          /api/v1/device-components/{id} [put]
func UpdateDeviceComponent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID узла"})
		return
	}
	var input dto.DeviceComponentUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Ошибка валидации при обновлении узла %d: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := dao.UpdateDeviceComponent(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Ошибка DAO при обновлении узла %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Узел не найден для обновления"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteDeviceComponent удаляет узел из дерева подчиненности
// @Summary         Удалить узел состава по ID
// @Tags            8. Конфигурация: Структура компонентов устройства
// @Param           id   path      int  true  "ID Узла"
// @Success         204  "No Content"
// @Router          /api/v1/device-components/{id} [delete]
func DeleteDeviceComponent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID узла"})
		return
	}
	found, err := dao.DeleteDeviceComponent(c.Request.Context(), id)
	if err != nil {
		logger.Error("Ошибка DAO при удалении узла %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Узел не найден в системе"})
		return
	}
	c.Status(http.StatusNoContent)
}
*/
