package drone

type checker struct {
	mode checkerMode
	args any
}

func newChecker(mode checkerMode, args any) *checker {
	return &checker{
		mode: mode,
		args: args,
	}
}
