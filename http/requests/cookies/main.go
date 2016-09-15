package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create a new request
	req, err := http.NewRequest("GET", "http://localhost", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Create two cookies
	cookie1 := &http.Cookie{
		Name:   "C1",
		Value:  "Cookie One",
		Domain: "localhost",
		Path:   "/",
	}

	cookie2 := &http.Cookie{
		Name:   "C2",
		Value:  "Cookie Two",
		Domain: "localhost",
		Path:   "/",
	}

	// Add both of the cookies to the request
	//
	// AddCookie adds a cookie to the request.  Per RFC 6265 section 5.4,
	// AddCookie does not attach more than one Cookie header field.  That
	// means all cookies, if any, are written into the same line,
	// separated by semicolon.
	req.AddCookie(cookie1)
	req.AddCookie(cookie2)

	// Print out all of the request's cookies using req.Cookies()
	//
	// Cookies parses and returns the HTTP cookies sent with the request.
	fmt.Println(req.Cookies())

	// Print out one single cookie from the req using req.Cookie(name)
	//
	// Cookie returns the named cookie provided in the request or
	// ErrNoCookie if not found.
	fmt.Println(req.Cookie("C1"))
}
