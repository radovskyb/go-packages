package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	// Unescape the escaped query and store the results in string var unescaped.
	unescaped, err := url.QueryUnescape("name%3DBenjamin+Radovsky%26age%3D24")
	if err != nil {
		log.Fatalln(err)
	}

	// Print the unescaped query.
	fmt.Println(unescaped)
}
