package drone

var (
	_            = Break
	_            = Continue
	_ stepOption = (*optionBreak)(nil)
)

type optionBreak struct {
	_break bool
}

// Break 遇到错误中断执行
func Break() *optionBreak {
	return &optionBreak{
		_break: true,
	}
}

// Continue 遇到错误继续执行
func Continue() *optionBreak {
	return &optionBreak{
		_break: false,
	}
}

func (b *optionBreak) applyStep(options *stepOptions) {
	options._break = b._break
}
