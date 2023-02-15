package drone

import (
	"os"

	"github.com/goexl/gox/field"
)

var _ worker = (*fileWorker)(nil)

type fileWorker struct {
	names []string
}

func newFileWorker(names ...string) *fileWorker {
	return &fileWorker{
		names: names,
	}
}

func (fc *fileWorker) work(base *Base) (err error) {
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
