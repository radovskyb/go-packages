package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// Create records to be written to the csv.Writer
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Benjamin", "Radovsky", "radosoft"},
		{"Randly", "Persy", "ranpers123"},
	}

	// Create a new csv.Writer w to write csv encoded
	// records to os.Stdout
	w := csv.NewWriter(os.Stdout)

	// Write all of the records to the csv.Writer w which also calls
	// flush internally as opposed to having to call it manually
	//
	// WriteAll writes multiple CSV records to w using Write
	// and then calls Flush.
	if err := w.WriteAll(records); err != nil {
		log.Fatalln(err)
	}

	// Check if w contains any errors and if so, log them
	if err := w.Error(); err != nil {
		log.Fatalln(err)
	}
}
