package internal

import (
	"strconv"
	"time"
)

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

func (v *Value) Timestamp() string {
	return v.Time().Format(time.RFC3339)
}

func (v *Value) Time() (timestamp time.Time) {
	if value, pie := strconv.ParseInt(v.content, 10, 64); nil == pie {
		timestamp = time.Unix(value, 0)
	} else {
		timestamp = time.Now()
	}

	return
}
