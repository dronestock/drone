package drone

import (
	"github.com/goexl/gox"
)

var (
	_               = Fields
	_ commandOption = (*optionFields)(nil)
	_ execOption    = (*optionFields)(nil)
)

type optionFields struct {
	fields []gox.Field
}

// Fields 字段
func Fields(fields ...gox.Field) *optionFields {
	return &optionFields{
		fields: fields,
	}
}

func (f *optionFields) applyCommand(options *commandOptions) {
	options.fields = f.fields
}

func (f *optionFields) applyExec(options *execOptions) {
	options.fields = f.fields
}
