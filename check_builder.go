package drone

type checkerBuilder interface {
	checker() *checker
}
