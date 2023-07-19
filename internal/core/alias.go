package core

type Alias struct {
	Name  string
	Value string
}

func NewAlias(name string, value string) *Alias {
	return &Alias{
		Name:  name,
		Value: value,
	}
}
