package logger

import (
	"log"
)

func Init() {
	// systemd captures stdout/stderr automatically
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
