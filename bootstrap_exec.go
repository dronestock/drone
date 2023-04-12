package drone

import (
	"runtime"
	"strings"

	"github.com/goexl/gox/args"
	"github.com/goexl/gox/field"
)

func (b *Base) exec() (code int, err error) {
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

func (b *Base) linux() (code int, err error) {
	name := "/bin/sh"
	ab := args.New().Build()
	ab.Flag("c").Add(strings.Join(b.Commands, ";"))
	code, err = b.Command(name).Args(ab.Build()).Build().Exec()

	return
}

func (b *Base) windows() (code int, err error) {
	name := "cmd"
	ab := args.New().Short("/").Long("/").Build()
	ab.Flag("C").Add(strings.Join(b.Commands, "&&"))
	code, err = b.Command(name).Args(ab.Build()).Build().Exec()

	return
}
