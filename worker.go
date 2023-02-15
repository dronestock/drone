package drone

type worker interface {
	work(base *Base) (err error)
}
