package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	t := &http.Transport{}

	req, err := http.NewRequest("GET", "http://betsee.com.au", nil)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := t.RoundTrip(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	if err := resp.Write(os.Stdout); err != nil {
		log.Fatalln(err)
	}
}
