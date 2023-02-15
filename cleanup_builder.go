package drone

type cleanupBuilder struct {
	base     *Base
	name     string
	cleanups []cleanup
}

func newCleanupBuilder(base *Base) *cleanupBuilder {
	return &cleanupBuilder{
		base:     base,
		name:     "开发者很懒，没设置清理名称",
		cleanups: make([]cleanup, 0),
	}
}

func (cb *cleanupBuilder) Name(name string) *cleanupBuilder {
	cb.name = name

	return cb
}

func (cb *cleanupBuilder) File(names ...string) *cleanupBuilder {
	cb.cleanups = append(cb.cleanups, newFileCleanup(names...))

	return cb
}

func (cb *cleanupBuilder) Build() {
	cb.base.cleanups = append(cb.base.cleanups, cb.cleanups...)
}
