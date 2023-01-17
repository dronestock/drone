package drone

import (
	"fmt"
	"os"
	"strings"

	"github.com/goexl/env"
)

func (b *Base) env(key string) (value string) {
	if strings.HasPrefix(key, envPrefix) || strings.HasPrefix(key, pluginEnvPrefix) {
		value = os.Getenv(key)
	} else if value = env.Get(b.droneEnv(key)); "" != value {
		return
	} else if value = env.Get(key); "" != value {
		return
	}

	return
}

func (b *Base) droneEnv(env string) string {
	return fmt.Sprintf("%s%s", pluginEnvPrefix, env)
}
