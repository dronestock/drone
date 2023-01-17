package drone

import (
	"context"
	"fmt"
	"os"
	"sync"
)

func (b *bootstrap) exec() (err error) {
	// 开始卡片信息写入
	go b.startCard()

	// 执行插件
	wg := new(sync.WaitGroup)
	ctx := context.Background()
	for _, step := range b.plugin.Steps() {
		if err = b.execStep(ctx, step, wg); nil != err && step.options._break {
			return
		}
	}
	wg.Wait()

	// 写入最终数据到卡片中
	_ = b.writeCard()

	// 记录日志
	if nil != err {
		b.Error(fmt.Sprintf("%s插件执行出错，请检查", b.name))
	} else {
		b.Info(fmt.Sprintf("%s插件执行成功，恭喜", b.name))

		// 退出程序，解决最外层panic报错的问题
		// 原理：如果到这个地方还没有发生错误，程序正常退出，外层panic得不到执行
		// 如果发生错误，则所有代码都会返回error直到panic检测到，然后程序整体panic
		os.Exit(0)
	}

	return
}
