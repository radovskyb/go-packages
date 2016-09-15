package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
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

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	fmt.Println(string(body))
}
