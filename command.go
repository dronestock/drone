package drone

func (b *Base) Command(command string) *commandBuilder {
	return newCommand(b, command)
}
