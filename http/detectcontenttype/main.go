package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Make a get request to http://google.com and store
	// it's response in the variable resp
	resp, err := http.Get("http://betsee.com.au")
	if err != nil {
		log.Fatalln(err)
	}

	// Read all of the response's body's contents into variable body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Close the response's body now that we are done reading from it
	if err := resp.Body.Close(); err != nil {
		log.Fatalln(err)
	}

	// Get the content type from the response's body and store it as
	// a string in the variable contentType
	//
	// DetectContentType implements the algorithm described
	// at http://mimesniff.spec.whatwg.org/ to determine the
	// Content-Type of the given data.  It considers at most the
	// first 512 bytes of data.  DetectContentType always returns
	// a valid MIME type: if it cannot determine a more specific one, it
	// returns "application/octet-stream".
	contentType := http.DetectContentType(body)

	// Print out the content type of the response body
	fmt.Println(contentType) // text/html; charset=utf-8
}
