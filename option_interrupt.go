package drone

var (
	_            = Interrupt
	_            = Continue
	_ stepOption = (*optionInterrupt)(nil)
)

type optionInterrupt struct {
	interrupt bool
}

// Interrupt 遇到错误中断执行
func Interrupt() *optionInterrupt {
	return &optionInterrupt{
		interrupt: true,
	}
}

// Continue 遇到错误继续执行
func Continue() *optionInterrupt {
	return &optionInterrupt{
		interrupt: false,
	}
}

func (i *optionInterrupt) applyStep(options *stepOptions) {
	options.interrupt = i.interrupt
}
