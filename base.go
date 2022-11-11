package drone

import (
	"math/rand"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

var _ Config = (*Base)(nil)

// Base 插件基础
type Base struct {
	simaqian.Logger

	// 是否启用默认配置
	Defaults bool `default:"${PLUGIN_DEFAULTS=${DEFAULTS=true}}"`
	// 是否显示详细信息
	Verbose bool `default:"${PLUGIN_VERBOSE=${VERBOSE=false}}"`
	// 是否显示调试信息
	Level string `default:"${PLUGIN_LEVEL=${LEVEL=debug}}"`
	// 是否在出错时打印输出
	Pwe bool `default:"${PLUGIN_PWE=${PWE=true}}"`

	// 是否重试
	Retry bool `default:"${PLUGIN_RETRY=${RETRY=true}}"`
	// 重试次数
	Counts int `default:"${PLUGIN_COUNTS=${COUNTS=5}}"`
	// 重试间隔
	Backoff time.Duration `default:"${PLUGIN_BACKOFF=${BACKOFF=5s}}"`

	// 卡片路径
	CardPath string `default:"${DRONE_CARD_PATH=/dev/stdout}"`
}

func (b *Base) Scheme() (scheme string) {
	return
}

func (b *Base) Card() (card any, err error) {
	return
}

func (b *Base) Interval() time.Duration {
	return time.Second
}

func (b *Base) Setup() (unset bool, err error) {
	unset = true
	err = nil

	return
}

func (b *Base) Fields() gox.Fields {
	return gox.Fields{
		field.Bool(`defaults`, b.Defaults),
		field.Bool(`verbose`, b.Verbose),
		field.String(`level`, b.Level),

		field.Bool(`retry`, b.Retry),
		field.Int(`counts`, b.Counts),
		field.Duration(`backoff`, b.Backoff),
	}
}

func (b *Base) BaseConfig() *Base {
	return b
}

func (b *Base) backoff() time.Duration {
	from := time.Duration(int64(float64(b.Backoff) * 0.3))
	offset := time.Duration(rand.Int63n(int64(b.Backoff / 2))).Truncate(time.Second)

	return from + offset
}
