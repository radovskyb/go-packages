package main

import (
	"fmt"
	"html/template"
	"os"
)

const s = `"Benjamin's Email": <radovskyb@gmail.com>`

var v = []interface{}{`"Benjamin's Email"`, `: `, `<radovskyb@gmail.com>`}

func main() {
	// HTMLEscapeString returns the escaped HTML equivalent of the plain text data s.
	fmt.Printf("HTMLEscapeString:\n%s\n\n", template.HTMLEscapeString(s))

	fmt.Println("HTMLEscape:")
	// HTMLEscape writes to w the escaped HTML equivalent of the plain text data b.
	template.HTMLEscape(os.Stdout, []byte(s))

	fmt.Println("\n\nHTMLEscaper:")
	// HTMLEscaper returns the escaped HTML equivalent of the textual
	// representation of its arguments.
	fmt.Println(template.HTMLEscaper(v...))

	fmt.Println("\nJSEscapeString:")
	// JSEscapeString returns the escaped JavaScript equivalent of
	// the plain text data s.
	fmt.Println(template.JSEscapeString(s))

	fmt.Println("\nJSEscape:")
	// JSEscape writes to w the escaped JavaScript equivalent of
	// the plain text data b.
	template.JSEscape(os.Stdout, []byte(s))

	fmt.Println("\n\nJSEscaper:")
	// JSEscaper returns the escaped JavaScript equivalent of the textual
	// representation of its arguments.
	fmt.Println(template.JSEscaper(v...))

	fmt.Println("\nURLQueryEscaper:")
	// URLQueryEscaper returns the escaped value of the textual representation of
	// its arguments in a form suitable for embedding in a URL query.
	fmt.Println(template.URLQueryEscaper(v...))
}
