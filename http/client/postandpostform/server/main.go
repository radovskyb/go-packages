package main

import (
	"fmt"
	"log"
	"net/http"
)

const html = `<form method="POST">
<input type="text" name="name">
<input type="submit" value="Send">
</form>`

func GetFormHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to text/html.
	w.Header().Set("Content-Type", "text/html")

	// Check if the request method is POST.
	if r.Method == "POST" {
		name := r.FormValue("name")
		if name != "" {
			fmt.Fprintln(w, "Received", name)
			return
		}
	}
	fmt.Fprintln(w, html)
}

func main() {
	http.HandleFunc("/", GetFormHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
