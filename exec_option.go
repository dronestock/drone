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
		args   []string
		dir    string
		fields gox.Fields
	}
)

func defaultExecOptions() *execOptions {
	return &execOptions{}
}
