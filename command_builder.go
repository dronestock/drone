package drone

import (
	"github.com/goexl/gox"
)

type commandBuilder struct {
	base    *Base
	command string
	options *commandOptions
}

func newCommand(base *Base, command string) *commandBuilder {
	return &commandBuilder{
		base:    base,
		command: command,
		options: &commandOptions{
			pwe: gox.Ifx(nil == base.Pwe, true, *base.Pwe),
		},
	}
}

func (cb *commandBuilder) Name(name string) *commandBuilder {
	cb.options.name = name

	return cb
}

func (cb *commandBuilder) Args(args ...any) *commandBuilder {
	cb.options.args = args

	return cb
}

func (cb *commandBuilder) Dir(dir string) *commandBuilder {
	cb.options.dir = dir

	return cb
}

func (cb *commandBuilder) Pwe(pwe bool) *commandBuilder {
	cb.options.pwe = pwe

	return cb
}

func (cb *commandBuilder) Async() *commandBuilder {
	cb.options.async = true

	return cb
}

func (cb *commandBuilder) Sync() *commandBuilder {
	cb.options.async = false

	return cb
}

func (cb *commandBuilder) Fields(fields ...gox.Field[any]) *commandBuilder {
	cb.options.fields = append(cb.options.fields, fields...)

	return cb
}

func (cb *commandBuilder) Field(field gox.Field[any]) *commandBuilder {
	return cb.Fields(field)
}

func (cb *commandBuilder) Collectors(collectors ...collectorBuilder) *commandBuilder {
	for _, builder := range collectors {
		cb.options.collectors = append(cb.options.collectors, builder.collector())
	}

	return cb
}

func (cb *commandBuilder) Collector(collector collectorBuilder) *commandBuilder {
	return cb.Collectors(collector)
}

func (cb *commandBuilder) Checkers(checkers ...checkerBuilder) *commandBuilder {
	for _, builder := range checkers {
		cb.options.checkers = append(cb.options.checkers, builder.checker())
	}

	return cb
}

func (cb *commandBuilder) Checker(checker checkerBuilder) *commandBuilder {
	return cb.Checkers(checker)
}
