package drone

type plugin interface {
	// Config 加载配置
	Config() (config Config)

	// Steps 插件运行步骤
	Steps() []*Step
}
