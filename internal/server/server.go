package server

import (
	"configurator/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	gin.SetMode(gin.ReleaseMode)
	var r *gin.Engine
	r = gin.New()
	r.Use(gin.Recovery())
	r.Static("/docs", "./docs")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/docs/swagger.json")))
	var v1 *gin.RouterGroup
	v1 = r.Group("/api/v1")
	{
		v1.GET("/components", GetAllComponents)
		var component *gin.RouterGroup
		component = v1.Group("/component")
		{
			component.POST("", CreateComponent)
			component.GET("/search", SearchComponents)
			component.GET("/:id", GetComponent)
			component.PUT("/:id", UpdateComponent)
			component.DELETE("/:id", DeleteComponent)
		}
		v1.GET("/params", GetAllParams)
		var param *gin.RouterGroup
		param = v1.Group("/param")
		{
			param.POST("", CreateParam)
			param.GET("/unattached", GetUnattachedParams)
			param.GET("/search", SearchParams)
			param.GET("/search/:id", GetComponentsByParam)
			param.GET("/:id", GetParam)
			param.PUT("/:id", UpdateParam)
			param.DELETE("/:id", DeleteParam)
		}
		var relation *gin.RouterGroup
		relation = v1.Group("/relation")
		{
			relation.POST("", BindParam)
			relation.DELETE("/:componentId/:paramId", UnbindParam)
		}
		var oid *gin.RouterGroup
		oid = v1.Group("/oid")
		{
			oid.GET("", GetOidsByExactNotation)
			oid.GET("/prefix", GetOidsByPrefixNotation)
			oid.GET("/mib", GetOidsByMib)
			oid.GET("/vendor", GetOidsByVendor)
			oid.GET("/exact", GetOidsByDotterMibAndVendor)
		}
		v1.GET("/indicator/devices", GetAllIndicators)
		var indicatorDevice *gin.RouterGroup
		indicatorDevice = v1.Group("/indicator/device")
		{
			indicatorDevice.POST("", CreateIndicator)
			indicatorDevice.GET("/:id", GetIndicator)
			indicatorDevice.PUT("/:id", UpdateIndicator)
			indicatorDevice.DELETE("/:id", DeleteIndicator)
		}
		v1.GET("/indicator/params", GetAllParamIndicators)
		var indicatorParam *gin.RouterGroup
		indicatorParam = v1.Group("/indicator/param")
		{
			indicatorParam.POST("", CreateParamIndicator)
			indicatorParam.GET("/:id", GetParamIndicator)
			indicatorParam.PUT("/:id", UpdateParamIndicator)
			indicatorParam.DELETE("/:id", DeleteParamIndicator)
		}
		v1.GET("/mappings", GetAllMappings)
		var mapping *gin.RouterGroup
		mapping = v1.Group("/mapping")
		{
			mapping.POST("", CreateMapping)
			mapping.GET("/:id", GetMapping)
			mapping.PUT("/:id", UpdateMapping)
			mapping.DELETE("/:id", DeleteMapping)
		}
	}
	return &Server{router: r}
}

func (s *Server) Run() error {
	var cfg *config.Config
	cfg = config.Get()
	var addr string
	addr = fmt.Sprintf(":%d", cfg.Server.Port)
	var err error
	err = s.router.Run(addr)
	return err
}
