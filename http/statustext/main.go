package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Get the string version for the given status code, in this case 307
	//
	// StatusText returns a text for the HTTP status code. It returns the empty
	// string if the code is unknown.
	codeString := http.StatusText(307)
	fmt.Println(codeString) // Temporary Redirect
}
