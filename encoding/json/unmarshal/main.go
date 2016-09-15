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
	jsonString := `[{"First":"Benjamin","Last":"Radovsky"},{"First":"Benji","Last":"Raddy"}]`

	// Create a new Person object to unmarshal `jsonString` into
	people := []Person{}

	// Unmarshal jsonString as a byte slice into Person object `person`
	//
	// Unmarshal parses the JSON-encoded data and stores the result
	// in the value pointed to by v.
	//
	// Unmarshal uses the inverse of the encodings that
	// Marshal uses, allocating maps, slices, and pointers as necessary,
	err := json.Unmarshal([]byte(jsonString), &people)
	if err != nil {
		log.Fatalln("Unmarshal error:", err)
	}

	// Print out `people` object
	fmt.Println("People - ", people)

	// Print out all of the `Person` objects stored in `people`
	for i, person := range people {
		// Print out each person's first and last names
		fmt.Println("\nPerson", i, "-", person)
		fmt.Println("Firstname:", person.First)
		fmt.Println("Lastname:", person.Last)
	}
}
