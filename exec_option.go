package drone

import (
	`github.com/storezhang/gox`
)

type (
	execOption interface {
		applyExec(options *execOptions)
	}

	execOptions struct {
		name   string
		args   []interface{}
		envs   []string
		dir    string
		async  bool
		fields gox.Fields

		checkerMode checkerMode
		checkerArgs interface{}
	}
)

func defaultExecOptions() *execOptions {
	return &execOptions{}
}
