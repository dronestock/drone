package drone

import (
	`fmt`
	`reflect`
	`strconv`
	`strings`
)

func toSlice(from string, field reflect.StructField) (to string, err error) {
	to = from
	if !canConvert(from, field) {
		return
	}

	separator := `,`
	values := strings.Split(from, separator)
	finals := make([]string, 0, len(values))
	for _, value := range values {
		if _, parseErr := strconv.ParseFloat(value, 64); nil == parseErr {
			finals = append(finals, value)
		} else {
			finals = append(finals, fmt.Sprintf(`"%s"`, value))
		}
	}
	to = fmt.Sprintf(`[%s]`, strings.Join(finals, separator))

	return
}

func canConvert(from string, field reflect.StructField) bool {
	return `` != strings.TrimSpace(from) && // 不能是空字符串
		reflect.Slice == field.Type.Kind() && // 只能是列表
		!(strings.HasPrefix(from, jsonArrayStart) && strings.HasSuffix(from, jsonArrayEnd)) // 不能是数组
}
