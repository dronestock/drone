package drone

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (b *bootstrap) execStep(step *Step, wg *sync.WaitGroup) (err error) {
	if step.options.async {
		err = b.execStepAsync(step, wg)
	} else {
		err = b.execStepSync(step)
	}

	return
}

func (b *bootstrap) execStepSync(step *Step) error {
	return b.execStepper(step.stepper, step.options)
}

func (b *bootstrap) execStepAsync(step *Step, wg *sync.WaitGroup) (err error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = b.execStepper(step.stepper, step.options); nil != err {
			panic(err)
		}
	}()

	return
}

func (b *bootstrap) execStepper(stepper stepper, options *stepOptions) (err error) {
	if !stepper.Runnable() {
		return
	}

	retry := options.retryable(b.Base)
	fields := gox.Fields[any]{
		field.New("name", options.name),
		field.New("async", options.async),
		field.New("retry", retry),
		field.New("break", options._break),
		field.New("counts", b.Counts),
	}

	b.Info("步骤执行开始", fields...)
	rand.Seed(time.Now().UnixNano())
	for count := 0; count < b.Counts; count++ {
		if err = stepper.Run(); (nil == err) && (0 == count && !retry) {
			break
		}

		backoff := b.backoff()
		b.Info(fmt.Sprintf("步骤第%d次执行遇到错误", count+1), fields.Connect(field.Error(err))...)
		b.Info(fmt.Sprintf("休眠%s，继续执行步骤", backoff), fields...)
		time.Sleep(backoff)
		b.Info(fmt.Sprintf("步骤重试第%d次执行", count+2), fields...)

		if count != b.Counts-1 {
			err = nil
		}
	}

	switch {
	case nil != err && retry:
		b.Error("步骤执行尝试所有重试后出错", fields.Connect(field.Error(err))...)
	case nil != err && !retry:
		b.Error("步骤执行出错", fields.Connect(field.Error(err))...)
	case nil == err:
		b.Info("步骤执行成功", fields...)
	}

	return
}
