package drone

type plugin interface {
	// 加载配置
	configuration() (configuration configuration)

	// 插件运行步骤
	steps() []*step
}
