package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"
)

func main() {
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// DumpRequest returns the as-received wire
			// representation of req, optionally including the
			// request body, for debugging.
			// DumpRequest is semantically a no-op, but in order to
			// dump the body, it reads the body data into memory and
			// changes req.Body to refer to the in-memory copy.
			// The documentation for http.Request.Write details which fields
			// of req are used.
			dump, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "%q", dump)
		}),
	)
	defer ts.Close()

	const body = "Go is a general-purpose language designed with systems programming in mind."

	req, err := http.NewRequest("POST", ts.URL, strings.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	req.Host = "http://betsee.com.au"

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", b)
}
