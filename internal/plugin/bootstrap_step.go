package plugin

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/dronestock/drone/internal/step"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/gox/rand"
)

func (b *Bootstrap) execStep(ctx context.Context, step *step.Step, wg *sync.WaitGroup) (err error) {
	if step.Options.Async {
		err = b.execStepAsync(ctx, step, wg)
	} else {
		err = b.execStepSync(ctx, step)
	}

	return
}

func (b *Bootstrap) execStepSync(ctx context.Context, step *step.Step) error {
	return b.execStepper(ctx, step.Stepper, step.Options)
}

func (b *Bootstrap) execStepAsync(ctx context.Context, step *step.Step, wg *sync.WaitGroup) (err error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = b.execStepper(ctx, step.Stepper, step.Options); nil != err {
			panic(err)
		}
	}()

	return
}

func (b *Bootstrap) execStepper(ctx context.Context, stepper step.Stepper, options *step.Options) (err error) {
	if !stepper.Runnable() {
		return
	}

	retry := options.Retryable(b.Base.Retry)
	fields := gox.Fields[any]{
		field.New("name", options.Name),
		field.New("async", options.Async),
		field.New("retry", retry),
		field.New("break", options.Break),
		field.New("counts", b.Counts),
	}

	b.Info("步骤执行开始", fields...)
	for count := 0; count < b.Counts; count++ {
		if err = stepper.Run(ctx); (nil == err) || (0 == count && !retry) {
			break
		}

		backoff := rand.New().Duration().Between(time.Second, b.Backoff).Build().Generate().Truncate(time.Second)
		b.Info(fmt.Sprintf("步骤第%d次执行遇到错误", count+1), fields.Add(field.Error(err))...)
		b.Info(fmt.Sprintf("休眠%s，继续执行步骤", backoff), fields...)
		time.Sleep(backoff)
		b.Info(fmt.Sprintf("步骤重试第%d次执行", count+2), fields...)

		if count != b.Counts-1 {
			err = nil
		}
	}

	switch {
	case nil != err && retry:
		b.Error("步骤执行尝试所有重试后出错", fields.Add(field.Error(err))...)
	case nil != err && !retry:
		b.Error("步骤执行出错", fields.Add(field.Error(err))...)
	case nil == err:
		b.Info("步骤执行成功", fields...)
	}

	return
}
