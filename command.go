package drone

import (
	"runtime"
	"strings"

	"github.com/goexl/gex"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

func (b *Base) Command(command string) *commandBuilder {
	return newCommand(b, command)
}

func (b *Base) commands() (err error) {
	if 0 == len(b.Commands) {
		return
	}

	b.Info("执行命令开始", field.New("commands", b.Commands))
	code := 0
	switch runtime.GOOS {
	case "windows":
		code, err = b.windows()
	case "linux":
		code, err = b.linux()
	}

	fields := gox.Fields[any]{
		field.New("code", code),
		field.New("commands", b.Commands),
	}
	if nil != err {
		b.Warn("执行命令出错", fields.Connect(field.Error(err))...)
	} else {
		b.Info("执行命令成功", fields...)
	}

	return
}

func (b *Base) linux() (int, error) {
	return gex.Exec("/bin/sh", gex.Args("-c", strings.Join(b.Commands, "; ")), gex.Terminal())
}

func (b *Base) windows() (int, error) {
	return gex.Exec("cmd", gex.Args("/C", strings.Join(b.Commands, "&& ")), gex.Terminal())
}
