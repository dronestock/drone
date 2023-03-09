package drone

import (
	"runtime"
	"strings"

	"github.com/goexl/gex"
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

func (b *Base) linux() (int, error) {
	return gex.New("/bin/sh").Args(args.New().Build().Arg("c", strings.Join(b.Commands, "; ")).Build()).Build().Exec()
}

func (b *Base) windows() (int, error) {
	ab := args.New()
	ab.Short("/")
	ab.Long("/")

	return gex.New("cmd").Args(ab.Build().Arg("C", strings.Join(b.Commands, "&& ")).Build()).Build().Exec()
}
