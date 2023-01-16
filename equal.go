package drone

var (
	_ checkerBuilder = (*_equal)(nil)
	_                = Equal
)

type _equal struct {
	equal string
}

// Equal 检查是否字符串相等
func Equal(equal string) *_equal {
	return &_equal{
		equal: equal,
	}
}

func (e *_equal) checker() *checker {
	return newChecker(checkerModeEqual, e.equal)
}
