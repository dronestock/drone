package drone

import (
	`strings`
	`time`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
)

var _ Configuration = (*Config)(nil)

// Config 插件基础配置
type Config struct {
	// 是否启用默认配置
	Defaults bool `default:"${PLUGIN_DEFAULTS=${DEFAULTS=true}}"`
	// 是否显示详细信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`
	// 是否显示调试信息
	Debug bool `default:"${PLUGIN_DEBUG=${DEBUG=false}}"`

	// 是否重试
	Retry bool `default:"${PLUGIN_RETRY=${RETRY=true}}"`
	// 重试次数
	Counts int `default:"${PLUGIN_COUNTS=${COUNTS=5}}"`
	// 重试间隔
	Backoff time.Duration `default:"${PLUGIN_BACKOFF=${BACKOFF=5s}}"`
}

func (c *Config) Parse(to map[string]string, configs ...string) {
	for _, config := range configs {
		c.parse(config, c.put(to))
	}
}

func (c *Config) Parses(to map[string][]string, configs ...string) {
	for _, config := range configs {
		c.parse(config, c.puts(to))
	}
}

func (c *Config) Setup() (unset bool, err error) {
	unset = true
	err = nil

	return
}

func (c *Config) Fields() gox.Fields {
	return gox.Fields{
		field.Bool(`defaults`, c.Defaults),
		field.Bool(`verbose`, c.Verbose),
		field.Bool(`debug`, c.Debug),

		field.Bool(`retry`, c.Retry),
		field.Int(`counts`, c.Counts),
		field.Duration(`backoff`, c.Backoff),
	}
}

func (c *Config) Basic() *Config {
	return c
}

func (c *Config) parse(original string, put func(configs []string)) {
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

	return
}

func (c *Config) puts(cache map[string][]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			cache[strings.TrimSpace(configs[0])] = c.splits(value, `,`, `|`, `||`)
		}
	}
}

func (c *Config) put(cache map[string]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			cache[strings.TrimSpace(configs[0])] = value
		}

		return
	}
}

func (c *Config) splits(config string, seps ...string) (configs []string) {
	configs = []string{config}
	for _, sep := range seps {
		if strings.Contains(config, sep) {
			configs = strings.Split(config, sep)
			break
		}
	}

	return
}
