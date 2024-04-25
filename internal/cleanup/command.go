package cleanup

import (
	"context"

	"github.com/dronestock/drone/internal/cleanup/internal/work"
	"github.com/goexl/args"
	"github.com/goexl/gex"
)

type Command struct {
	gex     *gex.Builder
	builder *Builder
}

func NewCommand(ctx context.Context, builder *Builder, pwe *bool, verbose bool, name string) (command *Command) {
	command = new(Command)
	command.gex = gex.New(name)
	command.builder = builder
	command.gex.Context(ctx)
	if nil == pwe || *pwe {
		command.gex.Pwe()
	}
	if verbose {
		command.gex.Echo()
	}

	return
}

func (c *Command) Arguments(arguments *args.Arguments) (command *Command) {
	c.gex.Args(arguments)
	command = c

	return
}

func (c *Command) Build() (builder *Builder) {
	c.builder.workers = append(c.builder.workers, work.NewCommand(c.gex))
	builder = c.builder

	return
}
