package drone

import (
	"fmt"
)

var (
	_               = Env
	_               = Environment
	_ commandOption = (*optionEnvironment)(nil)
	_ execOption    = (*optionEnvironment)(nil)
)

type optionEnvironment struct {
	key   string
	value string
}

// Env 环境变量
func Env(key string, value string) *optionEnvironment {
	return Environment(key, value)
}

// Environment 环境变量
func Environment(key string, value string) *optionEnvironment {
	return &optionEnvironment{
		key:   key,
		value: value,
	}
}

func (e *optionEnvironment) applyCommand(options *commandOptions) {
	options.environments = append(options.environments, fmt.Sprintf(environmentFormatter, e.key, e.value))
}

func (e *optionEnvironment) applyExec(options *execOptions) {
	options.environments = append(options.environments, fmt.Sprintf(environmentFormatter, e.key, e.value))
}
