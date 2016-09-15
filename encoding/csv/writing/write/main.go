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
	//
	// NewWriter returns a new Writer that writes to w.
	w := csv.NewWriter(os.Stdout)

	// Iterate over each of the records above and write
	// each one to the csv.Writer w
	for _, record := range records {
		// Write record to csv.Writer w
		//
		// Writer writes a single CSV record to w along with
		// any necessary quoting. A record is a slice of
		// strings with each string being one field.
		if err := w.Write(record); err != nil {
			log.Fatalln("Could not write record:", err)
		}
	}

	// Flush any buffered data to the underlying writer (os.Stdout)
	//
	// Flush writes any buffered data to the underlying io.Writer.
	// To check if an error occurred during the Flush, call Error.
	w.Flush()

	// Check if w contains any errors and if so, log them
	//
	// Error reports any error that has occurred during a
	// previous Write or Flush.
	if err := w.Error(); err != nil {
		log.Fatalln(err)
	}
}
