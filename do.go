package drone

import (
	`github.com/storezhang/simaqian`
)

type do func(logger simaqian.Logger) (undo bool, err error)
