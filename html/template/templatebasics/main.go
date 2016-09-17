package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// Create the html template to be parsed and used as html
const tmpl = `<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>{{.Title}}</title
</head>
<body>
	{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>No Rows</strong></div>{{end}}
</body>
</html>`

func main() {
	// Create a new template with the given name `webpage` and parse it
	// passing it the text from const `tmpl`
	//
	// New allocates a new, undefined template with the given name.
	//
	// Parse defines the template by parsing the text. Nested template
	// definitions will be associated with the top-level template t.
	// Parse may be called multiple times to parse definitions of templates
	// to associate with t. It is an error if a resulting template is
	// non-empty (contains content other than template definitions) and
	// would replace a non-empty template with the same name.
	// (In multiple calls to Parse with the same receiver template,
	// only one call can contain text other than space, comments,
	// and template definitions.)
	t, err := template.New("webpage").Parse(tmpl)
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Title string
		Items []string
	}{
		Title: "Items Page",
		Items: []string{
			"My Photos",
			"My Blog",
		},
	}

	fmt.Println("Template With Items:")

	// Apply the template `t` to the data from `data` and write the
	// output to `os.Stdout`
	//
	// Execute applies a parsed template to the specified data object,
	// and writes the output to wr.
	// If an error occurs executing the template or writing its output,
	// execution stops, but partial results may already have been written to
	// the output writer.
	// A template may be executed safely in parallel.
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}

	noItemsData := struct {
		Title string
		Items []string
	}{
		Title: "No Items Page",
		Items: []string{},
	}

	fmt.Println("\n\nTemplate With No Items:")

	// Execute the template again but this time apply the data `noItemsData`
	// to the template, which will output the template slightly differently
	// to before, but once again print the data applied template to `os.Stdout`
	if err := t.Execute(os.Stdout, noItemsData); err != nil {
		log.Fatalln(err)
	}
}
