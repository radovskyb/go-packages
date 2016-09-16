package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
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
	b, err := json.Marshal(people)
	if err != nil {
		log.Fatalln("Marshal error:", err)
	}

	// Create a new bytes buffer `dst`
	dst := new(bytes.Buffer)

	// Call json indent and indent the json encoded content stored in b
	// and encode with the given prefix and indent, into bytes buffer dst
	//
	// Indent appends to dst an indented form of the JSON-encoded src.
	// Each element in a JSON object or array begins on a new,
	// indented line beginning with prefix followed by one or more
	// copies of indent according to the indentation nesting.
	// The data appended to dst does not begin with the prefix nor
	// any indentation, and has no trailing newline, to make it
	// easier to embed inside other formatted JSON data.
	if err := json.Indent(dst, b, "", "\t"); err != nil {
		log.Fatalln(err)
	}

	// Print out dst's contents to os.Stdout
	_, err = dst.WriteTo(os.Stdout)
	if err != nil {
		log.Fatalln(err)
	}
}
