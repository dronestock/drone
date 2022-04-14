package drone

import (
	"github.com/goexl/gox"
)

type (
	commandOption interface {
		applyCommand(options *commandOptions)
	}

	commandOptions struct {
		name         string
		environments []string
		dir          string
		async        bool
		fields       gox.Fields

		collectors []*collector
		checkers   []*checker
	}
)

func defaultCommandOptions() *commandOptions {
	return &commandOptions{}
}
