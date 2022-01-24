package drone

var (
	_        = Aliases
	_ option = (*optionAliases)(nil)
)

type optionAliases struct {
	aliases []*alias
}

// Aliases 别名
func Aliases(aliases ...*alias) *optionAliases {
	return &optionAliases{
		aliases: aliases,
	}
}

func (a *optionAliases) apply(options *options) {
	options.aliases = a.aliases
}
