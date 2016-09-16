package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func main() {
	// Create a new cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Create a cookie slice to store cookies
	var cookies []*http.Cookie

	// Create some new cookies
	cookieOne := &http.Cookie{
		Name:   "C1",
		Value:  "cookieone",
		Path:   "/",
		Domain: "localhost",
	}

	cookieTwo := &http.Cookie{
		Name:   "C2",
		Value:  "cookietwo",
		Path:   "/",
		Domain: "localhost",
	}

	cookieThree := &http.Cookie{
		Name:   "C3",
		Value:  "cookiethree",
		Path:   "/",
		Domain: "localhost",
	}

	// Append the 3 cookies into the cookies slice
	cookies = append(cookies, cookieOne)
	cookies = append(cookies, cookieTwo)
	cookies = append(cookies, cookieThree)

	// Parse the url of the server localhost:9000
	cookieURL, err := url.Parse("http://localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}

	// Add the cookies slice to the cookie jar
	jar.SetCookies(cookieURL, cookies)

	// Print out the contents of the cookie jar for comparison
	fmt.Println(jar.Cookies(cookieURL))

	// Set up a client with the current cookie jar implemented
	client := &http.Client{
		Jar: jar,
	}

	// Send a get request over client to localhost:9000
	// which will carry the cookiejar over with it
	resp, err := client.Get("http://localhost:9000")
	if err != nil {
		log.Fatalln(err)
	}
	// Close the response body after it's been read
	defer resp.Body.Close()

	// Read the response's body from the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the response body as a string
	fmt.Println(string(body))
}
