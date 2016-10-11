package main

import (
	"fmt"
	"html"
)

func main() {
	// EscapeString:
	const s = `"My Email" <radovskyb@gmail.com>`

	// Escape the special characters in string `s` and store the escaped version
	// of the string in variable `escaped` and then print it
	//
	// EscapeString escapes special characters like "<" to become "&lt;". It
	// escapes only five such characters: <, >, &, ' and ".
	// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
	// always true.
	escaped := html.EscapeString(s)
	fmt.Println(escaped)

	// UnescapeString:

	// Unescape the escaped string `escaped` from above and print the unescaped string
	//
	// UnescapeString unescapes entities like "&lt;" to become "<". It unescapes a
	// larger range of entities than EscapeString escapes. For example, "&aacute;"
	// unescapes to "á", as does "&#225;" and "&xE1;".
	// UnescapeString(EscapeString(s)) == s always holds, but the converse isn't
	// always true.
	fmt.Println(html.UnescapeString(escaped))
}
