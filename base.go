package drone

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
)

var _ Config = (*Base)(nil)

// Base 插件基础
type Base struct {
	simaqian.Logger

	// 是否启用默认配置
	Defaults bool `default:"${DEFAULTS=true}"`
	// 是否显示详细信息
	Verbose bool `default:"${VERBOSE=false}"`
	// 日志配置
	Log log `default:"${LOG}"`
	// 是否在出错时打印输出
	Pwe bool `default:"${PWE=true}"`

	// 是否重试
	Retry bool `default:"${RETRY=true}"`
	// 重试次数
	Counts int `default:"${COUNTS=5}"`
	// 重试间隔
	Backoff time.Duration `default:"${BACKOFF=5s}"`

	// 代理
	Proxy *proxy `default:"${PROXY}"`
	// 卡片
	Card card `default:"${CARD}"`

	http *resty.Client
}

func (b *Base) Scheme() (scheme string) {
	return
}

func (b *Base) Carding() (card any, err error) {
	return
}

func (b *Base) Setup() (unset bool, err error) {
	unset = true
	err = nil

	return
}

func (b *Base) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("defaults", b.Defaults),
		field.New("verbose", b.Verbose),
		field.New("level", b.Level),

		field.New("retry", b.Retry),
		field.New("counts", b.Counts),
		field.New("backoff", b.Backoff),

		field.New("proxy", b.Proxy),
	}
}

func (b *Base) BaseConfig() *Base {
	return b
}

func (b *Base) backoff() (backoff time.Duration) {
	from := time.Duration(int64(float64(b.Backoff) * 0.3))
	if duration, re := rand.Int(rand.Reader, big.NewInt(int64(b.Backoff-from))); nil != re {
		backoff = b.Backoff
	} else {
		backoff = from + time.Duration(duration.Int64()).Truncate(time.Second)
	}

	return backoff
}
