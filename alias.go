package drone

var _ = NewAlias

type alias struct {
	name  string
	value string
}

// NewAlias 创建别名
func NewAlias(name string, value string) *alias {
	return &alias{
		name:  name,
		value: value,
	}
}
