package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

const jsonText = `{"Name":"Benjamin Radovksky","Age":24}`

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a new json.Decoder that reads from a strings reader that
	// reads from `jsonText`
	//
	// NewDecoder returns a new decoder that reads from r.
	//
	// The decoder introduces its own buffering and may
	// read data from r beyond the JSON values requested.
	decoder := json.NewDecoder(strings.NewReader(jsonText))

	// Create a new Person object to decode jsonText into
	person := new(Person)

	// Decode `jsonText` into `person`
	//
	// Decode reads the next JSON-encoded value from its
	// input and stores it in the value pointed to by v.
	//
	// See the documentation for Unmarshal for details about
	// the conversion of JSON into a Go value.
	if err := decoder.Decode(person); err != nil {
		log.Fatalln("Decode error:", err)
	}

	// Print out `person`'s Name and Age fields
	fmt.Println(person.Name)
	fmt.Println(person.Age)
}
