package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

const (
	master  = `Numbers:{{block "list" .}}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}{{println}}`
	overlay = `{{define "list"}} {{join . ", "}}{{end}}`
)

var (
	// Create a new FuncMap that will use strings.Join when join is called from a template.
	funcs = template.FuncMap{"join": strings.Join}

	// Create some data to pass to the templates.
	numbers = []string{"one", "two", "three"}
)

func main() {
	// Create a new template from the master template text
	masterTmpl, err := template.New("master").Funcs(funcs).Parse(master)
	if err != nil {
		log.Fatalln("New master template error:", err)
	}

	// Create a new template by cloning the `masterTmpl` template
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatalln("New overlay template error:", err)
	}

	// Execute the `masterTmpl` template and pass it guardians as template data
	if err := masterTmpl.Execute(os.Stdout, numbers); err != nil {
		log.Fatalln("masterTmpl.Execute error:", err)
	}

	// Execute the `overlayTmpl` template and pass it guardians as template data
	if err := overlayTmpl.Execute(os.Stdout, numbers); err != nil {
		log.Fatalln("overlayTmpl.Execute error:", err)
	}
}
