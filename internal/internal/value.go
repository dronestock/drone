package internal

import (
	"strconv"
	"strings"
	"time"

	"github.com/dronestock/drone/internal/internal/constant"
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

func (v *Value) Slices() (result []string) {
	items := strings.Split(v.content, constant.Common)
	if 0 != len(items) {
		result = make([]string, 0, len(items))
	}
	for _, item := range items {
		finally := strings.TrimSpace(item)
		if "" != finally {
			result = append(result, finally)
		}
	}

	return
}
