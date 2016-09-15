package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	dst := new(bytes.Buffer)

	src := []byte(`{
"Name":"Benjamin Radovksky",
"Age":24,
"Job":"Being Awesome"
}`)

	// Append the json encoded string in `src` into bytes.Buffer `dst`
	//
	// Compact appends to dst the JSON-encoded src with
	// insignificant space characters elided.
	err := json.Compact(dst, src)
	if err != nil {
		log.Fatalln("JSON compact error:", err)
	}

	// Print the bytes buffer dst's contents as a string
	fmt.Println(dst.String())
}
