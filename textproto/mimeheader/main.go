package main

import (
	"fmt"
	"net/textproto"
)

func main() {
	// Create a textproto.MIMEHeader headers map `h`
	h := textproto.MIMEHeader{}

	// Add a content-type header to the header `h` map
	h.Set("Content-Type", "text/html")

	// Get the content-type from headers map `h`
	ct := h.Get("content-type")

	// Print out the content-type header value returned from h.Get
	fmt.Println(ct)

	// Set a content-id mimeheader key
	h.Set("content-id", "123")

	// Get and print the content-id key
	fmt.Println(h.Get("CONTENT-ID"))

	// Delete the content-id mimeheader key from `h`
	h.Del("content-ID")

	// Try to get and print the content-id key
	// which will no longer be available in `h`
	// after deleting it above
	fmt.Println(h.Get("content-id"))
}
