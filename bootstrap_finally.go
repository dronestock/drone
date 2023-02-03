package drone

import (
	"os"
)

func (b *bootstrap) finally(err *error) {
	_ = b.commands()

	// 退出程序，解决最外层panic报错的问题
	// 原理：如果到这个地方还没有发生错误，程序正常退出，外层panic得不到执行
	// 如果发生错误，则所有代码都会返回error直到panic检测到，然后程序整体panic
	if nil == *err {
		os.Exit(0)
	}
}
