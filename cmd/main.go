package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"image/cmd/initialize"
	"image/internal/core"
	"image/internal/global"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// 初始化gin模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化viper
	global.GVA_VP = core.Viper()

	// 初始化zap
	global.GVA_LOG = core.Zap()
	zap.ReplaceGlobals(global.GVA_LOG)

	// 初始化mysql
	initialize.InitMysql()
	defer initialize.ReleaseMysql()

	// 初始化redis
	initialize.InitRedis()
	defer initialize.ReleaseRedis()

	// 初始化定时任务
	initialize.InitCron()

	// 初始化antPool
	initialize.InitAntsPool()

	// 初始化服务
	core.RunServer(initialize.InitRouter())
}
