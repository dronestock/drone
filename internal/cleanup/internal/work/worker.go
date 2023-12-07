package work

import (
	"github.com/goexl/log"
)

type Worker interface {
	Work(logger log.Logger) (err error)
}
