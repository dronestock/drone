package drone

// Plugin 插件接口，任何插件都需要实现这个接口
type Plugin interface {
	// Config 加载配置
	Config() (config Config)

	// Steps 插件运行步骤
	Steps() []*Step
}
