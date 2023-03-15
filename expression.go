package drone

type expression interface {
	// Name 名称
	Name() string

	// Exec 实际执行方法
	Exec(args ...any) (result any, err error)
}
