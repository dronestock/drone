package drone

var (
	_ checkerBuilder = (*equal)(nil)
	_                = Equal
)

type equal struct {
	equal string
}

// Equal 检查是否字符串相等
func Equal(eq string) *equal {
	return &equal{
		equal: eq,
	}
}

func (e *equal) checker() *checker {
	return newChecker(checkerModeEqual, e.equal)
}
