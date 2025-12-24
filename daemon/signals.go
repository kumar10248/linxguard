package daemon

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func HandleSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-sigChan
		log.Println("Received signal:", sig)
		log.Println("linxguard shutting down cleanly")
		os.Exit(0)
	}()
}
