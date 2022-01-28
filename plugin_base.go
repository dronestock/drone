package drone

import (
	`strings`
	`time`

	`github.com/storezhang/gox`
	`github.com/storezhang/gox/field`
	`github.com/storezhang/simaqian`
)

var _ Config = (*PluginBase)(nil)

// PluginBase 插件基础配置
type PluginBase struct {
	simaqian.Logger

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

func (pb *PluginBase) Parse(to map[string]string, configs ...string) {
	for _, config := range configs {
		pb.parse(config, pb.put(to))
	}
}

func (pb *PluginBase) Parses(to map[string][]string, configs ...string) {
	for _, config := range configs {
		pb.parse(config, pb.puts(to))
	}
}

func (pb *PluginBase) Setup() (unset bool, err error) {
	unset = true
	err = nil

	return
}

func (pb *PluginBase) Fields() gox.Fields {
	return gox.Fields{
		field.Bool(`defaults`, pb.Defaults),
		field.Bool(`verbose`, pb.Verbose),
		field.Bool(`debug`, pb.Debug),

		field.Bool(`retry`, pb.Retry),
		field.Int(`counts`, pb.Counts),
		field.Duration(`backoff`, pb.Backoff),
	}
}

func (pb *PluginBase) Base() *PluginBase {
	return pb
}

func (pb *PluginBase) parse(original string, put func(configs []string)) {
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

func (pb *PluginBase) puts(cache map[string][]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			if nil == cache {
				cache = make(map[string][]string)
			}
			cache[strings.TrimSpace(configs[0])] = pb.splits(value, `,`, `|`, `||`)
		}
	}
}

func (pb *PluginBase) put(cache map[string]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			if nil == cache {
				cache = make(map[string]string)
			}
			cache[strings.TrimSpace(configs[0])] = value
		}

		return
	}
}

func (pb *PluginBase) splits(config string, seps ...string) (configs []string) {
	configs = []string{config}
	for _, sep := range seps {
		if strings.Contains(config, sep) {
			configs = strings.Split(config, sep)
			break
		}
	}

	return
}
