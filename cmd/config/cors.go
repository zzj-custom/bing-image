package config

type CORS struct {
	Mode      string          `json:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `json:"whitelist" yaml:"whitelist"`
}

type CORSWhitelist struct {
	AllowOrigin      string `json:"allow-origin" yaml:"allow-origin" mapstructure:"allow-origin"`
	AllowMethods     string `json:"allow-methods" yaml:"allow-methods" mapstructure:"allow-methods"`
	AllowHeaders     string `json:"allow-headers" yaml:"allow-headers" mapstructure:"allow-headers"`
	ExposeHeaders    string `json:"expose-headers" yaml:"expose-headers" mapstructure:"expose-headers"`
	AllowCredentials bool   `json:"allow-credentials" yaml:"allow-credentials" mapstructure:"allow-credentials"`
}
