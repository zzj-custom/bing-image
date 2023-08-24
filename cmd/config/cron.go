package config

type Cron struct {
	Start       bool              `json:"start" yaml:"start" mapstructure:"start"`                      // 是否启用
	TaskList    map[string]string `json:"task-list" yaml:"task-list" mapstructure:"task-list"`          // CRON表达式
	WithSeconds bool              `json:"with-seconds" yaml:"with-seconds" mapstructure:"with-seconds"` // 是否精确到秒
}
