package plugin

import (
	"runtime"
	"strings"

	"github.com/dronestock/drone/internal/command"
	"github.com/goexl/args"
	"github.com/goexl/gox/field"
)

func (b *Bootstrap) exec() (handler command.Handler, err error) {
	if 0 == len(b.Commands) {
		return
	}

	b.Info("执行命令开始", field.New("commands", b.Commands))
	switch runtime.GOOS {
	case "windows":
		handler, err = b.windows()
	case "linux":
		handler, err = b.linux()
	}

	return
}

func (b *Bootstrap) linux() (handler command.Handler, err error) {
	name := "/bin/sh"
	arguments := args.New().Build()
	arguments.Option("c", strings.Join(b.Commands, ";"))
	handler, err = b.Command(name).Arguments(arguments.Build()).Build().Exec()

	return
}

func (b *Bootstrap) windows() (handler command.Handler, err error) {
	name := "cmd"
	arguments := args.New().Short("/").Long("/").Build()
	arguments.Option("C", strings.Join(b.Commands, "&&"))
	handler, err = b.Command(name).Arguments(arguments.Build()).Build().Exec()

	return
}
