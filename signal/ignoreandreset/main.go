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
	c := make(chan os.Signal, 1)

	// Notify when an interrupt or kill signal is sent to channel c
	signal.Notify(c, os.Interrupt, os.Kill)

	// Ignore causes the provided signals to be ignored. If they are
	// received by the program, nothing will happen. Ignore undoes the
	// effect of any prior calls to Notify for the provided signals.
	// If no signals are provided, all incoming signals will be ignored.
	//
	// Ignore any kill signals being sent
	signal.Ignore(os.Kill)

	// To reset any of the effects of calling notify for any provided signals,
	// simply call signal.Reset with the signals that are not needed anymore
	// or if os.Reset is called without any provided signals, all signal handlers
	// will be reset.
	//
	// Reset undoes the effect of any prior calls to Notify for
	// the provided signals. If no signals are provided, all
	// signal handlers will be reset.
	//
	// For now leave signal.Reset commented out
	// signal.Reset(os.Interrupt)

	// Block until a signal is received.
	s := <-c

	// Print any signals that are received
	fmt.Println("\nReceived signal:", s) // Received signal: interrupt
}
