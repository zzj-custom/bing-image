package config

type Crawler struct {
	BingImage *BingImage `json:"bing-image" yaml:"bing-image" mapstructure:"bing-image"`
}

type BingImage struct {
	URI    string `json:"uri" yaml:"uri" mapstructure:"uri"`
	Format string `json:"format" yaml:"format" mapstructure:"format"` // 非必填，json配置为js,默认返回xml 返回数据格式，json和xml,
	Idx    int    `json:"idx" yaml:"idx" mapstructure:"idx"`          // 非必填 请求图片截止天数, 0 今天, -1 截止中明天 （预准备的）1 截止至昨天，类推（目前最多获取到7天前的图片）
	N      int    `json:"n" yaml:"n" mapstructure:"n"`                // 1-8 返回请求数量，目前最多一次获取8张
	MKT    string `json:"mkt" yaml:"mkt" mapstructure:"mkt"`          // 非必填 地区
	Size   string `json:"size" yaml:"size" mapstructure:"size"`       // 非必填 默认1920×1080  1366×768 1280×768 1024×768 800×600 800×480 768×1280 720×1280 640×480 480×800 400×240 320×240 240×320
	Mbl    int    `json:"mbl" yaml:"mbl" mapstructure:"mbl"`          // 是否显示位置，例如：zh-CN
}
