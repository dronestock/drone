package drone

var (
	_               = Dir
	_ commandOption = (*optionDir)(nil)
	_ execOption    = (*optionDir)(nil)
)

type optionDir struct {
	dir string
}

// Dir 配置命令执行目录
func Dir(dir string) *optionDir {
	return &optionDir{
		dir: dir,
	}
}

func (d *optionDir) applyCommand(options *commandOptions) {
	options.dir = d.dir
}

func (d *optionDir) applyExec(options *execOptions) {
	options.dir = d.dir
}
