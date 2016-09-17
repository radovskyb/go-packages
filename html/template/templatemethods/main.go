package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	// Parse all template files found in the current directory that contain
	// the text .html in them
	t, err := template.ParseGlob("*.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Print out a listing of the defined templates associated with `t`
	//
	// DefinedTemplates returns a string listing the defined templates,
	// prefixed by the string "; defined templates are: ". If there are none,
	// it returns the empty string. Used to generate an error message.
	fmt.Println(t.DefinedTemplates())

	// Lookup and return the index template from inside of `t`'s template list
	//
	// Lookup returns the template with the given name that is associated with t,
	// or nil if there is no such template.
	indexTemplate := t.Lookup("index")

	if indexTemplate != nil {
		// Print out indexTemplate's name
		//
		// Name returns the name of the template.
		fmt.Printf("\n%s\n\n", indexTemplate.Name())

		// Execute the template indexTemplate returned from calling t.Lookup
		if err := indexTemplate.Execute(os.Stdout, map[string]string{
			"Title":  "Index Page",
			"Header": "Hello, World; From the index page!",
		}); err != nil {
			log.Fatalln(err)
		}

		fmt.Println()
	}

	fmt.Println()

	// Create the data for the about page
	data := struct {
		Title  string
		Header string
	}{
		Title:  "About Page",
		Header: "Welcome to the about page!",
	}

	// Execute the about page template and pass it `data`
	if err := t.ExecuteTemplate(os.Stdout, "about", data); err != nil {
		log.Fatalln(err)
	}
}
