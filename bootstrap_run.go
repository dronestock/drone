package drone

import (
	"context"
	"fmt"
	"sync"
)

func (b *bootstrap) run() (err error) {
	// 开始卡片信息写入
	go b.startCard()

	// 执行插件
	wg := new(sync.WaitGroup)
	ctx := context.Background()
	for _, step := range b.plugin.Steps() {
		if err = b.execStep(ctx, step, wg); nil != err && step.options.br {
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
