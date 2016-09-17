package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Create a new get request to http://google.com and pass on the
	// requests body to the reader rdr
	//
	// NewRequest returns a new Request given a method, URL, and optional body.
	//
	// If the provided body is also an io.Closer, the returned
	// Request.Body is set to body and will be closed by the Client
	// methods Do, Post, and PostForm, and Transport.RoundTrip.
	//
	// NewRequest returns a Request suitable for use with Client.Do or
	// Transport.RoundTrip.
	// To create a request for use with testing a Server Handler use either
	// ReadRequest or manually update the Request fields. See the Request
	// type's documentation for the difference between inbound and outbound
	// request fields.
	req, err := http.NewRequest("GET", "http://betsee.com.au", nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
}
