package drone

// Plugin 插件
type Plugin interface {
	// Config 加载配置
	Config() (config Config)

	// Steps 插件运行步骤
	Steps() []*Step
}
