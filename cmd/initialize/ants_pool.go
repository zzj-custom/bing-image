package initialize

import (
	"github.com/zzj-custom/pkg/pAntsPool"
	"go.uber.org/zap"
	"image/internal/global"
)

func InitAntsPool() {
	size := global.GVA_CONFIG.AntsPool.Size

	antsPool, err := pAntsPool.InitAsyncTaskPool(size)
	if err != nil {
		global.GVA_LOG.Error("初始化ants_pool失败", zap.Int("size", size))
	}
	global.GVA_ANTS_POOL = antsPool
}
