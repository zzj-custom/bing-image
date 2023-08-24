package global

import (
	"github.com/gomodule/redigo/redis"
	"github.com/panjf2000/ants/v2"
	"github.com/zzj-custom/pkg/pRedis"
	"image/cmd/config"
	"image/internal/cron"
	"sync"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB                  *gorm.DB
	GVA_DB_List             map[string]*gorm.DB
	GVA_REDIS               *redis.Pool
	GVA_CONFIG              config.Server
	GVA_VP                  *viper.Viper
	GVA_LOG                 *zap.Logger
	GVA_Timer               = cron.NewTimerTask()
	GVA_Concurrency_Control = &singleflight.Group{}
	GVA_ANTS_POOL           *ants.Pool

	lock sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DB_List[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DB_List[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

// GetGlobalRedisByName 切换redis
func GetGlobalRedisByName(name string) (*redis.Pool, error) {
	return pRedis.Pool(name)
}
