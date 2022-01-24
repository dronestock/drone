package drone

var (
	_            = Name
	_ option     = (*optionName)(nil)
	_ stepOption = (*optionName)(nil)
	_ execOption = (*optionName)(nil)
)

type optionName struct {
	name string
}

// Name 配置插件名称
func Name(name string) *optionName {
	return &optionName{
		name: name,
	}
}

func (n *optionName) apply(options *options) {
	options.name = n.name
}

func (n *optionName) applyStep(options *stepOptions) {
	options.name = n.name
}

func (n *optionName) applyExec(options *execOptions) {
	options.name = n.name
}
