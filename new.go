package drone

import (
	"github.com/dronestock/drone/internal/plugin"
)

var _ = New

func New(constructor plugin.Constructor) (bootstrap *plugin.Bootstrap) {
	return plugin.New(constructor)
}
