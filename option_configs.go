package drone

import (
	`strings`
)

var (
	_        = Configs
	_ option = (*optionConfigs)(nil)
)

type optionConfigs struct {
	configs []string
}

// Configs 配置组
func Configs(configs ...string) *optionConfigs {
	return &optionConfigs{
		configs: configs,
	}
}

func (c *optionConfigs) apply(options *options) {
	for _, config := range c.configs {
		options.configs = append(options.configs, strings.ToUpper(config))
	}
}
