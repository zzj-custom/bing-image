package config

type System struct {
	Env           string `json:"env" yaml:"env" mapstructure:"env"`                                  // 环境值
	Addr          int    `json:"addr" yaml:"addr" mapstructure:"addr"`                               // 端口值
	OssType       string `json:"oss-type" yaml:"oss-type" mapstructure:"oss-type"`                   // Oss类型
	UseMultipoint bool   `json:"use-multipoint" yaml:"use-multipoint" mapstructure:"use-multipoint"` // 多点登录拦截
	DefaultDB     string `json:"default-db" yaml:"default-db" mapstructure:"default-db"`             // 默认mysql和redis的库名
	LimitCountIP  int    `json:"iplimit-count" yaml:"iplimit-count" mapstructure:"limit-count"`
	LimitTimeIP   int    `json:"iplimit-time" yaml:"iplimit-time" mapstructure:"limit-time"`
	RouterPrefix  string `json:"router-prefix" yaml:"router-prefix" mapstructure:"router-prefix"`
}
