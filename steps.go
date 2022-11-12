package drone

// Steps 步骤集
type Steps []*Step

func (s *Steps) Add(steps ...*Step) {
	*s = append(*s, steps...)
}
