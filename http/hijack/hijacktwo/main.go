package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	// See if the webserver supports hijacking
	hj, ok := w.(http.Hijacker)

	// If not, return an http error
	if !ok {
		http.Error(w, "webserver doesn't support hijacking",
			http.StatusInternalServerError)
		return
	}

	// Hijack the connection
	conn, bufrw, err := hj.Hijack()
	// Check if there were any errors and if so return them as an http error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// defer to close the hijacked connection
	defer conn.Close()

	// Write a string to the hijacked connection
	bufrw.WriteString("Message sent over raw TCP")

	// Flush the message from the buffered writer onto the connection
	bufrw.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
