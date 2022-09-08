package drone

import (
	"github.com/goexl/gox"
)

var _ = NewExecOptions

type (
	execOption interface {
		applyExec(options *execOptions)
	}

	execOptions struct {
		name         string
		args         []any
		environments []string
		dir          string
		pwe          bool
		async        bool
		fields       gox.Fields

		collectors []*collector
		checkers   []*checker
	}
)

// NewExecOptions 创建运行选项
func NewExecOptions(options ...execOption) []execOption {
	return options
}

func defaultExecOptions(pwe bool) *execOptions {
	return &execOptions{
		pwe: pwe,
	}
}
