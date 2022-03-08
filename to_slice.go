package drone

import (
	`fmt`
	`reflect`
	`strconv`
	`strings`
)

func toSlice(from string, field reflect.StructField) (to string, err error) {
	to = from
	left := `[`
	right := `]`
	if reflect.Slice != field.Type.Kind() || (strings.Contains(from, left) && strings.Contains(from, right)) {
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
