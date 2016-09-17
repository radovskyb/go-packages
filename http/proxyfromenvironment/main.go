package main

import (
	"log"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	html := "Hello World!\n"

	// ProxyFromEnvironment returns the URL of the proxy to use for a
	// given request, as indicated by the environment variables
	// HTTP_PROXY, HTTPS_PROXY and NO_PROXY (or the lowercase versions
	// thereof). HTTPS_PROXY takes precedence over HTTP_PROXY for https
	// requests.
	//
	// The environment values may be either a complete URL or a
	// "host[:port]", in which case the "http" scheme is assumed.
	// An error is returned if the value is a different form.
	//
	// A nil URL and nil error are returned if no proxy is defined in the
	// environment, or a proxy should not be used for the given request,
	// as defined by NO_PROXY.
	//
	// As a special case, if req.URL.Host is "localhost" (with or without
	// a port number), then a nil URL and nil error will be returned.
	url, err := http.ProxyFromEnvironment(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if url != nil {
		html = html + "scheme " + url.Scheme + "\n"
		html = html + "host " + url.Scheme + "\n"
	} else {
		html = html + "proxy returns empty url\n"
	}

	_, err = w.Write([]byte(html))
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
