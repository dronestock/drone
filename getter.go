package drone

import (
	"os"
	"strings"

	"github.com/goexl/env"
)

type getter struct {
	bootstrap *bootstrap
}

func newGetter(bootstrap *bootstrap) *getter {
	return &getter{
		bootstrap: bootstrap,
	}
}

func (g *getter) Get(key string) (value string) {
	if strings.HasPrefix(key, envPrefix) || strings.HasPrefix(key, pluginEnvPrefix) {
		value = os.Getenv(key)
	} else if value = env.Get(g.bootstrap.droneEnv(key)); "" != value {
		return
	} else if value = env.Get(key); "" != value {
		return
	}

	return
}
