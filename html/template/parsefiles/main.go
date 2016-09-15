package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t := template.New("templates")

	var templates = template.Must(
		// Parse 2 template files use t.ParseFiles
		//
		// ParseFiles parses the named files and associates the resulting
		// templates with t. If an error occurs, parsing stops and the
		// returned template is nil; otherwise it is t. There must be at
		// least one file. Since the templates created by ParseFiles are
		// named by the base names of the argument files, t should usually
		// have the name of one of the (base) names of the files. If it
		// does not, depending on t's contents before calling ParseFiles,
		// t.Execute may fail. In that case use t.ExecuteTemplate to
		// execute a valid template.
		t.ParseFiles(
			"indexTemplate.html",

			// Parse the header and footer template files which will
			// be called from the index template
			"headerTmpl.html",
			"footerTmpl.html",

			"itemsTemplate.html",
		),
	)

	// Create the data to pass to the items template
	itemsTmplData := struct {
		Title string
		Items []string
	}{
		Title: "Items Template",
		Items: []string{
			"One",
			"Two",
			"Three",
			"Four",
			"Five",
		},
	}

	fmt.Println("itemsTemplate:")
	// Execute the itemsTemplate.html template from the list of templates
	// that were parsed using ParseFiles
	templates.ExecuteTemplate(os.Stdout, "itemsTemplate.html", itemsTmplData)

	// Create the data to pass to the index template
	indexTmplData := struct {
		Title string
	}{
		Title: "Index Page",
	}

	fmt.Println("\nindexTemplate:")
	// Now execute the indexTemplate.html template from the list of templates
	// that were parsed using ParseFiles
	templates.ExecuteTemplate(os.Stdout, "indexTemplate.html", indexTmplData)
}
