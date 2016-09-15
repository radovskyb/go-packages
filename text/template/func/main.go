package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

// Create a FuncMap with which to register the function to use with a template.
var funcMap = template.FuncMap{
	// The name "title" is what the function will be called in the template text
	"title": strings.Title,
}

// Create a simple template definition to test out the `title` function
// in the funcMap variable created above. We print the input text several ways:
// - the original
// - title-cased
// - title-cased and then printed with %q
// - printed with %q and then title-cased.
const templateText = `Input: {{printf "%q" .}}
Output 0: {{title .}}
Output 1: {{title . | printf "%q"}}
Output 2: {{printf "%q" . | title}}
`

func main() {
	// Create a new template, add the function map, then parse the template text.
	tmpl, err := template.New("titleFuncTest").Funcs(funcMap).Parse(templateText)
	if err != nil {
		log.Fatalln("New template error:", err)
	}

	// Run the template to verifty the output.
	err = tmpl.Execute(os.Stdout, "the go programming language.")
	if err != nil {
		log.Fatalln("tmpl.Execute error:", err)
	}
}
