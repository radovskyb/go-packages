package main

import (
	"encoding/xml"
	"os"
)

const text string = "\"'&<>"

func main() {
	// Escape is like EscapeText but omits the error return value.
	// It is provided for backwards compatibility with Go 1.0.
	// Code targeting Go 1.1 or later should use EscapeText.
	xml.Escape(os.Stdout, []byte(text)) // &#34;&#39;&amp;&lt;&gt;$
}
