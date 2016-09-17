package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	// Create a new Client object.
	client := &http.Client{}

	// Create a new url.Values struct
	data := url.Values{}

	// Add a value to send on the form (name="Benjamin")
	data.Add("name", "Benjamin")

	// Call the client.PostForm method and send the data to http://localhost:9000
	//
	// PostForm issues a POST to the specified URL,
	// with data's keys and values URL-encoded as the request body.
	//
	// The Content-Type header is set to application/x-www-form-urlencoded.
	// To set other headers, use NewRequest and DefaultClient.Do.
	//
	// When err is nil, resp always contains a non-nil resp.Body.
	// Caller should close resp.Body when done reading from it.
	resp, err := client.PostForm("http://localhost:9000", data)
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

	// Print out the response body's contents
	fmt.Println(string(body))
}
