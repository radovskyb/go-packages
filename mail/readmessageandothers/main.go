package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"strings"
)

func main() {
	msg := `Date: Mon, 23 Jun 2015 11:40:36 -0400
From: Gopher <from@example.com>
To: Another Gopher <to@example.com>
Subject: Gophers at Gophercon

Message body

Blah blah blah, bla
blahbedy blah...`

	r := strings.NewReader(msg)

	// ReadMessage reads a message from r.
	// The headers are parsed, and the body of the
	// message will be available  for reading from r.
	m, err := mail.ReadMessage(r)
	if err != nil {
		log.Fatalln(err)
	}

	// A Message represents a parsed mail message.
	header := m.Header

	// Use header.Date() to get the date from the messages header
	//
	// Date parses the Date header field.
	// Underneath it also simply uses header.Get("date") as part of
	// it's processing like getting the header fields below
	date, err := header.Date()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Date:", date)

	// Use header.Get(key) to get header values from the message
	//
	// Get gets the first value associated with the given key.
	// If there are no values associated with the key, Get returns "".
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	fmt.Println("Subject:", header.Get("Subject"))

	// Print a blank line before printing the message's body below
	fmt.Println()

	// Get the message's body
	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Print the message's body as a string
	fmt.Printf("%s", body)
}
