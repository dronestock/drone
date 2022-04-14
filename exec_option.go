package drone

type (
	execOption interface {
		applyExec(options *execOptions)
	}

	execOptions struct {
		*commandOptions

		args []interface{}
	}
)

func defaultExecOptions() *execOptions {
	return &execOptions{
		commandOptions: defaultCommandOptions(),
	}
}

func newExecOptions(options *commandOptions) *execOptions {
	return &execOptions{
		commandOptions: options,
	}
}
