package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
)

// NewFileSet creates a new file set.
//
// A FileSet represents a set of source files.
// Methods of file sets are synchronized; multiple goroutines
// may invoke them concurrently.
var fset = token.NewFileSet()

// This is some sample documentation that will be parsed below.
func main() {
	// Parse this Go source file into an *ast.File.
	//
	// ParseFile parses the source code of a single Go source file and returns
	// the corresponding ast.File node. The source code may be provided via
	// the filename of the source file, or via the src parameter.
	//
	// If src != nil, ParseFile parses the source from src and the filename is
	// only used when recording position information. The type of the argument
	// for the src parameter must be string, []byte, or io.Reader.
	// If src == nil, ParseFile parses the file specified by filename.
	//
	// The mode parameter controls the amount of source text parsed and other
	// optional parser functionality. Position information is recorded in the
	// file set fset, which must not be nil.
	//
	// If the source couldn't be read, the returned AST is nil and the error
	// indicates the specific failure. If the source was read but syntax
	// errors were found, the result is a partial AST (with ast.Bad* nodes
	// representing the fragments of erroneous source code). Multiple errors
	// are returned via a scanner.ErrorList which is sorted by file position.
	file, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}

	// Print out all of this file's imports.
	//
	// Imports are imports in a file.
	fmt.Println("Imports:")
	for _, i := range file.Imports {
		// Path is a an import path.
		fmt.Println(strings.Trim(i.Path.Value, "\""))
	}

	fmt.Println()

	// Inspect every single ast.Node in this file.
	//
	// Inspect traverses an AST in depth-first order: It starts by calling
	// f(node); node must not be nil. If f returns true, Inspect invokes f
	// recursively for each of the non-nil children of node, followed by a
	// call of f(nil).
	ast.Inspect(file, func(n ast.Node) bool {
		// Check if the current node is a function declaration.
		//
		// A FuncDecl node represents a function declaration.
		if fnc, ok := n.(*ast.FuncDecl); ok {
			// Print out the function's name if it's not an exported function.
			//
			// IsExported reports whether id is an exported Go symbol
			// (that is, whether it begins with an uppercase letter).
			if !fnc.Name.IsExported() {
				fmt.Printf("Name:\n%s\n\n", fnc.Name.Name)
			}

			// Print out the function's documentation.
			//
			// Text returns the text of the comment.
			// Comment markers (//, /*, and */), the first space of a line comment, and
			// leading and trailing empty lines are removed. Multiple empty lines are
			// reduced to one, and trailing space on lines is trimmed. Unless the result
			// is empty, it is newline-terminated.
			fmt.Printf("Documentation:\n%s\n", fnc.Doc.Text())

			fmt.Println("Function:")
			// Print out the entire function.
			//
			// Fprint "pretty-prints" an AST node to output.
			// It calls Config.Fprint with default settings.
			printer.Fprint(os.Stdout, fset, fnc)

			// Now set the function's body and doc to nil so only the
			// function's signature will be printed in the next
			// call to printer.Fprint.
			fnc.Body = nil
			fnc.Doc = nil

			// Now print out only the function's signature without it's documentation.
			fmt.Println("\n\nSignature:")
			printer.Fprint(os.Stdout, fset, fnc)
		}
		return true
	})
}
