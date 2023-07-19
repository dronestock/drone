package plugin

import (
	"runtime"
	"strings"

	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

func (b *Bootstrap) exec() (code int, err error) {
	if 0 == len(b.Commands) {
		return
	}

	b.Info("执行命令开始", field.New("commands", b.Commands))
	switch runtime.GOOS {
	case "windows":
		code, err = b.windows()
	case "linux":
		code, err = b.linux()
	}

	return
}

func (b *Bootstrap) linux() (code int, err error) {
	name := "/bin/sh"
	ab := args.New().Build()
	ab.Option("c", strings.Join(b.Commands, ";"))
	code, err = b.Command(name).Args(ab.Build()).Build().Exec()

	return
}

func (b *Bootstrap) windows() (code int, err error) {
	name := "cmd"
	ab := args.New().Short("/").Long("/").Build()
	ab.Option("C", strings.Join(b.Commands, "&&"))
	code, err = b.Command(name).Args(ab.Build()).Build().Exec()

	return
}
