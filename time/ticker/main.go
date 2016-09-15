// A Ticker holds a channel that delivers `ticks' of a clock at intervals.
package main

import (
	"fmt"
	"time"
)

type Ticker struct {
	ticker *time.Ticker
	close  chan struct{}
}

func NewTicker() *Ticker {
	return &Ticker{
		time.NewTicker(time.Second * 1),
		make(chan struct{}),
	}
}

func (t *Ticker) Tick() {
	for {
		select {
		case tick := <-t.ticker.C:
			fmt.Println(tick)
		case <-t.close:
			fmt.Println("Ticker is finished ticking!")
			return
		}
	}
}

func (t *Ticker) Stop() {
	// Stop the ticker which cleans ticker resources and to avoid any leaks
	//
	// Stop turns off a ticker.  After Stop, no more ticks will be sent.
	// Stop does not close the channel, to prevent a read from the channel succeeding
	// incorrectly.
	t.ticker.Stop()

	// Now close the custom Ticker object's close channel, to let the custom
	// Tick method's function to exit
	close(t.close)
}

func main() {
	// Create a new ticker that ticks once per second
	//
	// NewTicker returns a new Ticker containing a channel that will send the
	// time with a period specified by the duration argument.
	// It adjusts the intervals or drops ticks to make up for slow receivers.
	// The duration d must be greater than zero; if not, NewTicker will panic.
	// Stop the ticker to release associated resources.
	ticker := NewTicker()

	// Stop the ticker after 3 seconds
	go func() {
		time.Sleep(time.Second * 3)
		ticker.Stop()
	}()

	// Call tickers Tick function
	ticker.Tick()
}
