package server

import (
	"filler/internal/config"
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()

	r.Static("/docs", "./docs")

	url := ginSwagger.URL("/docs/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1 := r.Group("/api/v1")
	{
		components := v1.Group("/components")
		{
			components.POST("", CreateComponent)
			components.GET("", GetAllComponents)
			components.GET("/search", SearchComponents)
			components.GET("/:id", GetComponent)
			components.PUT("/:id", UpdateComponent)
			components.DELETE("/:id", DeleteComponent)
		}

		params := v1.Group("/params")
		{
			params.POST("", CreateParam)
			params.GET("/unattached", GetUnattachedParams)
			params.GET("/search", SearchParams)
			params.GET("/:id", GetParam)
			params.PUT("/:id", UpdateParam)
			params.DELETE("/:id", DeleteParam)
		}

		relations := v1.Group("/relations")
		{
			relations.POST("", BindParam)
			relations.DELETE("/:componentId/:paramId", UnbindParam)
		}
	}

	return &Server{router: r}
}

func (s *Server) Run() error {
	cfg := config.Get()
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	return s.router.Run(addr)
}
