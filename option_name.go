package drone

var (
	_        = Name
	_ option = (*optionName)(nil)
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
