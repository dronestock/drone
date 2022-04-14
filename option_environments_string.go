package drone

var (
	_            = StringEnv
	_            = StringEnvironment
	_            = StringEnvs
	_            = StringEnvironments
	_ execOption = (*optionStringEnvironments)(nil)
)

type optionStringEnvironments struct {
	environments []string
}

// StringEnv 环境变量
func StringEnv(env string) *optionStringEnvironments {
	return StringEnvironment(env)
}

// StringEnvironment 环境变量
func StringEnvironment(environment string) *optionStringEnvironments {
	return &optionStringEnvironments{
		environments: []string{environment},
	}
}

// StringEnvs 环境变量列表
func StringEnvs(envs ...string) *optionStringEnvironments {
	return StringEnvironments(envs...)
}

// StringEnvironments 环境变量列表
func StringEnvironments(environments ...string) *optionStringEnvironments {
	return &optionStringEnvironments{
		environments: environments,
	}
}

func (se *optionStringEnvironments) applyExec(options *execOptions) {
	options.environments = append(options.environments, se.environments...)
}
