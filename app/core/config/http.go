package config

type http struct {
	Listen string `yaml:"listen" env:"GO_HTTP_LISTEN"`
	Enable string `yaml:"enable" env:"GO_HTTP_ENABLE"`
}
