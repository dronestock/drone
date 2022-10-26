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

func parseAliases(aliases ...*alias) (err error) {
	for _, _alias := range aliases {
		config := os.Getenv(_alias.name)
		if `` == config {
			config = os.Getenv(dronePluginEnv(_alias.name))
		}
		if err = os.Setenv(_alias.value, config); nil != err {
			return
		}
		if err = os.Setenv(dronePluginEnv(_alias.value), config); nil != err {
			return
		}
	}

	return
}

func dronePluginEnv(env string) string {
	return fmt.Sprintf(`%s%s`, pluginEnvPrefix, env)
}
