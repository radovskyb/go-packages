package main

import (
	"encoding/xml"
	"log"
	"os"
)

type Person struct {
	First string `xml:"Firstname"`
	Last  string `xml:"Lastname"`
}

func main() {
	// Create a new xml encoder that writes encoded xml to os.Stdout
	//
	// NewEncoder returns a new encoder that writes to w.
	encoder := xml.NewEncoder(os.Stdout)

	// Create a new Person object to encode using encoder
	p := &Person{"Benjamin", "Radovsky"}

	// Create a prefix and indent strings to be used with encoder.Indent
	prefix := "--> "
	indent := "    "

	// Set the encoder to generate xml that has the prefix of `prefix`
	// and indents each line with `indent`
	//
	// Indent sets the encoder to generate XML in which each element
	// begins on a new indented line that starts with prefix and is followed by
	// one or more copies of indent according to the nesting depth.
	encoder.Indent(prefix, indent)

	// Call encoder.Encode and pass it Person object p
	//
	// Encode writes the XML encoding of v to the stream.
	//
	// See the documentation for Marshal for details about the conversion
	// of Go values to XML.
	//
	// Encode calls Flush before returning.
	if err := encoder.Encode(p); err != nil {
		log.Fatalln(err)
	}
}
