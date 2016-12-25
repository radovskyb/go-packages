package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "An error occurred", http.StatusInternalServerError)
	}

	req, err := http.NewRequest("GET", "http://betsee.com.au", nil)
	if err != nil {
		log.Fatal(err)
	}

	// NewRecorder returns an initialized ResponseRecorder.
	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
}
