package drone

import (
	"os"
)

type cleanupBuilder struct {
	base    *Base
	name    string
	workers []worker
}

func newCleanupBuilder(base *Base) *cleanupBuilder {
	return &cleanupBuilder{
		base:    base,
		name:    "这个开发者很懒，没设置清理名称",
		workers: make([]worker, 0),
	}
}

func (cb *cleanupBuilder) Name(name string) *cleanupBuilder {
	cb.name = name

	return cb
}

func (cb *cleanupBuilder) File(names ...string) *cleanupBuilder {
	cb.workers = append(cb.workers, newFileWorker(names...))

	return cb
}

func (cb *cleanupBuilder) Write(filename string, data []byte, mode os.FileMode) *cleanupBuilder {
	cb.workers = append(cb.workers, newWriteWorker(filename, data, mode))

	return cb
}

func (cb *cleanupBuilder) Build() {
	cb.base.cleanups = append(cb.base.cleanups, newCleanup(cb.name, cb.workers...))
}
