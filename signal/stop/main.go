package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	// Notify causes package signal to relay incoming signals to c. If no
	// signals are provided, all incoming signals will be relayed to c.
	// Otherwise, just the provided signals will.
	c := make(chan os.Signal, 1)

	// Notify when an interrupt or kill signal is sent to channel c
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-c

	// Stop relaying any incoming signals to chan c
	//
	// Stop causes package signal to stop relaying incoming signals to c.
	// It undoes the effect of all prior calls to Notify using c. When
	// Stop returns, it is guaranteed that c will receive no more signals.
	signal.Stop(c)

	// Print any signals that are received
	fmt.Println("Got signal:", s) // Wont print because signal.Stop() was called
}
