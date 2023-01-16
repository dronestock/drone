package drone

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (bb *bootstrapBuilder) execStep(step *Step, wg *sync.WaitGroup, base *Base) (err error) {
	if step.options.async {
		err = bb.execStepAsync(step, wg, base)
	} else {
		err = bb.execStepSync(step, base)
	}

	return
}

func (bb *bootstrapBuilder) execStepSync(step *Step, base *Base) error {
	return bb.execStepper(step.stepper, step.options, base)
}

func (bb *bootstrapBuilder) execStepAsync(step *Step, wg *sync.WaitGroup, base *Base) (err error) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err = bb.execStepper(step.stepper, step.options, base); nil != err {
			panic(err)
		}
	}()

	return
}

func (bb *bootstrapBuilder) execStepper(stepper stepper, options *stepOptions, base *Base) (err error) {
	if !stepper.Runnable() {
		return
	}

	retry := options.retryable(base)
	fields := gox.Fields[any]{
		field.New("name", options.name),
		field.New("async", options.async),
		field.New("retry", retry),
		field.New("break", options._break),
		field.New("counts", base.Counts),
	}

	base.Info("步骤执行开始", fields...)
	rand.Seed(time.Now().UnixNano())
	for count := 0; count < base.Counts; count++ {
		if err = stepper.Run(); (nil == err) && (0 == count && !retry) {
			break
		}

		backoff := base.backoff()
		base.Info(fmt.Sprintf("步骤第%d次执行遇到错误", count+1), fields.Connect(field.Error(err))...)
		base.Info(fmt.Sprintf("休眠%s，继续执行步骤", backoff), fields...)
		time.Sleep(backoff)
		base.Info(fmt.Sprintf("步骤重试第%d次执行", count+2), fields...)

		if count != base.Counts-1 {
			err = nil
		}
	}

	switch {
	case nil != err && retry:
		base.Error("步骤执行尝试所有重试后出错", fields.Connect(field.Error(err))...)
	case nil != err && !retry:
		base.Error("步骤执行出错", fields.Connect(field.Error(err))...)
	case nil == err:
		base.Info("步骤执行成功", fields...)
	}

	return
}
