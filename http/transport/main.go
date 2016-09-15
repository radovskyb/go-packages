package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	t := &http.Transport{}

	req, err := http.NewRequest("GET", "https://www.google.com.au", nil)

	resp, err := t.RoundTrip(req)

	if err != nil {
		log.Fatalln(err)
	}

	resp.Write(os.Stdout)
}
