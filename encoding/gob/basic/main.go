package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

// This example shows the basic usage of the package: Create an encoder,
// transmit some values, receive them with a decoder.
func main() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer // Stand-in for a network connection

	// NewEncoder returns a new encoder that will transmit on the io.Writer.
	enc := gob.NewEncoder(&network) // Will write to a connection

	// NewDecoder returns a new decoder that reads from the io.Reader.
	// If r does not also implement io.ByteReader, it will be wrapped in a
	// bufio.Reader.
	dec := gob.NewDecoder(&network) // Will read to a connection

	// Encode and send some values.
	//
	// Encode transmits the data item represented by the empty
	// interface value, guaranteeing that all necessary type
	// information has been transmitted first.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatalln("Encode error:", err)
	}

	err = enc.Encode(P{30, 40, 50, "Theorem"})
	if err != nil {
		log.Fatalln("Encode error:", err)
	}

	// Receive and decode some values.
	var q Q

	// Decode reads the next value from the input stream and stores
	// it in the data represented by the empty interface value.
	// If e is nil, the value will be discarded. Otherwise,
	// the value underlying e must be a pointer to the
	// correct type for the next data item received.
	// If the input is at EOF, Decode returns io.EOF and
	// does not modify e.
	err = dec.Decode(&q)
	if err != nil {
		log.Fatalln("Decode error:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)

	err = dec.Decode(&q)
	if err != nil {
		log.Fatalln("Decode error:", err)
	}
	fmt.Printf("%q: {%d, %d}\n", q.Name, *q.X, *q.Y)
}
