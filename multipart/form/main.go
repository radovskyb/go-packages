package main

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10000)

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
	log.Fatalln(http.ListenAndServe(":9000", nil))
}
