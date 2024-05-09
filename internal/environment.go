package internal

import (
	"fmt"

	"github.com/dronestock/drone/internal/internal/constant"
)

func DroneEnvironment(environment string) string {
	return fmt.Sprintf("%s%s", constant.PrefixPluginEnvironment, environment)
}

func CIEnvironment(environment string) string {
	return fmt.Sprintf("%s%s", constant.PrefixCIEnvironment, environment)
}
