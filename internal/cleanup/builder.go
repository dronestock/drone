package cleanup

import (
	"context"
	"os"

	"github.com/dronestock/drone/internal/cleanup/internal/work"
)

type Builder struct {
	ctx     context.Context
	pwe     *bool
	verbose bool

	cleanups *[]*Cleanup
	name     string
	workers  []work.Worker
}

func NewBuilder(ctx context.Context, pwe *bool, verbose bool, cleanups *[]*Cleanup) *Builder {
	return &Builder{
		ctx:     ctx,
		pwe:     pwe,
		verbose: verbose,

		cleanups: cleanups,
		name:     "这个开发者很懒，没设置清理名称",
		workers:  make([]work.Worker, 0),
	}
}

func (b *Builder) Name(name string) *Builder {
	b.name = name

	return b
}

func (b *Builder) File(names ...string) *Builder {
	b.workers = append(b.workers, work.NewFile(names...))

	return b
}

func (b *Builder) Write(filename string, data []byte, mode os.FileMode) *Builder {
	b.workers = append(b.workers, work.NewWrite(filename, data, mode))

	return b
}

func (b *Builder) Command(name string) *Command {
	return NewCommand(b.ctx, b, b.pwe, b.verbose, name)
}

func (b *Builder) Build() {
	*b.cleanups = append(*b.cleanups, New(b.name, b.workers...))
}
