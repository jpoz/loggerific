package main

import (
	"log"
	"os"
	"time"

	"github.com/gofrs/uuid"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func writeLog(t time.Time) {
	u, _ := uuid.NewV4()
	log.Printf("[%s] event", u)
}

func main() {
	u, _ := uuid.NewV4()
	log.SetOutput(os.Stdout)
	log.Printf("[%s] hello", u)
	doEvery(100*time.Millisecond, writeLog)
}
