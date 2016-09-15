package main

import (
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

	http.ListenAndServe(":9000", proxy)
}
