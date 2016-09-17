package main

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

type templateFile struct {
	Name, Text string
}

func createTestDir(templates []templateFile) string {
	dir := os.TempDir()
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalln("Chdir error:", err)
	}
	for _, tmpl := range templates {
		f, err := os.Create(tmpl.Name)
		if err != nil {
			log.Fatalf("Creating template file %s error: %v\n", tmpl.Name, err)
		}
		_, err = f.WriteString(tmpl.Text)
		if err != nil {
			log.Fatalln(err)
		}
		if err := f.Close(); err != nil {
			log.Fatalln(err)
		}
	}
	return dir
}

func main() {
	// Create a temporary directory and populate it with our sample
	// template definition files; usually the template files would already
	// exist in some location known to the program.
	dir := createTestDir([]templateFile{
		// T0.tmpl is a plain template file that just invokes T1.
		{"T0.tmpl", `T0 invokes T1: ({{template "T1"}}){{println}}`},
		// T1.tmpl defines a template, T1 that invokes T2.
		{"T1.tmpl", `{{define "T1"}}T1 invokes T2: ({{template "T2"}}){{end}}`},
		// T2.tmpl defines a template T2.
		{"T2.tmpl", `{{define "T2"}}This is T2{{end}}`},
	})

	// Clean up after the test; another quirk of running as an example.
	defer os.RemoveAll(dir)

	// pattern is the glob pattern used to find all the template files.
	pattern := filepath.Join(dir, "*.tmpl")

	// Here starts the example proper.
	// T0.tmpl is the first name matched, so it becomes the starting template,
	// the value returned by ParseGlob.
	tmpl := template.Must(template.ParseGlob(pattern))

	// Execute the template `tmpl`
	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("tmpl.Execute error:", err)
	}
}
