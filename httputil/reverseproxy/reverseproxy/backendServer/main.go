package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Received from localhost:9001 using a reverse proxy")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9001", nil)
}
