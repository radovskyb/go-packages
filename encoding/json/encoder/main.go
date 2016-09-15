package main

import (
	"encoding/json"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new json.Encoder
	//
	// NewEncoder returns a new encoder that writes to w.
	encoder := json.NewEncoder(os.Stdout)

	// Encode a Person object into json using `encoder`
	//
	// Encode writes the JSON encoding of v to the stream,
	// followed by a newline character.
	//
	// See the documentation for Marshal for details about the
	// conversion of Go values to JSON.
	if err := encoder.Encode(&Person{"Benjamin", 24}); err != nil {
		log.Fatalln("Encode error:", err)
	}
}
