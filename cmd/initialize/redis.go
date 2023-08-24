package initialize

import (
	"github.com/zzj-custom/pkg/pRedis"
	"image/internal/global"
)

func InitRedis() {
	rds := global.GVA_CONFIG.Redis
	if err := pRedis.InitMultiPools(rds); err != nil {
		panic("redis初始化失败")
	}

	// 获取默认的redis
	defaultPool, err := pRedis.Pool()
	if err != nil {
		panic("redis初始化失败")
	}
	global.GVA_REDIS = defaultPool
}

func ReleaseRedis() {
	rds := global.GVA_REDIS
	_ = rds.Close()
}
