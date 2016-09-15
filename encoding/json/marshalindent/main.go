package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	First string
	Last  string
}

func main() {
	people := []*Person{
		{"Benjamin", "Radovsky"},
		{"Benji", "Raddy"},
	}

	// Create a prefix and indent strings to use with json.MarshalIndent
	prefix := "-->"
	indent := "\t"

	// Marshal people into json with `prefix` and `indent` used as the
	// encoded prefix and indentation and store the results in byte slice b
	//
	b, err := json.MarshalIndent(people, prefix, indent)
	if err != nil {
		log.Fatalln("MarshalIndent error:", err)
	}

	// Print out b as a string
	fmt.Println(string(b))
}
