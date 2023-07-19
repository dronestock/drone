package plugin

import (
	"os"
	"os/signal"
	"syscall"
)

func (b *Bootstrap) gracefully() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go b.exit(signals)
}

func (b *Bootstrap) exit(signals chan os.Signal) {
	<-signals
	os.Exit(2)
}
