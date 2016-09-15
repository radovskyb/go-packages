package main

import (
	"bytes"
	"fmt"
	"index/suffixarray"
	"log"
)

func main() {
	// Create some data
	data := []byte("some data")

	// Create a new index for data
	index1 := suffixarray.New(data)

	// Create a new bytes.Buffer
	var buf bytes.Buffer

	// Write the index index1 to the buffer
	if err := index1.Write(&buf); err != nil {
		log.Fatalln("index1.Write error:", err)
	}

	// Create a second index but not yet containing any data
	var index2 suffixarray.Index

	// Read the data from index1 stored in buf, into index2
	if err := index2.Read(&buf); err != nil {
		log.Fatalln("index2.Read error:", err)
	}

	// Print out index2's bytes as a string
	fmt.Println(string(index2.Bytes()))
}
