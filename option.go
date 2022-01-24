package drone

var _ = NewOptions

type (
	option interface {
		apply(options *options)
	}

	options struct {
		name    string
		configs []string
		aliases []*alias
	}
)

// NewOptions 创建选项
func NewOptions(options ...option) []option {
	return options
}

func defaultOptions() *options {
	return &options{}
}
