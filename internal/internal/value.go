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

func (v *Value) Timestamp() (stamp string) {
	if created, pie := strconv.ParseInt(v.content, 10, 64); nil == pie {
		stamp = time.Unix(created, 0).Format(time.RFC3339)
	} else {
		stamp = time.Now().Format(time.RFC3339)
	}

	return
}
