package drone

var (
	_            = Equal
	_ execOption = (*optionEqual)(nil)
)

type optionEqual struct {
	equal string
}

// Equal 检查是否字符串相等
func Equal(equal string) *optionEqual {
	return &optionEqual{
		equal: equal,
	}
}

func (c *optionEqual) applyExec(options *execOptions) {
	options.checkers = append(options.checkers, newChecker(checkerModeEqual, c.equal))
}
