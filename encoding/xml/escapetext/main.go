package main

import (
	"encoding/xml"
	"log"
	"os"
)

const text string = "\".&<>"

func main() {
	// EscapeText writes to w the properly escaped XML equivalent
	// of the plain text data s.
	//
	// &#34;.&amp;&lt;&gt;
	if err := xml.EscapeText(os.Stdout, []byte(text)); err != nil {
		log.Fatalln(err)
	}
}
