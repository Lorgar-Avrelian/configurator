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
			param.GET("/:id", GetParam)
			param.PUT("/:id", UpdateParam)
			param.DELETE("/:id", DeleteParam)
		}
		var relation *gin.RouterGroup
		relation = v1.Group("/relation")
		{
			relation.POST("", BindParam)
			relation.DELETE("/:componentId/:paramId", UnbindParam)
			relation.GET("/search/:id", GetComponentsByParam)
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
