package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
)

func main() {
	const body = "Go is a general-purpose language designed with systems programming in mind."

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Date", "Wed, 19 Jul 1972 19:00:00 GMT")
			fmt.Fprintln(w, body)
		}),
	)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	// DumpResponse is like DumpRequest but dumps a response.
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%q", dump)
}
