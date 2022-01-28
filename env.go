package drone

import (
	`fmt`
	`os`
	`strings`
)

func parseAliases(aliases ...*alias) (err error) {
	for _, _alias := range aliases {
		config := os.Getenv(_alias.name)
		if `` == config {
			config = os.Getenv(droneConfigName(_alias.name))
		}
		if err = os.Setenv(_alias.value, config); nil != err {
			return
		}
		if err = os.Setenv(droneConfigName(_alias.value), config); nil != err {
			return
		}
	}

	return
}

func parseConfigs(envs ...string) (err error) {
	for _, env := range envs {
		if err = parseStrings(env); nil != err {
			return
		}
	}

	return
}

func parseStrings(env string) (err error) {
	if values := parseValues(env); `` != values {
		err = setEnv(env, values)
	}
	if values := parseValues(droneConfigName(env)); `` != values {
		err = setEnv(env, values)
	}

	return
}

func parseValues(from string) (to string) {
	values := strings.Split(os.Getenv(from), `,`)
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
	to = fmt.Sprintf(`[%s]`, strings.Join(converts, `,`))

	return
}

func setEnv(env string, value string) (err error) {
	if err = os.Setenv(env, value); nil != err {
		return
	}
	err = os.Setenv(droneConfigName(env), value)

	return
}

func droneConfigName(env string) string {
	return fmt.Sprintf(`PLUGIN_%s`, env)
}
