package main

import (
	"fmt"
	"time"
)

func main() {
	f := func() {
		fmt.Println("Called by AfterFunc()")
	}

	// Call function f after half a second
	//
	// AfterFunc waits for the duration to elapse and then calls f
	// in its own goroutine. It returns a Timer that can
	// be used to cancel the call using its Stop method.
	timer := time.AfterFunc(time.Second/2, f)

	defer timer.Stop()

	time.Sleep(time.Second)

	fmt.Println("Finished!")
}
