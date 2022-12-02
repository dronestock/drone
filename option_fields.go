package drone

import (
	"github.com/goexl/gox"
)

var (
	_            = Fields
	_ execOption = (*optionFields)(nil)
)

type optionFields struct {
	fields gox.Fields[any]
}

// Fields 字段
func Fields(fields ...gox.Field[any]) *optionFields {
	return &optionFields{
		fields: fields,
	}
}

func (f *optionFields) applyExec(options *execOptions) {
	options.fields = f.fields
}
