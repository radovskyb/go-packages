package main

import (
	"encoding/xml"
	"os"
)

const text string = "\".&<>"

func main() {
	// EscapeText writes to w the properly escaped XML equivalent
	// of the plain text data s.
	xml.EscapeText(os.Stdout, []byte(text)) // &#34;.&amp;&lt;&gt;
}
