package drone

var (
	_            = Pwe
	_            = Sync
	_            = DisablePwe
	_ execOption = (*optionPwe)(nil)
)

type optionPwe struct {
	pwe bool
}

// Pwe 启用PWE
func Pwe() *optionPwe {
	return &optionPwe{
		pwe: true,
	}
}

// DisablePwe 禁用PWE
func DisablePwe() *optionPwe {
	return &optionPwe{
		pwe: false,
	}
}

func (p *optionPwe) applyExec(options *execOptions) {
	options.pwe = p.pwe
}
