package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	// Create a new replacer `rep`
	rep := strings.NewReplacer("<", "&lt;", ">", "&gt;")

	// Call rep.WriteString which replaces all characters
	// according to the old new string pairs from replacer
	// `rep`, and then writes the new replaced string to the
	// given writer, in this case, os.Stout
	//
	// WriteString writes s to w with all replacements performed.
	_, err := rep.WriteString(os.Stdout, "This is <b>HTML<b>\n")
	if err != nil {
		log.Fatalln(err)
	}
}
