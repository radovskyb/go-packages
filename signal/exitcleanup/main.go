package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a buffered channel os type os.Signal
	c := make(chan os.Signal, 1)

	// Create a channel of type bool for storing whether the
	// cleanup after catching any signals has completed or not
	done := make(chan bool)

	// Use signal.Nofity to listen for either interrupt or terminate signals
	// which get relayed to channel c
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Run this go routine that will take care of signal handling
	// and what to do once one is received
	go func() {
		// Wait for a signal to be received and store it's type in a the sig
		sig := <-c

		// Print out the type of signal received
		fmt.Println("\nReceived signal:", sig)

		// Now clean up any resources that need to be freed (Demo function only)
		CleanupResources()

		// Once CleanupResources has finished execution, pass true on to the done chan
		done <- true
	}()

	// Print out a message to let the user know how to quit the application
	fmt.Println("Press Ctrl-c to send interrupt signal")

	// Wait for done to receive a value
	<-done

	// Once done has received it's value, print an exit message
	fmt.Println("Finished cleaning up. Exiting.")
}

// CleanupResources is just a demo function of what might be called
// when a cleanup might be called before the program exits. This could contain
// such things as pushing files to memory or printing out messages like below etc.
func CleanupResources() {
	// Print out a cleaning up resources message
	fmt.Println("Cleaning up resources...")

	// Simulate resource cleaning
	time.Sleep(time.Second)
}
