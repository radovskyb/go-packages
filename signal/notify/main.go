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
	//
	// Package signal will not block sending to c: the caller must ensure that
	// c has sufficient buffer space to keep up with the expected signal rate.
	// For a channel used for notification of just one signal value, a buffer
	// of size 1 is sufficient.
	//
	// It is allowed to call Notify multiple times with the same channel:
	// each call expands the set of signals sent to that channel. The only
	// way to remove signals from the set is to call Stop.
	//
	// It is allowed to call Notify multiple times with different channels and
	// the same signals: each channel receives copies of incoming signals independently.
	c := make(chan os.Signal, 1)

	// Notify when an interrupt or kill signal is sent to channel c
	signal.Notify(c, os.Interrupt, os.Kill)

	// Block until a signal is received.
	s := <-c

	// Print any signals that are received
	fmt.Println("\nReceived signal:", s) // Received signal: interrupt
}
