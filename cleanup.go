package drone

type cleanup struct {
	name    string
	workers []worker
}

func newCleanup(name string, workers ...worker) *cleanup {
	return &cleanup{
		name:    name,
		workers: workers,
	}
}

func (c *cleanup) clean(base *Base) (err error) {
	for _, worker := range c.workers {
		err = worker.work(base)
	}

	return
}
