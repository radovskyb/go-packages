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

	// Marshal people into json and store the results in byte slice b
	//
	// Marshal returns the JSON encoding of v.
	b, err := json.Marshal(people)
	if err != nil {
		log.Fatalln("Marshal error:", err)
	}

	// Print out b as a string
	fmt.Println(string(b))
}
