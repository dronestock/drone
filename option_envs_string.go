package drone

var (
	_            = StringEnv
	_            = StringEnvs
	_ execOption = (*optionStringEnvs)(nil)
)

type optionStringEnvs struct {
	envs []string
}

// StringEnv 环境变量
func StringEnv(env string) *optionStringEnvs {
	return &optionStringEnvs{
		envs: []string{env},
	}
}

// StringEnvs 环境变量
func StringEnvs(envs ...string) *optionStringEnvs {
	return &optionStringEnvs{
		envs: envs,
	}
}

func (e *optionStringEnvs) applyExec(options *execOptions) {
	options.envs = e.envs
}
