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
	//
	// ParseGlob creates a new Template and parses the template definitions from the
	// files identified by the pattern, which must match at least one file. The
	// returned template will have the (base) name and (parsed) contents of the
	// first file matched by the pattern. ParseGlob is equivalent to calling
	// ParseFiles with the list of files matched by the pattern.
	t, err := template.ParseGlob("*.html")
	if err != nil {
		log.Fatalln(err)
	}

	// Create the data for the index page
	data := struct {
		Title  string
		Header string
	}{
		Title:  "Index Page",
		Header: "Hello, World; From the index page!",
	}

	// Execute the index page template and pass it `data`
	err = t.ExecuteTemplate(os.Stdout, "index", data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println()

	// Create the data for the about page
	data = struct {
		Title  string
		Header string
	}{
		Title:  "About Page",
		Header: "Welcome to the about page!",
	}

	// Execute the about page template and pass it `data`
	err = t.ExecuteTemplate(os.Stdout, "about", data)
	if err != nil {
		log.Fatalln(err)
	}
}
