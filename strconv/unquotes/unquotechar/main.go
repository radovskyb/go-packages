package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// UnquoteChar decodes the first character or byte in the
	// escaped string or character literal represented by the string s.
	// It returns four values:
	//
	//	1) value, the decoded Unicode code point or byte value;
	//	2) multibyte, a boolean indicating whether the decoded
	//     character requires a multibyte UTF-8 representation;
	//	3) tail, the remainder of the string after the character; and
	//	4) an error that will be nil if the character is syntactically
	//     valid.
	//
	// The second argument, quote, specifies the type of literal being
	// parsed and therefore which escaped quote character is permitted.
	//
	// If set to a single quote, it permits the sequence \' and
	// disallows unescaped '. If set to a double quote, it permits
	// \" and disallows unescaped ".
	//
	// If set to zero, it does not permit either escape and allows
	// both quote characters to appear unescaped.
	v, mb, t, err := strconv.UnquoteChar(`\"I'm inside of quotes\"`, '"')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value:", string(v))
	fmt.Println("multibyte:", mb)
	fmt.Println("tail:", t)
}
