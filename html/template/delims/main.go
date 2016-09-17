package main

import (
	"html/template"
	"log"
	"os"
)

const delimsTmpl = `<# range . #><# . #><# println #><# end #>`

func main() {
	// Create a new template and before parsing it, set it's template's
	// code delimiters to `<#` and `#>` instead of `{{` and `}}`
	//
	// Delims sets the action delimiters to the specified strings, to be used in
	// subsequent calls to Parse, ParseFiles, or ParseGlob. Nested template
	// definitions will inherit the settings. An empty delimiter stands for the
	// corresponding default: {{ or }}.
	// The return value is the template, so calls can be chained.
	t := template.Must(template.New("delimsTmpl").Delims("<#", "#>").
		Parse(delimsTmpl))

	// Create a names slice to pass as data to template `t`
	names := []string{"Benjamin", "Radovsky", "Benji", "Radosoft"}

	// Execute template `t` to print to os.Stdout and pass
	// `names` as it's data to it
	if err := t.Execute(os.Stdout, names); err != nil {
		log.Fatalln(err)
	}
}
