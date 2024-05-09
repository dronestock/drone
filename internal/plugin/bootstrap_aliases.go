package plugin

import (
	"os"

	"github.com/dronestock/drone/internal"
)

func (b *Bootstrap) parseAliases() (err error) {
	for _, alias := range b.aliases {
		config := os.Getenv(alias.Name)
		if "" == config {
			config = os.Getenv(internal.CIEnvironment(alias.Name))
		}
		if "" == config {
			config = os.Getenv(internal.DroneEnvironment(alias.Name))
		}

		if soe := os.Setenv(alias.Value, config); nil != soe {
			err = soe
		} else if sde := os.Setenv(internal.DroneEnvironment(alias.Value), config); nil != sde {
			err = sde
		} else if sce := os.Setenv(internal.CIEnvironment(alias.Value), config); nil != sce {
			err = sce
		}
	}

	return
}
