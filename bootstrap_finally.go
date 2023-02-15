package drone

import (
	"os"

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
		b.Debug("清理插件出错")
	}

	// 退出程序，解决最外层panic报错的问题
	// 原理：如果到这个地方还没有发生错误，程序正常退出，外层panic得不到执行
	// 如果发生错误，则所有代码都会返回error直到panic检测到，然后程序整体panic
	if nil == *err {
		os.Exit(0)
	}
}
