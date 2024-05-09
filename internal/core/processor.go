package core

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/dronestock/drone/internal/internal/constant"
)

type Processor struct{}

func NewProcessor() *Processor {
	return new(Processor)
}

func (b *Processor) Process(tag string, field reflect.StructField) (to string, err error) {
	to = tag
	if !b.canConvert(tag, field) {
		return
	}

	separator := ","
	values := strings.Split(tag, separator)
	finals := make([]string, 0, len(values))
	for _, value := range values {
		if _, parseErr := strconv.ParseFloat(value, 64); nil == parseErr {
			finals = append(finals, value)
		} else {
			finals = append(finals, fmt.Sprintf(`"%s"`, value))
		}
	}
	to = fmt.Sprintf("[%s]", strings.Join(finals, separator))

	return
}

func (b *Processor) canConvert(from string, field reflect.StructField) bool {
	return "" != strings.TrimSpace(from) && // 不能是空字符串
		reflect.Slice == field.Type.Kind() && // 只能是列表
		!(strings.HasPrefix(from, constant.JsonArrayStart) && strings.HasSuffix(from, constant.JsonArrayEnd)) // 不能是数组
}
