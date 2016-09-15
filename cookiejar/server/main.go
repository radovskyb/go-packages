package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mw := io.MultiWriter(os.Stdout, w)

	cookies := r.Cookies()
	for _, cookie := range cookies {
		fmt.Fprintln(mw, "Received cookie: "+cookie.Value)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
