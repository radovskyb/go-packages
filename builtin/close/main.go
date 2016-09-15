package main

import "fmt"

func main() {
	channel := make(chan int, 1)
	channel <- 1

	x, ok := <-channel
	fmt.Println(x, ok)

	// The close built-in function closes a channel, which must be
	// either bidirectional or send-only. It should be executed only
	// by the sender, never the receiver, and has the effect of
	// shutting down the channel after the last sent value is received.
	// After the last value has been received from a closed channel c,
	// any receive from c will succeed without blocking, returning the
	// zero value for the channel element. The form will also set ok
	// to false for a closed channel.
	close(channel)

	x, ok = <-channel
	fmt.Println(x, ok)
}
