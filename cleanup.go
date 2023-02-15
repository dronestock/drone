package drone

type cleanup interface {
	clean(base *Base) (err error)
}
