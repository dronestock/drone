package plugin

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/dronestock/drone/internal"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (b *Bootstrap) finally(err *error) {
	if finally := recover(); nil != finally {
		fmt.Println(finally)
		fmt.Println(b.stack())
		os.Exit(internal.ExitCodeFailed)
	}

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
		os.Exit(internal.ExitCodeOk)
	} else {
		os.Exit(internal.ExitCodeFailed)
	}
}

func (b *Bootstrap) stack() string {
	stack := 10
	skip := 4
	callers := make([]uintptr, stack+1)
	count := runtime.Callers(skip+2, callers)
	frames := runtime.CallersFrames(callers[:count])

	stacks := make([]string, 0, stack)
	for {
		frame, more := frames.Next()
		stacks = append(stacks, fmt.Sprintf("%s()\n\t%s:%d", frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}

	return strings.Join(stacks, "\n")
}
