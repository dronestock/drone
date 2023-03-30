package drone

import (
	"os"
	"os/signal"
	"syscall"
)

func (b *bootstrap) gracefully() {
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go b.exit(signals)
}

func (b *bootstrap) exit(signals chan os.Signal) {
	select {
	case <-signals:
		os.Exit(2)
	}
}
