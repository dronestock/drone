package plugin

import (
	"context"
	"fmt"
	"sync"
)

func (b *Bootstrap) run() (err error) {
	// 设置整体超时时间
	b.ctx, b.cancel = context.WithTimeout(context.Background(), b.Timeout)
	defer b.cancel()
	// 开始卡片信息写入
	go b.startCard()

	// 执行插件
	wg := new(sync.WaitGroup)
	for _, step := range b.plugin.Steps() {
		if err = b.execStep(b.ctx, step, wg); nil != err && step.Options.Break {
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
	}

	return
}
