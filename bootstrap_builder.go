package drone

var _ = New

type bootstrapBuilder struct {
	constructor constructor
	name        string
	aliases     []*alias
}

func New(constructor constructor) *bootstrapBuilder {
	return &bootstrapBuilder{
		constructor: constructor,
	}
}

func (bb *bootstrapBuilder) Name(name string) *bootstrapBuilder {
	bb.name = name

	return bb
}

func (bb *bootstrapBuilder) Boot() (err error) {
	bb.name = name

	return bb
}
