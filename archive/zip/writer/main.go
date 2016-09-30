package main

import (
	"archive/zip"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create a new file to write a zip archive to
	f, err := os.Create(filepath.Join("..", "archive.zip"))
	if err != nil {
		log.Fatalln(err)
	}
	// Close the new archive file when we're done
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Create a new zip writer that writes to f
	w := zip.NewWriter(f)

	// Create some files to add to the archive
	var files = []struct {
		Name, Body string
	}{
		{"zipheaderone.txt", "zip content one"},
		{"zipheadertwo.txt", "zip content two"},
	}

	// Iterate over each file in files
	for _, file := range files {
		// Add a file to the zip archive using the current file's name
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatalln(err)
		}

		// Write the file's content to the archive
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatalln(err)
		}
	}

	// Close the zip writer when we're done
	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}
}
