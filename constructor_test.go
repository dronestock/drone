package drone

type plugin struct {
	Base
}

func newPlugin() Plugin {
	return new(plugin)
}

func (p *plugin) Steps() Steps {
	return Steps{
		NewDebugStep(),
	}
}

func (p *plugin) Config() Config {
	return p
}
