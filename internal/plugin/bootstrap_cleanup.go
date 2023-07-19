package plugin

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (b *Bootstrap) cleanup() (err error) {
	for _, cleanup := range b.cleanups {
		fields := gox.Fields[any]{
			field.New("name", cleanup.Name),
		}
		if ce := cleanup.Clean(b.Logger); nil != ce {
			err = ce
			b.Warn("执行清理出错", fields.Add(field.Error(ce))...)
		} else {
			b.Debug("执行清理成功", fields...)
		}
	}

	return
}
