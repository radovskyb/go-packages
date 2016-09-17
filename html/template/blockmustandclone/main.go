package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

const (
	// Create a template text variable that uses the template keyword `block`
	// to easily reuse part of a template
	master  = `Numbers:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}{{"\n"}}`
	overlay = `{{define "list"}} {{join . ", "}}{{end}} `
)

var (
	// Create a new FuncMap that will use strings.Join when join is called from a template.
	funcs = template.FuncMap{"join": strings.Join}

	// Create some data to pass to the templates.
	numbers = []string{"one", "two", "three"}
)

func main() {
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatal(err)
	}

	// Clone returns a duplicate of the template, including all associated
	// templates. The actual representation is not copied, but the name space of
	// associated templates is, so further calls to Parse in the copy will add
	// templates to the copy but not to the original. Clone can be used to prepare
	// common templates and use them with variant definitions for other templates
	// by adding the variants after the clone is made.
	//
	// It returns an error if t has already been executed.
	//
	// Must is a helper that wraps a call to a function returning
	// (*Template, error) and panics if the error is non-nil. It
	// is intended for use in variable initializations such as
	// var t = template.Must(template.New("name").Parse("html"))
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
	}

	// Execute and write both templates to os.Stdout and parse `numbers`
	// to both as their data
	if err := masterTmpl.Execute(os.Stdout, numbers); err != nil {
		log.Fatalln(err)
	}

	if err := overlayTmpl.Execute(os.Stdout, numbers); err != nil {
		log.Fatalln(err)
	}
}
