package work

import (
	"github.com/goexl/gex"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
)

var _ Worker = (*Command)(nil)

type Command struct {
	builder *gex.Builder
}

func NewCommand(builder *gex.Builder) *Command {
	return &Command{
		builder: builder,
	}
}

func (c *Command) Work(logger log.Logger) (err error) {
	if code, ee := c.builder.Build().Exec(); nil != ee {
		err = ee
		logger.Warn("命令执行出错", field.New("code", code), field.Error(err))
	} else {
		logger.Warn("命令执行出错", field.New("code", code))
	}

	return
}
