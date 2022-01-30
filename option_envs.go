package drone

var (
	_            = Envs
	_ execOption = (*optionEnvs)(nil)
)

type optionEnvs struct {
	envs []string
}

// Envs 参数
func Envs(envs ...string) *optionEnvs {
	return &optionEnvs{
		envs: envs,
	}
}

func (e *optionEnvs) applyExec(options *execOptions) {
	options.envs = e.envs
}
