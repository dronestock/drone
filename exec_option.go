package drone

var _ = NewExecOptions

type (
	execOption interface {
		applyExec(options *execOptions)
	}

	execOptions struct {
	}
)

// NewExecOptions 创建运行选项
func NewExecOptions(options ...execOption) []execOption {
	return options
}

func defaultExecOptions(pwe bool) *execOptions {
	return &execOptions{
		pwe: pwe,
	}
}
