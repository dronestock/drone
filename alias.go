package drone

type alias struct {
	name  string
	value string
}

func newAlias(name string, value string) *alias {
	return &alias{
		name:  name,
		value: value,
	}
}
