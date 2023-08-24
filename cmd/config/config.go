package config

import (
	"github.com/zzj-custom/pkg/pAntsPool"
	"github.com/zzj-custom/pkg/pRedis"
	"image/pkg/pMysql"
)

type Server struct {
	Zap    Zap                       `json:"zap" yaml:"zap" mapstructure:""zap`
	Redis  []*pRedis.MultiDialConfig `json:"redis" yaml:"redis" mapstructure:"redis"`
	System System                    `json:"system" yaml:"system" mapstructure:"system"`
	// gorm
	Mysql map[string]*pMysql.Database `json:"mysql" yml:"mysql" mapstructure:"mysql"`
	// oss
	Local     Local     `json:"local" yaml:"local" mapstructure:"local"`
	Qiniu     Qiniu     `json:"qiniu" yaml:"qiniu" mapstructure:"qiniu"`
	AliyunOSS AliyunOSS `json:"aliyun-oss" yaml:"aliyun-oss" mapstructure:"aliyun-oss"`

	Cron Cron `json:"cron" yaml:"cron" mapstructure:"cron"`

	// 跨域配置
	Cors CORS `json:"cors" yaml:"cors" mapstructure:"cors"`

	// ants_pool
	AntsPool pAntsPool.Config `json:"ants-pool" yaml:"ant" mapstructure:"ants-pool"`

	// 爬虫数据配置
	Crawler Crawler `json:"crawler" yaml:"crawler" mapstructure:"crawler"`
}
