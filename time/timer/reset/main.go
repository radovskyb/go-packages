package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(5 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer expired!")
	}()

	// Change the timer to end after half a second instead of 5 seconds
	//
	// Reset changes the timer to expire after duration d.
	// It returns true if the timer had been active, false if the timer had
	// expired or been stopped.
	reset := timer.Reset(time.Second / 2)

	time.Sleep(time.Duration(1) * time.Second)

	fmt.Println("Did the timer reset from 5 seconds to a half a second?:", reset)
}
