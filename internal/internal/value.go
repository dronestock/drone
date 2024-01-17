package internal

type Value struct {
	content string
}

func NewValue(content string) *Value {
	return &Value{
		content: content,
	}
}

func (v *Value) String() string {
	return v.content
}
