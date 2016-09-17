package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var url string
	var file string

	// Prepare a form that you will submit to that URL.
	var buf bytes.Buffer
	// NewWriter returns a new multipart Writer with a random boundary,
	// writing to w.
	w := multipart.NewWriter(&buf)

	// Add an image file
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// CreateFormFile is a convenience wrapper around CreatePart. It creates
	// a new form-data header with the provided field name and file name.
	fw, err := w.CreateFormFile("image", file)
	if err != nil {
		log.Fatalln(err)
	}

	// Copy file f's contents into field writer fw
	if _, err = io.Copy(fw, f); err != nil {
		log.Fatalln(err)
	}

	// Add the other fields
	//
	// CreateFormField calls CreatePart with a header using the
	// given field name.
	if fw, err = w.CreateFormField("key"); err != nil {
		log.Fatalln(err)
	}
	if _, err = fw.Write([]byte("KEY")); err != nil {
		log.Fatalln(err)
	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	//
	// Close finishes the multipart message and writes the trailing
	// boundary end line to the output.
	if err := w.Close(); err != nil {
		log.Fatalln(err)
	}

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &buf)
	if err != nil {
		log.Fatalln(err)
	}

	// Don't forget to set the content type, this will contain the boundary.
	//
	// FormDataContentType returns the Content-Type for an HTTP
	// multipart/form-data with this Writer's Boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	if err := res.Body.Close(); err != nil {
		log.Fatalln(err)
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Bad status: %s", res.Status)
	}
}
