package router

import (
	"github.com/gin-gonic/gin"
	v1 "image/internal/api/v1"
)

type ImageRouter struct{}

func (s *ImageRouter) Init(Router *gin.RouterGroup) (R gin.IRoutes) {
	router := Router.Group("image")
	_ = v1.ApiGroupApp.Api
	{
		// router.POST("xx", api.xx)
	}
	return router
}
