package drone

import (
	"fmt"
)

func (b *bootstrap) droneEnv(env string) string {
	return fmt.Sprintf("%s%s", prefixPluginEnv, env)
}
