package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Open a zip file for reading
	r, err := zip.OpenReader(filepath.Join("..", "archive.zip"))
	if err != nil {
		log.Fatalln(err)
	}
	// Close the archive when we're done
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Iterate through the files in the archive
	for _, f := range r.File {
		// Print the current file's name
		fmt.Printf("%s:\n", f.Name)

		// Open the current file to read it's contents
		rc, err := f.Open()
		if err != nil {
			log.Fatalln(err)
		}

		// Copy the contents of the file to stdout
		_, err = io.Copy(os.Stdout, rc)
		if err != nil {
			log.Fatalln(err)
		}

		// Close the file handle
		if err := rc.Close(); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("\n")
	}
}
