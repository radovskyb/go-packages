package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"strings"
)

func main() {
	in := `first_name;last_name;username
"Benjamin";"Radovsky";radosoft
# This is a comment and will be ignored
"Randle";"Persy";randpers123
# This line will be ignored too.`

	// Create a new csv reader
	r := csv.NewReader(strings.NewReader(in))

	// Set some csv options:

	// Use a semicolon `;` as a field delimiter instead of a comma `,`
	r.Comma = ';'

	// Set the comment character for start of line to a hash `#`
	r.Comment = '#'

	// Read all records from input in
	//
	// ReadAll reads all the remaining records from r.
	// Each record is a slice of fields.
	// A successful call returns err == nil, not err == EOF.
	// Because ReadAll is defined to read until EOF, it does
	// not treat end of file as an error to be reported.
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	// Print out all records as a [][]string
	fmt.Println(records)

	fmt.Println()

	// Loop through records and print out each
	// record individuall yon it's own line
	for _, record := range records {
		fmt.Println(record)

		// To further extract record values out of the record
		// slice, another range loop could be used. example:
		// for _, v := range record {
		// 	fmt.Println(v)
		// }
	}
}
