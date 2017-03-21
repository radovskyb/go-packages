package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

const numbersTmplSrc = `Numbers:{{range .}}{{println}}- {{upper .}}{{end}}{{"\n"}}`

// Create a new FuncMap to pass to a template.
//
// FuncMap is the type of the map defining the mapping from names to
// functions. Each function must have either a single return value, or two
// return values of which the second has type error. In that case, if the
// second (error) argument evaluates to non-nil during execution, execution
// terminates and Execute returns that error. FuncMap has the same base type
// as FuncMap in "text/template", copied here so clients need not import
// "text/template".
var funcs = template.FuncMap{"upper": strings.ToUpper}

func main() {
	// Create a new template and parse it using template names but also add
	// a new template function that calls strings.ToUpper when calling `upper`
	// in the template's code as a new self defined function
	//
	// Funcs adds the elements of the argument map to the template's function map.
	// It panics if a value in the map is not a function with appropriate return
	// type. However, it is legal to overwrite elements of the map. The return
	// value is the template, so calls can be chained.
	namesTmpl, err := template.New("numbers").Funcs(funcs).Parse(numbersTmplSrc)
	if err != nil {
		log.Fatalln(err)
	}

	data := []string{"one", "two", "three"}

	// Execute the template and pass it `data` as it's data
	if err := namesTmpl.Execute(os.Stdout, data); err != nil {
		log.Fatalln(err)
	}
}
