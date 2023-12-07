package work

import (
	"os"

	"github.com/goexl/log"

	"github.com/goexl/gox/field"
)

var _ Worker = (*File)(nil)

type File struct {
	names []string
}

func NewFile(names ...string) *File {
	return &File{
		names: names,
	}
}

func (f *File) Work(logger log.Logger) (err error) {
	for _, name := range f.names {
		if re := os.RemoveAll(name); nil != re {
			err = re
			logger.Warn("执行文件清理出错", field.New("filename", name), field.Error(re))
		} else {
			logger.Debug("执行文件清理成功", field.New("filename", name), field.Error(re))
		}
	}

	return
}
