package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var username = []byte("admin")
var password = []byte("password")

func BasicAuth(w http.ResponseWriter, r *http.Request, user, pass []byte) bool {
	// Get the username and password from the request's header
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	// Check if we got a username and password
	if len(s) != 2 {
		// If not, return false
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	if pair[0] == string(user) && pair[1] == string(pass) {
		return true
	}

	fmt.Println("User has failed to log in using the following credentials:")
	fmt.Printf("Username: %s\nPassword: %s\n", pair[0], pair[1])
	return false
}

func Home(w http.ResponseWriter, r *http.Request) {
	if BasicAuth(w, r, username, password) {
		_, err := w.Write([]byte("Protected Area\n"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// BasicAuth returns the username and password provided in the request's
		// Authorization header, if the request uses HTTP Basic Authentication.
		// See RFC 2617, Section 2.
		username, password, _ := r.BasicAuth()
		fmt.Println("User has successfully logged in with the following credentials:")
		fmt.Printf("Username: %s\nPassword: %s\n", username, password)
		return
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="Protected Area"`)
	w.WriteHeader(401)
	_, err := w.Write([]byte("401 Unauthorized\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
