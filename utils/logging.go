package logging

import (
	"log"
)

func SetUpLog() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("LOG: ")
}
