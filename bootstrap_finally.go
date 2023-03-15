package drone

import (
	"os"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (b *bootstrap) finally(err *error) {
	if code, ee := b.exec(); nil != ee {
		b.Warn("执行命令出错", field.New("code", code), field.New("commands", b.Commands), field.Error(ee))
	} else {
		b.Debug("执行命令成功", field.New("code", code), field.New("commands", b.Commands))
	}
	if ce := b.cleanup(); nil != ce {
		b.Warn("清理插件出错", field.Error(ce))
	} else {
		b.Debug("清理插件成功")
	}

	fields := gox.Fields[any]{
		field.New("duration", time.Since(b.started).Truncate(time.Second)),
	}
	b.Info("插件执行完成", fields...)
	if nil == *err {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
