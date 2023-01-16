package drone

type collectorBuilder interface {
	collector() *collector
}
