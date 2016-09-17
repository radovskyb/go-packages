package main

import (
	"log"
	"os"
	"text/template"
)

// Define a template called letter
const letter = `Dear {{.Name}}
{{if .Attended}}
It was a pleasure to see you at the event.
{{- else}}
Sorry you couldn't be at the event.
{{- end}}
{{with .Gift -}}Thank you for the lovely {{.}}.{{end}}

Best wishes,
Benjamin
{{println}}`

// Create a structure to hold the data that will
// be used with a template
type Recipient struct {
	Name, Gift string
	Attended   bool
}

func main() {
	// Prepare some data to insert into the template below
	recipients := []Recipient{
		{"Benji Radovsky", "tea set", true},
		{"Rado Soft", "rubik's cube", false},
		{"Blabehdy Dah", "lala leelah", true},
	}

	// Create a new template and parse the letter into it.
	tmpl := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient
	for _, r := range recipients {
		err := tmpl.Execute(os.Stdout, r)
		if err != nil {
			log.Fatalln("tmpl.Execute error:", err)
		}
	}
}
