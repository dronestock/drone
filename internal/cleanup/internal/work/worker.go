package work

import (
	"github.com/goexl/simaqian"
)

type Worker interface {
	Work(logger simaqian.Logger) (err error)
}
