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

	// Get issues a GET to the specified URL. If the response is one of the
	// following redirect codes, Get follows the redirect after calling the
	// Client's CheckRedirect function:
	//
	//    301 (Moved Permanently)
	//    302 (Found)
	//    303 (See Other)
	//    307 (Temporary Redirect)
	//
	// An error is returned if the Client's CheckRedirect function fails
	// or if there was an HTTP protocol error. A non-2xx response doesn't
	// cause an error.
	//
	// When err is nil, resp always contains a non-nil resp.Body.
	// Caller should close resp.Body when done reading from it.
	//
	// To make a request with custom headers, use NewRequest and Client.Do.
	resp, err := client.Get("http://betsee.com.au")
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
