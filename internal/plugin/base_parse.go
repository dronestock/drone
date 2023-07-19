package plugin

import (
	"strings"
)

func (b *Base) Parse(to map[string]string, configs ...string) {
	for _, config := range configs {
		b.parse(config, b.put(to))
	}
}

func (b *Base) Parses(to map[string][]string, configs ...string) {
	for _, config := range configs {
		b.parse(config, b.puts(to))
	}
}

func (b *Base) parse(original string, put func(configs []string)) {
	var _configs []string
	defer func() {
		put(_configs)
	}()

	if _configs = strings.Split(original, "@"); 2 <= len(_configs) {
		return
	}
	if _configs = strings.Split(original, "=>"); 2 <= len(_configs) {
		return
	}
	if _configs = strings.Split(original, "->"); 2 <= len(_configs) {
		return
	}
	if _configs = strings.Split(original, " "); 2 <= len(_configs) {
		return
	}
}

func (b *Base) puts(cache map[string][]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			cache[strings.TrimSpace(configs[0])] = b.splits(value, `,`, `|`, `||`)
		}
	}
}

func (b *Base) put(cache map[string]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			cache[strings.TrimSpace(configs[0])] = value
		}
	}
}

func (b *Base) splits(config string, seps ...string) (configs []string) {
	configs = []string{config}
	for _, sep := range seps {
		if strings.Contains(config, sep) {
			configs = strings.Split(config, sep)

			break
		}
	}

	return
}
