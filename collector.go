package drone

type collector struct {
	mode collectorMode
	args any
}

func newCollector(mode collectorMode, args any) *collector {
	return &collector{
		mode: mode,
		args: args,
	}
}
