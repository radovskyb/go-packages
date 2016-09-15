package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
)

func main() {
	backendServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "This call was relayed by the reverse proxy.")
		}),
	)
	defer backendServer.Close()

	rpURL, err := url.Parse(backendServer.URL)
	if err != nil {
		log.Fatalln(err)
	}

	frontendProxy := httptest.NewServer(
		httputil.NewSingleHostReverseProxy(rpURL),
	)
	defer frontendProxy.Close()

	resp, err := http.Get(frontendProxy.URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", body)
}
