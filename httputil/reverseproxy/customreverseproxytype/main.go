package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type CustomProxy struct {
	*httputil.ReverseProxy
	Target string
}

func (cp *CustomProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Redirecting to %s\n", cp.Target)
	cp.ReverseProxy.ServeHTTP(w, r)
}

func NewCustomProxy(target string) *CustomProxy {
	director := func(req *http.Request) {
		url, err := url.Parse(target)
		if err != nil {
			log.Fatalln("An error occured:", err.Error())
		}

		req.URL.Scheme = url.Scheme
		req.URL.Host = url.Host
		req.Host = url.Host
	}

	return &CustomProxy{
		ReverseProxy: &httputil.ReverseProxy{
			Director: director,
		},
		Target: target,
	}
}

func main() {
	proxy := NewCustomProxy("http://google.com")
	http.ListenAndServe(":9000", proxy)
}
