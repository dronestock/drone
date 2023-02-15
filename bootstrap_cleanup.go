package drone

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (b *Base) cleanup() (err error) {
	for _, cleanup := range b.cleanups {
		fields := gox.Fields[any]{
			field.New("name", cleanup.name),
		}
		if ce := cleanup.clean(b); nil != ce {
			err = ce
			b.Warn("执行清理出错", fields.Connect(field.Error(ce))...)
		} else {
			b.Debug("执行清理成功", fields...)
		}
	}

	return
}
