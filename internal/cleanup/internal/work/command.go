package work

import (
	"github.com/goexl/gex"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
)

var _ Worker = (*Command)(nil)

type Command struct {
	command *gex.Command
}

func NewCommand(command *gex.Command) *Command {
	return &Command{
		command: command,
	}
}

func (c *Command) Work(logger log.Logger) (err error) {
	if code, ee := c.command.Build().Exec(); nil != ee {
		err = ee
		logger.Warn("命令执行出错", field.New("code", code), field.Error(err))
	} else {
		logger.Warn("命令执行出错", field.New("code", code))
	}

	return
}
