package drone

var (
	_ checkerBuilder = (*contains)(nil)
	_                = Contains
)

type contains struct {
	contains string
}

// Contains 检查是否包含字符串
func Contains(cas string) *contains {
	return &contains{
		contains: cas,
	}
}

func (c *contains) checker() *checker {
	return newChecker(checkerModeContains, c.contains)
}
