package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = "google.com"
			req.Host = "google.com"
		},
	}

	log.Fatal(http.ListenAndServe(":9000", proxy))
}
