package plugin

import (
	"context"
	"os"
	"path/filepath"
	"time"

	"github.com/dronestock/drone/internal/cleanup"
	"github.com/dronestock/drone/internal/command"
	"github.com/dronestock/drone/internal/config"
	"github.com/dronestock/drone/internal/core"
	"github.com/dronestock/drone/internal/internal"
	"github.com/dronestock/drone/internal/internal/constant"
	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/log"
)

var _ Config = (*Base)(nil)

// Base 插件基础
type Base struct {
	log.Logger

	// 是否启用默认配置
	Defaults *bool `default:"${DEFAULT=true}"`
	// 是否显示详细信息
	Verbose bool `default:"${VERBOSE=false}"`
	// 日志配置
	Level string `default:"${LEVEL=info}"`
	// 是否在出错时打印输出
	Pwe *bool `default:"${PWE=true}"`

	// 是否重试
	Retry *bool `default:"${RETRY=true}"`
	// 重试次数
	Counts int `default:"${COUNTS=5}"`
	// 重试间隔
	Backoff time.Duration `default:"${BACKOFF=5s}"`
	// 超时时间
	Timeout time.Duration `default:"${TIMEOUT=60m}"`

	// 代理
	Proxy *config.Proxy `default:"${PROXY}"`
	// 多个代理
	Proxies []*config.Proxy `default:"${PROXIES}"`
	// 卡片
	Card config.Card `default:"${CARD}"`
	// 命令列表
	Commands []string `default:"${COMMANDS}"`

	started  time.Time
	ctx      context.Context
	cancel   context.CancelFunc
	cleanups []*cleanup.Cleanup
	http     *http.Client
}

func (b *Base) Scheme() (scheme string) {
	return
}

func (b *Base) Carding() (card any, err error) {
	return
}

func (b *Base) Setup() (err error) {
	return
}

func (b *Base) Before() (err error) {
	return
}

func (b *Base) After() (err error) {
	return
}

func (b *Base) Elapsed() time.Duration {
	return b.Value("BUILD_STARTED").Time().Sub(b.Value("BUILD_CREATED").Time())
}

func (b *Base) Value(key string) *internal.Value {
	return internal.NewValue(os.Getenv(gox.StringBuilder(constant.DroneEnv, key).String()))
}

func (b *Base) Cleanup() *cleanup.Builder {
	return cleanup.NewBuilder(b.ctx, b.Pwe, b.Verbose, &b.cleanups)
}

func (b *Base) Since() time.Duration {
	return time.Since(b.started).Truncate(time.Second)
}

func (b *Base) Command(command string) (builder *command.Builder) {
	builder = gex.New(command)
	builder.Context(b.ctx)
	if nil == b.Pwe || *b.Pwe {
		builder.Pwe()
	}
	if b.Verbose {
		builder.Echo()
	}

	return
}

func (b *Base) Expressions() (expressions core.Expressions) {
	return
}

func (b *Base) Fields() gox.Fields[any] {
	return gox.Fields[any]{
		field.New("default", b.Defaults),
		field.New("verbose", b.Verbose),
		field.New("level", b.Level),

		field.New("retry", b.Retry),
		field.New("counts", b.Counts),
		field.New("backoff", b.Backoff),

		field.New("proxy", b.Proxy),
		field.New("commands", b.Commands),
	}
}

func (b *Base) base() *Base {
	return b
}

func (b *Base) Default() bool {
	return nil != b.Defaults && *b.Defaults
}

func (b *Base) Home(paths ...string) (final string) {
	if home, uhe := os.UserHomeDir(); nil == uhe {
		finals := make([]string, 0, len(paths)+1)
		finals = append(finals, home)
		finals = append(finals, paths...)
		final = filepath.Join(finals...)
	}

	return
}
