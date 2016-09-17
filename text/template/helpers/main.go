package main

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
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
	// Load the helpers.
	tmpls := template.Must(template.ParseGlob(pattern))

	// Add one driver template to the bunch; we do this with an
	// explicit template definition.
	_, err := tmpls.Parse(
		"{{define `driver1`}}Driver 1 calls T1: ({{template `T1`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver1 template error: ", err)
	}

	// Add another driver template.
	_, err = tmpls.Parse(
		"{{define `driver2`}}Driver 2 calls T2: ({{template `T2`}})\n{{end}}")
	if err != nil {
		log.Fatal("parsing driver2 template error: ", err)
	}

	// We load all the templates before execution. This package does not require
	// that behavior but html/template's escaping does, so it's a good habit.
	err = tmpls.ExecuteTemplate(os.Stdout, "driver1", nil)
	if err != nil {
		log.Fatalf("driver1 execution: %s", err)
	}

	err = tmpls.ExecuteTemplate(os.Stdout, "driver2", nil)
	if err != nil {
		log.Fatalf("driver2 execution: %s", err)
	}
}
