package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Create a new channel c
	c := make(chan struct{})

	// Seed the random number generator
	rand.Seed(time.Now().UTC().UnixNano())

	go func() {
		// Sleep for a random duration
		//
		// Sleep pauses the current goroutine for at least the duration d.
		// A negative or zero duration causes Sleep to return immediately.
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

		// Close channel c, therefore putting a nil value into the channel
		close(c)
	}()

	// Receive from channel c if the time.Sleep using the random number
	// generated, occurs before 250 milliseconds are up, otherwise
	// call the <-time.After select case clause.
	//
	// After waits for the duration to elapse and then sends the current time
	// on the returned channel.
	// It is equivalent to NewTimer(d).C.
	select {
	case <-c:
		fmt.Println("Received from c")
	case <-time.After(250 * time.Millisecond):
		fmt.Println("Timed out!")
	}
}
