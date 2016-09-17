package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	rpURL, err := url.Parse("http://localhost:9001")
	if err != nil {
		log.Fatalln(err)
	}

	// NewSingleHostReverseProxy returns a new ReverseProxy that rewrites
	// URLs to the scheme, host, and base path provided in target. If the
	// target's path is "/base" and the incoming request was for "/dir",
	// the target request will be for /base/dir.
	proxy := httputil.NewSingleHostReverseProxy(rpURL)
	log.Fatal(http.ListenAndServe(":9000", proxy))
}
