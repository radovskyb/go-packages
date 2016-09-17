package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10000); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	formData := r.MultipartForm

	if formData == nil {
		fmt.Fprintln(w, "No files uploaded!")
	} else {
		// Remove any temporary file associated with formData when the
		// function exits
		//
		// RemoveAll removes any temporary files associated with a Form.
		defer formData.RemoveAll()
	}
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
