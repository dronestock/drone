package drone

import (
	"fmt"
	"os"
)

func parseAliases(aliases ...*alias) (err error) {
	for _, _alias := range aliases {
		config := os.Getenv(_alias.name)
		if `` == config {
			config = os.Getenv(droneConfigName(_alias.name))
		}
		if err = os.Setenv(_alias.value, config); nil != err {
			return
		}
		if err = os.Setenv(droneConfigName(_alias.value), config); nil != err {
			return
		}
	}

	return
}

func droneConfigName(env string) string {
	return fmt.Sprintf(`PLUGIN_%s`, env)
}
