package initialize

import (
	"github.com/gin-gonic/gin"
	"image/cmd/initialize/router"
	"image/internal/middleware"
	"image/internal/response"
)

func InitRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(gin.Recovery(), middleware.DefaultLogger())

	// 健康监测
	engine.GET("/heath", func(context *gin.Context) {
		response.Ok(context)
	})

	// 初始化路由
	routerGroup := engine.Group("/v1")
	{
		router.GroupApp.ImageRouter.Init(routerGroup)
	}

	return engine
}
