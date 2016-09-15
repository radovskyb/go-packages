package main

import (
	"fmt"
	"time"
)

func main() {
	// NewTimer creates a new Timer that will send
	// the current time on its channel after at least duration d.
	timer := time.NewTimer(time.Second * 2)

	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()

	time.Sleep(time.Second)

	// Stop prevents the Timer from firing.
	// It returns true if the call stops the timer, false if the timer has already
	// expired or been stopped.
	// Stop does not close the channel, to prevent a read from the channel succeeding
	// incorrectly.
	stop := timer.Stop()

	fmt.Println("Timer stopped before expired:", stop)

	fmt.Println()

	timer = time.NewTimer(time.Second / 2)

	go func() {
		<-timer.C
		fmt.Println("Timer expired")
	}()

	time.Sleep(time.Second)

	// Stop prevents the Timer from firing.
	// It returns true if the call stops the timer, false if the timer has already
	// expired or been stopped.
	// Stop does not close the channel, to prevent a read from the channel succeeding
	// incorrectly.
	stop = timer.Stop()

	fmt.Println("Timer stopped before expired:", stop)
}
