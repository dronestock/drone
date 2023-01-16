package drone

var _ stepper = (*defaultStepper)(nil)

type defaultStepper struct{}

func (ds *defaultStepper) Runnable() bool {
	return true
}

func (ds *defaultStepper) Run() (err error) {
	return
}
