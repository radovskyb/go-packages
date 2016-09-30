package main

import (
	"archive/tar"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Create a new file to write a tar archive to
	f, err := os.Create(filepath.Join("..", "archive.tar"))
	if err != nil {
		log.Fatalln(err)
	}
	// Close the new archive file when we're done
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	// Create a new tar writer that writes to f
	tw := tar.NewWriter(f)

	// Create some files to add to the archive
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "this archive contains a readme"},
		{"tester.txt", "testing testing, one two three."},
		{"todoiscool.txt", "read packages one by one"},
	}

	// Iterate over each file in files
	for _, file := range files {
		// Create a tar header for the current file
		hdr := &tar.Header{
			Name: file.Name,             // File name
			Mode: int64(os.ModePerm),    // Read/Write/Executable
			Size: int64(len(file.Body)), // Size of body
		}

		// Write the tar header to the tar writer
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}

		// Write the file body to the tar writer
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}

	// Close the tar writer when we're done
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}
}
