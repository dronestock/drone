package drone

import (
	"fmt"
	"os"
	"strings"

	"github.com/goexl/env"
)

func envGetter(key string) (value string) {
	if strings.HasPrefix(key, envPrefix) || strings.HasPrefix(key, pluginEnvPrefix) {
		value = os.Getenv(key)
	} else if value = env.Get(dronePluginEnv(key)); `` != value {
		return
	} else if value = env.Get(key); `` != value {
		return
	}

	return
}

func dronePluginEnv(env string) string {
	return fmt.Sprintf(`%s%s`, pluginEnvPrefix, env)
}
