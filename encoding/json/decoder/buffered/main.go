package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Person struct {
	Name string
}

func main() {
	// Create a new strings reader that contains a json encoded string as
	// well as some extra text remaining after the json string
	reader := strings.NewReader(`{"Name": "Benjamin"}[Extra string to be buffered!]`)

	// Create a new json.Decoder
	decoder := json.NewDecoder(reader)

	// Create a new Person object to decode into
	person := new(Person)

	// Decode the json values from the decoder into `person`
	if err := decoder.Decode(person); err != nil {
		log.Fatalln("Decode error:", err)
	}

	// Print out person
	fmt.Println(person)

	// Read all of the remaining string from `reader` after
	// the json encoded part has been decoded into person above by
	// reading from reader returned from decoder.Buffered which returns
	// a reader containing only the remaining string after the json
	// encoded string has been decoded
	//
	// Buffered returns a reader of the data remaining in the Decoder's
	// buffer. The reader is valid until the next call to Decode.
	if extraString, err := ioutil.ReadAll(decoder.Buffered()); err != nil {
		log.Fatalln("ReadAll error:", err)
	} else {
		fmt.Println(string(extraString))
	}

}
