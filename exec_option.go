package drone

import (
	`github.com/goexl/gox`
)

type (
	execOption interface {
		applyExec(options *execOptions)
	}

	execOptions struct {
		name         string
		args         []interface{}
		environments []string
		dir          string
		async        bool
		fields       gox.Fields

		checkerMode checkerMode
		checkerArgs interface{}
	}
)

func defaultExecOptions() *execOptions {
	return &execOptions{}
}
