package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Create a new Client object.
	client := &http.Client{}

	// Post issues a POST to the specified URL.
	//
	// Caller should close resp.Body when done reading from it.
	//
	// If the provided body is an io.Closer, it is closed after the
	// request.
	//
	// To set custom headers, use NewRequest and Client.Do.
	resp, err := client.Post("http://localhost:9000", "text/html", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Slurp up the response's body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// Close the response's body.
	if err := resp.Body.Close(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))
}
