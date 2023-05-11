package drone

import (
	"os"
)

var _ worker = (*writeWorker)(nil)

type writeWorker struct {
	filename string
	data []byte
	mode os.FileMode
}

func newWriteWorker(filename string, data []byte, mode os.FileMode) *writeWorker {
	return &writeWorker{
		filename: filename,
		data:data,
		mode:mode,
	}
}

func (wc *writeWorker) work(base *Base) error {
	return os.WriteFile(wc.filename, wc.data, wc.mode)
}
