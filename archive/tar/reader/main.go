package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Open a tar archive for reading
	r, err := os.Open(filepath.Join("..", "archive.tar"))
	if err != nil {
		log.Fatalln(err)
	}
	// Close the archive when we're done
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Create a tar reader that reads from r
	tr := tar.NewReader(r)

	// Iterate through the files in the archive
	for {
		// Advance to the next entry in the tar archive
		hdr, err := tr.Next()

		// If err is io.EOF, we've reached the end of
		// the archive so exit the loop
		if err == io.EOF {
			break
		}

		// If the error is something else, log it
		if err != nil {
			log.Fatalln(err)
		}

		// Print the name of the current file
		fmt.Printf("%s:\n", hdr.Name)

		// Print the contents of the current file
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}

		fmt.Println("\n")
	}
}
