package drone

var (
	_            = String
	_ execOption = (*optionString)(nil)
)

type optionString struct {
	output *string
}

// String 字符串输出
func String(output *string) *optionString {
	return &optionString{
		output: output,
	}
}

func (s *optionString) applyExec(options *execOptions) {
	options.collectors = append(options.collectors, newCollector(collectorModeString, s.output))
}
