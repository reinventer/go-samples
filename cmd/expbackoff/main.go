package main

import (
	"github.com/reinventer/go-samples/expbackoff"

	"log"
	"time"
)

func main() {
	b := expbackoff.New(2*time.Millisecond, 10*time.Second)
	defer b.Stop()

	c := b.Fetch()
	for i := 0; i < 30; i++ {
		t := time.Now()
		<-c

		log.Printf("tick %d (%s)", i, time.Since(t))

		if i == 15 {
			b.Reset()
			log.Print("reset")
		} else if i < 15 || i > 15 {
			b.Jam()
		}
	}
}
