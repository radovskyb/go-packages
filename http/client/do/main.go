package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Do sends an HTTP request and returns an HTTP response, following
	// policy (e.g. redirects, cookies, auth) as configured on the client.
	//
	// An error is returned if caused by client policy (such as
	// CheckRedirect), or if there was an HTTP protocol error.
	// A non-2xx response doesn't cause an error.
	//
	// When err is nil, resp always contains a non-nil resp.Body.
	//
	// Callers should close resp.Body when done reading from it. If
	// resp.Body is not closed, the Client's underlying RoundTripper
	// (typically Transport) may not be able to re-use a persistent TCP
	// connection to the server for a subsequent "keep-alive" request.
	//
	// The request Body, if non-nil, will be closed by the underlying
	// Transport, even on errors.
	//
	// Generally Get, Post, or PostForm will be used instead of Do.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println(string(body))
}
