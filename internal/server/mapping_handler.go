package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	_ "configurator/internal/model"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateMapping создаёт новое сопоставление
// @Summary         Создать сопоставление
// @Tags            7. Конфигурация: Сопоставления параметров
// @Accept          json
// @Produce         json
// @Param           request body dto.MappingCreateDto true "Данные сопоставления"
// @Success         201  {object}  dto.MappingDto
// @Failure         400  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/mapping [post]
func CreateMapping(c *gin.Context) {
	var input dto.MappingCreateDto
	var err error
	var res *dto.MappingDto
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during mapping creation: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	res, err = service.CreateMapping(c.Request.Context(), input)
	if err != nil {
		logger.Error("Service error occurred while creating mapping: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create mapping"})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetMapping возвращает сопоставление по ID вместе с иерархией дочерних элементов
// @Summary         Получить сопоставление по ID
// @Tags            7. Конфигурация: Сопоставления параметров
// @Produce         json
// @Param           id   path      int  true  "ID сопоставления"
// @Success         200  {object}  dto.MappingDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/mapping/{id} [get]
func GetMapping(c *gin.Context) {
	var id int64
	var err error
	var res *dto.MappingDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mapping ID format"})
		return
	}
	res, err = service.GetMappingByID(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving mapping %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve mapping"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mapping not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetAllMappings возвращает все сопоставления в виде иерархического дерева
// @Summary         Получить все сопоставления
// @Tags            7. Конфигурация: Сопоставления параметров
// @Produce         json
// @Success         200  {array}   dto.MappingDto
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/mappings [get]
func GetAllMappings(c *gin.Context) {
	var res []dto.MappingDto
	var err error
	res, err = service.GetAllMappings(c.Request.Context())
	if err != nil {
		logger.Error("Service error occurred while retrieving all mappings: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve all mappings"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateMapping обновляет сопоставление по ID
// @Summary         Обновить сопоставление по ID
// @Tags            7. Конфигурация: Сопоставления параметров
// @Accept          json
// @Produce         json
// @Param           id      path      int  true  "ID сопоставления"
// @Param           request body dto.MappingCreateDto true "Новые данные"
// @Success         200  {object}  dto.MappingDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/mapping/{id} [put]
func UpdateMapping(c *gin.Context) {
	var id int64
	var err error
	var input dto.MappingCreateDto
	var res *dto.MappingDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mapping ID format"})
		return
	}
	if err = c.ShouldBindJSON(&input); err != nil {
		logger.Warn("Validation failed during mapping update for ID %d: %v", id, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body format"})
		return
	}
	res, err = service.UpdateMapping(c.Request.Context(), id, input)
	if err != nil {
		logger.Error("Service error occurred while updating mapping %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update mapping"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mapping not found for update"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteMapping удаляет сопоставление по ID
// @Summary         Удалить сопоставление по ID
// @Tags            7. Конфигурация: Сопоставления параметров
// @Param           id   path      int  true  "ID сопоставления"
// @Success         204  "No Content"
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/mapping/{id} [delete]
func DeleteMapping(c *gin.Context) {
	var id int64
	var err error
	var found bool
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mapping ID format"})
		return
	}
	found, err = service.DeleteMapping(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while deleting mapping %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete mapping"})
		return
	}
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mapping not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetMappingOwn возвращает сопоставление по ID без указания положения внутри дерева вложенности
// @Summary         Получить изолированное сопоставление по ID
// @Tags            7. Конфигурация: Сопоставления параметров
// @Produce         json
// @Param           id   path      int  true  "ID сопоставления"
// @Success         200  {object}  dto.MappingDto
// @Failure         400  {object}  map[string]string
// @Failure         404  {object}  map[string]string
// @Failure         500  {object}  map[string]string
// @Router          /api/v1/mapping/{id}/own [get]
func GetMappingOwn(c *gin.Context) {
	var id int64
	var err error
	var res *dto.MappingDto
	id, err = strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid mapping ID format"})
		return
	}
	res, err = service.GetMappingByIDOwn(c.Request.Context(), id)
	if err != nil {
		logger.Error("Service error occurred while retrieving single mapping %d: %v", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve mapping"})
		return
	}
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mapping not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}
