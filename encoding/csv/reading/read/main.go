package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	in := `first_name,last_name,username
"Benjamin","Radovsky",radosoft
"Randle","Persy",randpers123`

	// Create a new csv reader
	r := csv.NewReader(strings.NewReader(in))

	// Iterate through the records in the csv `input`
	for {
		// Get a single record from the input
		//
		// Read reads one record from r. The record is a slice
		// of strings with each string representing one field.
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Could not retrieve record:", err)
		}

		// Print the record
		fmt.Println(record)
	}
}
