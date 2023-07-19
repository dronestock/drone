package step

import (
	"github.com/goexl/gox"
)

type Options struct {
	Name  string
	Async bool
	Retry bool
	Break bool
}

func (o *Options) Retryable(base *bool) (retry bool) {
	retry = o.Retry
	// 优先以本步骤的配置为准，如果本步骤配置为不重试，再以全局配置为准
	if retry {
		retry = gox.Ift(nil == base, true, *base)
	}

	return
}
