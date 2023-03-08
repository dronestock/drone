package drone

import (
	"github.com/goexl/gox"
)

type commandParams struct {
	name         string
	args         []any
	environments []string
	dir          string
	pwe          bool
	async        bool
	fields       gox.Fields[any]

	collectors []*collector
	checkers   []*checker
}

func newCommandParams(pwe bool) *commandParams {
	return &commandParams{
		args: make([]any, 0, 4),
		pwe:  pwe,
	}
}
