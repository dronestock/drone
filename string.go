package drone

var (
	_ collectorBuilder = (*_string)(nil)
	_                  = String
)

type _string struct {
	output *string
}

// String 字符串输出
func String(output *string) *_string {
	return &_string{
		output: output,
	}
}

func (s *_string) collector() *collector {
	return newCollector(collectorModeString, s.output)
}
