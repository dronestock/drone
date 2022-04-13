package drone

type checker struct {
	mode checkerMode
	args interface{}
}

func newChecker(mode checkerMode, args interface{}) *checker {
	return &checker{
		mode: mode,
		args: args,
	}
}
