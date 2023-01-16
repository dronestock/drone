package drone

var (
	_ checkerBuilder = (*_contains)(nil)
	_                = Contains
)

type _contains struct {
	contains string
}

// Contains 检查是否包含字符串
func Contains(contains string) *_contains {
	return &_contains{
		contains: contains,
	}
}

func (c *_contains) checker() *checker {
	return newChecker(checkerModeContains, c.contains)
}
