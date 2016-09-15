package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// The Vector type has unexported fields, which the package cannot access.
// We therefore write a BinaryMarshal/BinaryUnmarshal method pair to allow us
// to send and receive the type with the gob package. These interfaces are
// defined in the "encoding" package.
// We could equivalently use the locally defined GobEncode/GobDecoder
// interfaces.
type Vector struct {
	x, y, z int
}

// MarshalBinary returns v.x, v.y and v.z as a bytes
func (v Vector) MarshalBinary() ([]byte, error) {
	// A simple encoding: plain text
	var b bytes.Buffer
	fmt.Fprintln(&b, v.x, v.y, v.z)
	return b.Bytes(), nil
}

// UnmarshalBinary modifies the receiver so it must take a pointer receiver.
// UnmarshalBinary unmarshalls the given data into the vector object
func (v *Vector) UnmarshalBinary(data []byte) error {
	// A simple encoding: plain text
	b := bytes.NewBuffer(data)
	_, err := fmt.Fscanln(b, &v.x, &v.y, &v.z)
	return err
}

// This example transmits a value that implements the
// custom encoding and decoding methods.
func main() {
	var network bytes.Buffer // Stand-in for the network.

	// Create a new gob encoder and send a value.
	enc := gob.NewEncoder(&network)

	// Encode a new Vector with some vector values
	err := enc.Encode(Vector{3, 4, 5})
	if err != nil {
		log.Fatalln("Encode error:", err)
	}

	// Create a new gob decoder and receive a value.
	dec := gob.NewDecoder(&network)

	// Create a new Vector to store the decoded results into
	var v Vector

	// Decode into &v
	err = dec.Decode(&v)
	if err != nil {
		log.Fatalln("Decode error:", err)
	}

	// Print out Vector v
	fmt.Println(v)
}
