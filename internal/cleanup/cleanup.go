package cleanup

import (
	"github.com/dronestock/drone/internal/cleanup/internal/work"
	"github.com/goexl/simaqian"
)

type Cleanup struct {
	Name    string
	Workers []work.Worker
}

func New(name string, workers ...work.Worker) *Cleanup {
	return &Cleanup{
		Name:    name,
		Workers: workers,
	}
}

func (c *Cleanup) Clean(logger simaqian.Logger) (err error) {
	for _, worker := range c.Workers {
		err = worker.Work(logger)
	}

	return
}
