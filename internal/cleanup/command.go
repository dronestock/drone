package cleanup

import (
	"context"

	"github.com/dronestock/drone/internal/cleanup/internal/work"
	"github.com/goexl/args"
	"github.com/goexl/gex"
)

type Command struct {
	command *gex.Command
	builder *Builder
}

func NewCommand(ctx context.Context, builder *Builder, pwe *bool, verbose bool, name string) (command *Command) {
	command = new(Command)
	command.command = gex.New(name)
	command.builder = builder
	command.command.Context(ctx)
	if nil == pwe || *pwe {
		command.command.Pwe()
	}
	if verbose {
		command.command.Echo()
	}

	return
}

func (c *Command) Arguments(arguments *args.Arguments) (command *Command) {
	c.command.Arguments(arguments)
	command = c

	return
}

func (c *Command) Build() (builder *Builder) {
	c.builder.workers = append(c.builder.workers, work.NewCommand(c.command))
	builder = c.builder

	return
}
