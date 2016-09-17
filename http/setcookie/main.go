package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	// Create a new cookie
	cookie := &http.Cookie{Name: "loggedin", Value: "true",
		Expires: time.Now().Add(time.Hour * 24)}

	// Set cookie `loggedin`
	//
	// SetCookie adds a Set-Cookie header to the provided ResponseWriter's headers.
	// The provided cookie must have a valid Name. Invalid cookies may be
	// silently dropped.
	http.SetCookie(w, cookie)

	// Print out that the cookie has been set
	fmt.Fprintln(w, "Set cookie loggedin to value true")
}

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	// Try to get the cookie `loggedin` and if theres an error
	// then write it to the user and return
	cookie, err := r.Cookie("loggedin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print the cookie's value
	fmt.Fprintln(w, "Cookie loggedin is set to", cookie.Value) // True if set
}

func main() {
	http.HandleFunc("/", HandlerOne)
	http.HandleFunc("/two", HandlerTwo)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
