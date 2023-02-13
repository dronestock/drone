package drone

import (
	"runtime"
	"strings"

	"github.com/goexl/gex"
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
	return gex.Exec("/bin/sh", gex.Args("-c", strings.Join(b.Commands, "; ")), gex.Terminal())
}

func (b *Base) windows() (int, error) {
	return gex.Exec("cmd", gex.Args("/C", strings.Join(b.Commands, "&& ")), gex.Terminal())
}
