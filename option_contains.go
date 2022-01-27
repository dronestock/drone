package drone

var (
	_            = Contains
	_ execOption = (*optionContains)(nil)
)

type optionContains struct {
	contains string
}

// Contains 检查是否包含字符串
func Contains(contains string) *optionContains {
	return &optionContains{
		contains: contains,
	}
}

func (c *optionContains) applyExec(options *execOptions) {
	options.checkerMode = checkerModeContains
	options.checkerArgs = c.contains
}
