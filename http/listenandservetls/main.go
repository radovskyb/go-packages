package main

import (
	"fmt"
	"net/http"
)

func HelloWorldTLS(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World From This Server Using TLS!"))
}

func main() {
	http.HandleFunc("/", HelloWorldTLS)
	err := http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Println(err)
	}
}
