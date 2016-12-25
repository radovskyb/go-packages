package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// Create some sample artificial input to be used to write to a writer
	const input = "This is some sample input to be written to a writer"

	// Create a new buffered writer to os.Stdout using bufio.NewWriter
	w := bufio.NewWriter(os.Stdout)

	// Use Fprintln to write the input string to writer w
	fmt.Fprintln(w, input)

	// Since we have not yet flushed the data from w to os.Stdout,
	// and input is still sitting in w, we can now use w.Reset
	// to discard any unflushed data and errors and reset w to
	// write its output to os.Stdout again (could use another io.Writer
	// instead here if we wanted to change it).
	w.Reset(os.Stdout)

	// Once again use Fprintln to write the input string to writer w
	fmt.Fprintln(w, input)

	// We now flush the buffered writer w, using w.Flush().
	// w.Flush() only prints the input string one time, since we
	// reset w after the first Fprintln, before flushing w to os.Stdout
	//
	// Flush writes any buffered data to the underlying io.Writer
	// by actually calling the writers Write() method
	if err := w.Flush(); err != nil {
		log.Fatalln(err)
	}
}
