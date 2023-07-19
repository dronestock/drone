package plugin

import (
	"os"

	"github.com/dronestock/drone/internal"
)

func (b *Bootstrap) parseAliases() (err error) {
	for _, alias := range b.aliases {
		config := os.Getenv(alias.Name)
		if "" == config {
			config = os.Getenv(internal.DroneEnv(alias.Name))
		}
		if err = os.Setenv(alias.Value, config); nil != err {
			return
		}
		if err = os.Setenv(internal.DroneEnv(alias.Value), config); nil != err {
			return
		}
	}

	return
}
