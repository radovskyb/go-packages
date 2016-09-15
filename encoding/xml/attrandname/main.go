package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	// An Attr represents an attribute in an XML element (Name=Value).
	//
	// A Name represents an XML name (Local) annotated
	// with a name space identifier (Space).
	// In tokens returned by Decoder.Token, the Space identifier
	// is given as a canonical URL, not the short prefix used
	// in the document being parsed.
	attr := xml.Attr{xml.Name{"", "AttrOne"}, "MyAttr"}
	fmt.Printf("%s: %s\n", attr.Name.Local, attr.Value)
}
