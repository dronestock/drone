package work

import (
	"os"

	"github.com/goexl/simaqian"
)

var _ Worker = (*Write)(nil)

type Write struct {
	filename string
	data     []byte
	mode     os.FileMode
}

func NewWrite(filename string, data []byte, mode os.FileMode) *Write {
	return &Write{
		filename: filename,
		data:     data,
		mode:     mode,
	}
}

func (w *Write) Work(_ simaqian.Logger) error {
	return os.WriteFile(w.filename, w.data, w.mode)
}
