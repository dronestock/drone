package cleanup

import (
	"github.com/dronestock/drone/internal/cleanup/internal/work"
	"github.com/goexl/gex"
	"github.com/goexl/gox/args"
)

type Command struct {
	gex     *gex.Builder
	builder *Builder
}

func (c *Command) Args(args *args.Args) (command *Command) {
	c.gex.Args(args)
	command = c

	return
}

func (c *Command) Build() (builder *Builder) {
	c.builder.workers = append(c.builder.workers, work.NewCommand(c.gex))
	builder = c.builder

	return
}
