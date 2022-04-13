package drone

import (
	"strings"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

var _ Config = (*Plugin)(nil)

// Plugin 插件基础
type Plugin struct {
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

func (p *Plugin) Parse(to map[string]string, configs ...string) {
	for _, config := range configs {
		p.parse(config, p.put(to))
	}
}

func (p *Plugin) Parses(to map[string][]string, configs ...string) {
	for _, config := range configs {
		p.parse(config, p.puts(to))
	}
}

func (p *Plugin) Setup() (unset bool, err error) {
	unset = true
	err = nil

	return
}

func (p *Plugin) Fields() gox.Fields {
	return gox.Fields{
		field.Bool(`defaults`, p.Defaults),
		field.Bool(`verbose`, p.Verbose),
		field.Bool(`debug`, p.Debug),

		field.Bool(`retry`, p.Retry),
		field.Int(`counts`, p.Counts),
		field.Duration(`backoff`, p.Backoff),
	}
}

func (p *Plugin) Plugin() *Plugin {
	return p
}

func (p *Plugin) parse(original string, put func(configs []string)) {
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

func (p *Plugin) puts(cache map[string][]string) func(configs []string) {
	return func(configs []string) {
		if nil != configs && 2 <= len(configs) {
			value := strings.TrimSpace(configs[1])
			if `` == value {
				return
			}

			cache[strings.TrimSpace(configs[0])] = p.splits(value, `,`, `|`, `||`)
		}
	}
}

func (p *Plugin) put(cache map[string]string) func(configs []string) {
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

func (p *Plugin) splits(config string, seps ...string) (configs []string) {
	configs = []string{config}
	for _, sep := range seps {
		if strings.Contains(config, sep) {
			configs = strings.Split(config, sep)
			break
		}
	}

	return
}
