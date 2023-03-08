package drone

import (
	"github.com/goexl/gox"
	"github.com/goexl/gox/args"
)

type commandBuilder struct {
	base    *Base
	command string
	params  *commandParams
}

func newCommandBuilder(base *Base, command string) *commandBuilder {
	return &commandBuilder{
		base:    base,
		command: command,
		params:  newCommandParams(gox.Ift(nil == base.Pwe, true, *base.Pwe)),
	}
}

func (cb *commandBuilder) Name(name string) *commandBuilder {
	cb.params.name = name

	return cb
}

func (cb *commandBuilder) Args(args *args.Args) *commandBuilder {
	for _, arg := range args.String() {
		cb.params.args = append(cb.params.args, arg)
	}

	return cb
}

func (cb *commandBuilder) Dir(dir string) *commandBuilder {
	cb.params.dir = dir

	return cb
}

func (cb *commandBuilder) Pwe(pwe bool) *commandBuilder {
	cb.params.pwe = pwe

	return cb
}

func (cb *commandBuilder) Async() *commandBuilder {
	cb.params.async = true

	return cb
}

func (cb *commandBuilder) Sync() *commandBuilder {
	cb.params.async = false

	return cb
}

func (cb *commandBuilder) Fields(fields ...gox.Field[any]) *commandBuilder {
	cb.params.fields = append(cb.params.fields, fields...)

	return cb
}

func (cb *commandBuilder) Field(field gox.Field[any]) *commandBuilder {
	return cb.Fields(field)
}

func (cb *commandBuilder) Collectors(collectors ...collectorBuilder) *commandBuilder {
	for _, builder := range collectors {
		cb.params.collectors = append(cb.params.collectors, builder.collector())
	}

	return cb
}

func (cb *commandBuilder) Collector(collector collectorBuilder) *commandBuilder {
	return cb.Collectors(collector)
}

func (cb *commandBuilder) Checkers(checkers ...checkerBuilder) *commandBuilder {
	for _, builder := range checkers {
		cb.params.checkers = append(cb.params.checkers, builder.checker())
	}

	return cb
}

func (cb *commandBuilder) Checker(checker checkerBuilder) *commandBuilder {
	return cb.Checkers(checker)
}

func (cb *commandBuilder) Build() *command {
	return newCommand(cb.base, cb.command, cb.params)
}
