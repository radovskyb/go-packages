package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Create 2 random time durations between 0-100 milliseconds.
	d1 := time.Duration(rand.Intn(100)) * time.Millisecond
	d2 := time.Duration(rand.Intn(100)) * time.Millisecond

	// Print out the 2 time durations.
	fmt.Printf("d1: %s, d2: %s\n", d1, d2)

	// Create a new context with d1 as it's timeout.
	//
	// WithTimeout returns WithDeadline(parent, time.Now().Add(timeout)).
	//
	// WithDeadline returns a copy of the parent context with the deadline adjusted
	// to be no later than d. If the parent's deadline is already earlier than d,
	// WithDeadline(parent, d) is semantically equivalent to parent. The returned
	// context's Done channel is closed when the deadline expires, when the returned
	// cancel function is called, or when the parent context's Done channel is
	// closed, whichever happens first.
	//
	// 	func slowOperationWithTimeout(ctx context.Context) (Result, error) {
	// 		ctx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	// 		defer cancel()  // releases resources if slowOperation completes before timeout elapses
	// 		return slowOperation(ctx)
	// 	}
	//
	// Canceling this context releases resources associated with it, so code should
	// call cancel as soon as the operations running in this Context complete.
	ctx, cancel := context.WithTimeout(context.Background(), d1)

	// Even though ctx should have expired already, it is good
	// practice to call its cancelation function in any case.
	// Failure to do so may keep the context and its parent alive
	// longer than necessary.
	defer cancel()

	// Print a specified message based on whether the context's deadline has expired yet or not.
	select {
	case <-ctx.Done():
		fmt.Println("timed out before time.After completed")
	case <-time.After(d2):
		fmt.Println("time.After completed")
	}
}
