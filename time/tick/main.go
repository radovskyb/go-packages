package main

import (
	"fmt"
	"time"
)

func main() {
	// Tick is a convenience wrapper for NewTicker providing access to the ticking
	// channel only. While Tick is useful for clients that have no need to shut down
	// the Ticker, be aware that without a way to shut it down the underlying
	// Ticker cannot be recovered by the garbage collector; it "leaks".
	c := time.Tick(time.Second)

	// Print out the current time, every time a value is received from channel c
	// which in this case will be received once every second.
	for now := range c {
		fmt.Printf("Current time: %v\n", now)
	}
}
