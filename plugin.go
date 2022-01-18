package drone

type plugin interface {
	// Configuration 加载配置
	Configuration() (configuration Configuration)

	// Steps 插件运行步骤
	Steps() []*Step
}
