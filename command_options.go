package drone

import (
	"github.com/goexl/gox"
)

type commandOptions struct {
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
