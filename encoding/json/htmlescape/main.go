package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	dst := new(bytes.Buffer)

	src := []byte(`{
"<script>Name</script>":"Benjamin Radovsky",
"Age":24
"Job":"Qualified Awesomee"
}`)

	// Escape html in the json encoded []byte stored in `src`
	// and append the escaped json to the bytes buffer `dst`
	//
	// HTMLEscape appends to dst the JSON-encoded src with <, >, &, U+2028 and U+2029
	// characters inside string literals changed to \u003c, \u003e, \u0026, \u2028, \u2029
	// so that the JSON will be safe to embed inside HTML <script> tags.
	// For historical reasons, web browsers don't honor standard HTML
	// escaping within <script> tags, so an alternative JSON encoding must
	// be used.
	json.HTMLEscape(dst, src)

	// Print the bytes buffer dst's contents as a string
	fmt.Println(dst.String())
}
