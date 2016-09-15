package main

import (
	"fmt"
	"log"
	"net/mail"
)

func main() {
	// Create an email list to parse
	const list = "Benji <radovskyb@gmail.com>, ExamplePerson <example@gmail.com>"

	// ParseAddressList parses the given string as a list of addresses.
	addresses, err := mail.ParseAddressList(list)
	if err != nil {
		log.Fatalln(err)
	}

	// Range over the addresses and print out each name and email address
	//
	// Each iteration returns an Address type
	//
	// Address represents a single mail address.
	// An address such as "Barry Gibbs <bg@example.com>" is represented
	// as Address{Name: "Barry Gibbs", Address: "bg@example.com"}.
	for _, address := range addresses {
		fmt.Println(address.Name, "-", address.Address)
	}

	fmt.Println()

	// ParseAddress Parses a single RFC 5322 address,
	// e.g. "Barry Gibbs <bg@example.com>"
	address, err := mail.ParseAddress("Benjamin <radovskyb@gmail.com>")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(address.Name, "-", address.Address)

	// The address can also be printed using it's `String()` method
	//
	// String formats the address as a valid RFC 5322 address.
	// If the address's name contains non-ASCII characters the
	// name will be rendered according to RFC 2047.
	fmt.Println(address) // "Benjamin" <radovskyb@gmail.com>
}
