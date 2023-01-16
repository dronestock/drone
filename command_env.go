package drone

import (
	"fmt"
)

func (cb *commandBuilder) Environment(key string, value string) *commandBuilder {
	cb.options.environments = append(cb.options.environments, fmt.Sprintf(environmentFormatter, key, value))

	return cb
}

func (cb *commandBuilder) Env(key string, value string) *commandBuilder {
	return cb.Environment(key, value)
}

func (cb *commandBuilder) StringEnvironment(env string) *commandBuilder {
	cb.options.environments = append(cb.options.environments, env)

	return cb
}

func (cb *commandBuilder) StringEnv(env string) *commandBuilder {
	return cb.StringEnvironment(env)
}

func (cb *commandBuilder) StringEnvironments(envs ...string) *commandBuilder {
	cb.options.environments = append(cb.options.environments, envs...)

	return cb
}

func (cb *commandBuilder) StringEnvs(envs ...string) *commandBuilder {
	return cb.StringEnvironments(envs...)
}
