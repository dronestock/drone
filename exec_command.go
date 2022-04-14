package drone

func (b *Base) execWithCommandOptions(command string, opts ...commandOption) (err error) {
	_options := defaultCommandOptions()
	for _, opt := range opts {
		opt.applyCommand(_options)
	}

	return b.exec(command, newExecOptions(_options))
}
