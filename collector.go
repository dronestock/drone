package drone

type collector struct {
	mode collectorMode
	args interface{}
}

func newCollector(mode collectorMode, args interface{}) *collector {
	return &collector{
		mode: mode,
		args: args,
	}
}
