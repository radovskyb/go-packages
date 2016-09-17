package main

import (
	"log"
	"net/http"
)

func HelloWorldTLS(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, World From This Server Using TLS!"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", HelloWorldTLS)
	log.Fatal(http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil))
}
