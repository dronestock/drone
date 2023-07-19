package internal

import (
	"fmt"
)

func DroneEnv(env string) string {
	return fmt.Sprintf("%s%s", PrefixPluginEnv, env)
}
