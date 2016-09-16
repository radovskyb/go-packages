package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Person struct {
	First string `xml:"Firstname"`
	Last  string `xml:"Lastname"`
}

func main() {
	// Marshal returns the XML encoding of v.
	//
	// Marshal handles an array or slice by marshalling each of the elements.
	// Marshal handles a pointer by marshalling the value it points at or, if the
	// pointer is nil, by writing nothing.  Marshal handles an interface value by
	// marshalling the value it contains or, if the interface value is nil, by
	// writing nothing.  Marshal handles all other data by writing one or more XML
	// elements containing the data.
	//
	// The name for the XML elements is taken from, in order of preference:
	//     - the tag on the XMLName field, if the data is a struct
	//     - the value of the XMLName field of type xml.Name
	//     - the tag of the struct field used to obtain the data
	//     - the name of the struct field used to obtain the data
	//     - the name of the marshalled type
	//
	// The XML element for a struct contains marshalled elements for each of the
	// exported fields of the struct, with these exceptions:
	//     - the XMLName field, described above, is omitted.
	//     - a field with tag "-" is omitted.
	//     - a field with tag "name,attr" becomes an attribute with
	//       the given name in the XML element.
	//     - a field with tag ",attr" becomes an attribute with the
	//       field name in the XML element.
	//     - a field with tag ",chardata" is written as character data,
	//       not as an XML element.
	//     - a field with tag ",innerxml" is written verbatim, not subject
	//       to the usual marshalling procedure.
	//     - a field with tag ",comment" is written as an XML comment, not
	//       subject to the usual marshalling procedure. It must not contain
	//       the "--" string within it.
	//     - a field with a tag including the "omitempty" option is omitted
	//       if the field value is empty. The empty values are false, 0, any
	//       nil pointer or interface value, and any array, slice, map, or
	//       string of length zero.
	//     - an anonymous struct field is handled as if the fields of its
	//       value were part of the outer struct.
	//
	// If a field uses a tag "a>b>c", then the element c will be nested inside
	// parent elements a and b.  Fields that appear next to each other that name
	// the same parent will be enclosed in one XML element.
	//
	// See MarshalIndent for an example.
	//
	// Marshal will return an error if asked to marshal a channel, function, or map.
	marshalled, err := xml.Marshal(Person{First: "Benjamin", Last: "Radovsky"})
	if err != nil {
		log.Fatalln(err)
	}

	// Print out the xml marshalled Person object
	fmt.Println(string(marshalled))

	// Create a Person object `unmarshalled` to store the unmarshalled xml from `marshalled`
	var unmarshalled Person

	// Unmarshal `marshalled` into the address of Person object `unmarshalled`
	//
	// Unmarshal parses the XML-encoded data and stores the result in
	// the value pointed to by v, which must be an arbitrary struct,
	// slice, or string. Well-formed data that does not fit into v is
	// discarded.
	//
	// Because Unmarshal uses the reflect package, it can only assign
	// to exported (upper case) fields.  Unmarshal uses a case-sensitive
	// comparison to match XML element names to tag values and struct
	// field names.
	if err := xml.Unmarshal(marshalled, &unmarshalled); err != nil {
		log.Fatalln(err)
	}

	// Print out Person `unmarshalled`
	fmt.Println(unmarshalled)

	// Print out the `First` and `Last` values stored in Person `unmarshalled`
	fmt.Println(unmarshalled.First)
	fmt.Println(unmarshalled.Last)
}
