package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	First string `xml:"Firstname"`
	Last  string `xml:"Lastname"`
}

func main() {
	// Create a new Person object `p` to be marshalled into xml
	p := &Person{First: "Benjamin", Last: "Radovsky"}

	// Create a prefix and indent strings to be used with xml.MarshalIndent
	prefix := "--> "
	indent := "    "

	// Marshal `p` into xml and have each line start with the prefix `prefix`
	// and have each line's indent be set to `indent`
	//
	// MarshalIndent works like Marshal, but each XML element begins on a new
	// indented line that starts with prefix and is followed by one or more
	// copies of indent according to the nesting depth.
	marshalled, err := xml.MarshalIndent(p, prefix, indent)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the marshalled xml as a string
	fmt.Println(string(marshalled))
}
