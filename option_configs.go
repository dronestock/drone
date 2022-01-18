package drone

var (
	_        = NewConfigs
	_ option = (*optionConfigs)(nil)
)

type optionConfigs struct {
	configs []string
}

// NewConfigs 配置组
func NewConfigs(configs ...string) *optionConfigs {
	return &optionConfigs{
		configs: configs,
	}
}

func (c *optionConfigs) apply(options *options) {
	options.configs = c.configs
}
