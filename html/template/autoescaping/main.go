package main

import (
	"log"
	"os"
	"text/template"
)

// The template will be defined with the name "T" when
// using the text like `tmpl` contains with it's define keyword
const tmpl = `{{define "T"}}Hello, {{.}}!{{end}}`

func main() {
	// Create a new templte named `Foo` and parse it with the text from `tmpl`
	t, err := template.New("Foo").Parse(tmpl)
	if err != nil {
		log.Fatalln(err)
	}

	// Execute the template with t.ExecuteTemplate with also takes
	// the template's defined name
	//
	// ExecuteTemplate applies the template associated with t that has the given name
	// to the specified data object and writes the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel.
	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('Pwned!')</script>")
	if err != nil {
		log.Fatalln(err)
	}
}
