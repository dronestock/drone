package drone

import (
	"os"
)

func (b *bootstrap) parseAliases() (err error) {
	for _, _alias := range b.aliases {
		config := os.Getenv(_alias.name)
		if "" == config {
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
