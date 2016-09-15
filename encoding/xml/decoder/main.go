package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	First string `xml:"Firstname"`
	Last  string `xml:"Lastname"`
}

const xmlInput string = `<Person><Firstname>Benjamin</Firstname><Lastname>Radovsky</Lastname></Person>`

func main() {
	// Create a new bytes reader to read from string `xmlInput`
	xmlReader := bytes.NewReader([]byte(xmlInput))

	// Create a new xml decoder for `xmlReader`
	//
	// NewDecoder creates a new XML parser reading from r.
	// If r does not implement io.ByteReader, NewDecoder will
	// do its own buffering.
	decoder := xml.NewDecoder(xmlReader)

	// Create a new Person object to decode from `decoder` into
	person := &Person{}

	// Decode the xml from decoder into the address of Person object `person`
	//
	// Decode works like xml.Unmarshal, except it reads the decoder
	// stream to find the start element.
	err := decoder.Decode(person)
	if err != nil {
		log.Fatalln("Decode error:", err)
	}

	// Print out `person`
	fmt.Println(person)

	// Print out `person`'s first and last name strings
	fmt.Println(person.First)
	fmt.Println(person.Last)
}
