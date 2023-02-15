package drone

import (
	"os"

	"github.com/goexl/gox/field"
)

var _ cleanup = (*fileCleanup)(nil)

type fileCleanup struct {
	names []string
}

func newFileCleanup(names ...string) *fileCleanup {
	return &fileCleanup{
		names: names,
	}
}

func (fc *fileCleanup) clean(base *Base) (err error) {
	for _, name := range fc.names {
		if re := os.Remove(name); nil != re {
			err = re
			base.Warn("执行文件清理出错", field.New("filename", name), field.Error(re))
		} else {
			base.Debug("执行文件清理成功", field.New("filename", name), field.Error(re))
		}
	}

	return
}
