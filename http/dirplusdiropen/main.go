package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// A Dir implements FileSystem using the native file system restricted to a
	// specific directory tree.
	//
	// While the FileSystem.Open method takes '/'-separated paths, a Dir's string
	// value is a filename on the native file system, not a URL, so it is separated
	// by filepath.Separator, which isn't necessarily '/'.
	//
	// An empty Dir is treated as ".".
	dir := http.Dir("./")

	file, err := dir.Open("file.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the file info from the file
	fi, err := file.Stat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print the file's name
	fmt.Println(fi.Name())

	// Get the file's contents
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Print them to the user as a string
	fmt.Fprintln(w, string(contents))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
