package server

import (
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOidsByExactNotation возвращает OID по точному совпадению dotter notation
// @Summary         Поиск OID по точной dotter notation
// @Tags            4. Парсер: OID
// @Produce         json
// @Param           notation query    string  true  "Точная dotter notation"
// @Success         200      {array}  dto.OidDto
// @Failure         400      {object} map[string]string
// @Failure         500      {object} map[string]string
// @Router          /api/v1/oid [get]
func GetOidsByExactNotation(c *gin.Context) {
	var notation string
	var res []dto.OidDto
	var err error
	notation = c.Query("notation")
	if notation == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'notation' не может быть пустым"})
		return
	}
	res, err = service.GetOidsByExactNotation(c.Request.Context(), notation)
	if err != nil {
		logger.Error("Service error occurred while fetching exact OID: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetOidsByPrefixNotation возвращает OID по префиксу с пагинацией и сортировкой в БД
// @Summary         Поиск OID по префиксу с пагинацией
// @Description     Возвращает отсортированный список OID (по 100 на страницу)
// @Tags            4. Парсер: OID
// @Produce         json
// @Param           prefix   query    string  true  "Префикс dotter_notation"
// @Param           page     query    int     false "Номер страницы (по умолчанию: 1)"
// @Success         200      {array}  dto.OidDto
// @Failure         400      {object} map[string]string
// @Failure         500      {object} map[string]string
// @Router          /api/v1/oid/prefix [get]
func GetOidsByPrefixNotation(c *gin.Context) {
	var prefix string
	var page int
	var res []dto.OidDto
	var err error
	prefix = c.Query("prefix")
	if prefix == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'prefix' не может быть пустым"})
		return
	}
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	if 1 > page {
		page = 1
	}
	res, err = service.GetOidsByPrefixNotation(c.Request.Context(), prefix, page)
	if err != nil {
		logger.Error("Service error occurred while fetching OIDs by prefix: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetOidsByMib возвращает OID по названию MIB
// @Summary         Получить OID по названию MIB
// @Tags            4. Парсер: OID
// @Produce         json
// @Param           name     query    string  true  "Название MIB"
// @Success         200      {array}  dto.OidDto
// @Failure         400      {object} map[string]string
// @Failure         500      {object} map[string]string
// @Router          /api/v1/oid/mib [get]
func GetOidsByMib(c *gin.Context) {
	var name string
	var res []dto.OidDto
	var err error
	name = c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 'name' не может быть пустым"})
		return
	}
	res, err = service.GetOidsByMib(c.Request.Context(), name)
	if err != nil {
		logger.Error("Service error occurred while fetching OIDs by MIB: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetOidsByVendor возвращает OID производителя с пагинацией и сортировкой в БД
// @Summary         Получить OID по производителю с пагинацией
// @Tags            4. Парсер: OID
// @Produce         json
// @Param           vendor   query    string  false "Имя вендора или его директория"
// @Param           page     query    int     false "Номер страницы (дефолт: 1)"
// @Success         200      {array}  dto.OidDto
// @Failure         500      {object} map[string]string
// @Router          /api/v1/oid/vendor [get]
func GetOidsByVendor(c *gin.Context) {
	var vendor string
	var page int
	var res []dto.OidDto
	var err error
	vendor = c.Query("vendor")
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	if 1 > page {
		page = 1
	}
	res, err = service.GetOidsByVendor(c.Request.Context(), vendor, page)
	if err != nil {
		logger.Error("Service error occurred while fetching OIDs by vendor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetOidsByDotterMibAndVendor возвращает OID по нотации, MIB и производителю
// @Summary         Поиск OID по dotter notation, названию MIB и производителю
// @Tags            4. Парсер: OID
// @Produce         json
// @Param           notation query    string  true  "Точная dotter notation"
// @Param           mib      query    string  true  "Точное название MIB"
// @Param           vendor   query    string  false "Имя или директория вендора"
// @Success         200      {array}  dto.OidDto
// @Failure         400      {object} map[string]string
// @Failure         500      {object} map[string]string
// @Router          /api/v1/oid/exact [get]
func GetOidsByDotterMibAndVendor(c *gin.Context) {
	var notation string
	var mibName string
	var vendorIdent string
	var hasVendor bool
	var vendorPtr *string
	var res []dto.OidDto
	var err error
	notation = c.Query("notation")
	mibName = c.Query("mib")
	vendorIdent, hasVendor = c.GetQuery("vendor")
	if notation == "" || mibName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters 'notation' и 'mib' couldn't be empty"})
		return
	}
	vendorPtr = nil
	if hasVendor && vendorIdent != "" {
		vendorPtr = &vendorIdent
	}
	res, err = service.GetOidsByDotterMibAndVendor(c.Request.Context(), notation, mibName, vendorPtr)
	if err != nil {
		logger.Error("Service error occurred while fetching OIDs by dotter, MIB and vendor: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
