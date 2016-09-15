package main

import (
	"fmt"
	"net/http"
)

func GetForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<form method="POST">
	<input type="text" name="name">
	<input type="submit" value="Send">
	</form>`

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
	http.HandleFunc("/", GetForm)
	http.ListenAndServe(":9000", nil)
}
