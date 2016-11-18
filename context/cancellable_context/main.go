package main

import (
	"context"
	"fmt"
)

func main() {
	// Create a new cancellable context from a new blank context.
	//
	// WithCancel returns a copy of parent with a new Done channel. The returned
	// context's Done channel is closed when the returned cancel function is called
	// or when the parent context's Done channel is closed, whichever happens first.
	//
	// Canceling this context releases resources associated with it, so code should
	// call cancel as soon as the operations running in this Context complete.
	//
	// Background returns a non-nil, empty Context. It is never canceled, has no
	// values, and has no deadline. It is typically used by the main function,
	// initialization, and tests, and as the top-level Context for incoming
	// requests.
	ctx, cancel := context.WithCancel(context.Background())

	// Even though ctx should have been cancelled already, it is good
	// practice to call its cancelation function in any case.
	// Failure to do so may keep the context and its parent alive
	// longer than necessary.
	defer cancel()

	// Create a new message channel.
	msg := make(chan string)

	// Call fnc 10 times in their own goroutines.
	for i := 0; i < 10; i++ {
		go fnc(i, msg, ctx, cancel)
	}

	// Receive and print a single message from the msg channel.
	fmt.Println(<-msg)

	// Wait until the cotext is cancelled before exiting main.
	<-ctx.Done()
}

// fnc takes in an integer, message string channel, a context and also a cancel function
// for that context.
func fnc(i int, msg chan string, ctx context.Context, cancel context.CancelFunc) {
	select {
	// Check if the context was already cancelled.
	//
	// Done returns a channel that's closed when work done on behalf of this
	// context should be canceled. Done may return nil if this context can
	// never be canceled. Successive calls to Done return the same value.
	//
	// WithCancel arranges for Done to be closed when cancel is called;
	// WithDeadline arranges for Done to be closed when the deadline
	// expires; WithTimeout arranges for Done to be closed when the timeout
	// elapses.
	case <-ctx.Done():
		fmt.Printf("fnc %d cancelled\n", i)
		return
	// Send a string on the msg channel and then cancel the context so the
	// other goroutines exit cleanly.
	case msg <- fmt.Sprintf("fnc %d finished first\n", i):
		cancel()
	}
}
