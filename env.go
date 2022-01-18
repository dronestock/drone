package drone

import (
	`fmt`
	`os`
	`strings`
)

func parseConfigs(envs ...string) (err error) {
	for _, env := range envs {
		if err = parseStrings(env); nil != err {
			return
		}
	}

	return
}

func parseStrings(env string) (err error) {
	if err = parseValues(env); nil != err {
		return
	}
	err = parseValues(fmt.Sprintf(`PLUGIN_%s`, env))

	return
}

func parseValues(env string) (err error) {
	values := strings.Split(os.Getenv(env), `,`)
	converts := make([]string, 0, len(values))
	for _, value := range values {
		if `` == value {
			continue
		}
		converts = append(converts, fmt.Sprintf(`"%s"`, value))
	}

	if 0 == len(converts) {
		return
	}
	if err = os.Setenv(env, fmt.Sprintf(`[%s]`, strings.Join(converts, `,`))); nil != err {
		return
	}

	return
}
