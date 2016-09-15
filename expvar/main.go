package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"
	"time"
)

// KeyValue represents a single entry in a Map.
func kvfunc(kv expvar.KeyValue) {
	fmt.Println(kv.Key, kv.Value)
}

func main() {
	// Publish declares a named exported variable. This should be called from a
	// package's init function when it creates its Vars. If the name is already
	// registered then this will log.Panic.
	expvar.Publish("uptime", time.Now())

	var inerInt int64 = 10
	pubInt := expvar.NewInt("Int")
	pubInt.Set(inerInt) // pubInt: 10
	pubInt.Add(2)       // pubInt: 12

	var inerFloat float64 = 1.2
	pubFloat := expvar.NewFloat("Float")
	pubFloat.Set(inerFloat) // pubFloat: 1.2
	pubFloat.Add(0.1)       // pubFloat: 1.3

	var inerString string = "Hello, World!"
	pubString := expvar.NewString("String")
	pubString.Set(inerString)

	pubMap := expvar.NewMap("Map").Init()
	pubMap.Set("Int", pubInt)
	pubMap.Set("Float", pubFloat)
	pubMap.Set("String", pubString)

	pubMap.Add("Int", 1)
	pubMap.Add("NewInt", 123)
	pubMap.AddFloat("Float", 0.5)
	pubMap.AddFloat("NewFloat", 0.25)

	// Just for demonstration purposes
	fmt.Println("Int from pubMap:", pubMap.Get("Int"))

	// Get retrieves a named exported variable.
	fmt.Println("Int from expvar:", expvar.Get("Int"))

	// Do calls f for each entry in the map.
	// The map is locked during the iteration,
	// but existing entries may be concurrently updated.
	pubMap.Do(kvfunc)

	var x interface{} = `{"employees":[
	{"firstName":"John", "lastName":"Doe"},
	{"firstName":"Anna", "lastName":"Smith"},
	{"firstName":"Peter", "lastName":"Jones"}
	]}`

	// Func implements Var by calling the function
	// and formatting the returned value using JSON.
	f := expvar.Func(func() interface{} { return x })

	expvar.Publish("function f", f)

	// Do calls f for each exported variable.
	// The global variable map is locked during the iteration,
	// but existing entries may be concurrently updated.
	expvar.Do(kvfunc)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Gophers!")
	})

	log.Fatalln(http.ListenAndServe(":9000", nil))
}
