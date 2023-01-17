package drone

import (
	"github.com/goexl/gox"
)

type stepOptions struct {
	name   string
	async  bool
	retry  bool
	_break bool
}

func (o *stepOptions) retryable(base *Base) (retry bool) {
	retry = o.retry
	// 优先以本步骤的配置为准，如果本步骤配置为不重试，再以全局配置为准
	if retry {
		retry = gox.If(nil == base.Retry, true, *base.Retry)
	}

	return
}
