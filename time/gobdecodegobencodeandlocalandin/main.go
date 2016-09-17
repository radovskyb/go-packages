package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatalln(err)
	}
	t := time.Date(2015, 8, 6, 0, 0, 0, 0, loc)

	fmt.Println("Time:", t)
	fmt.Println("Time Location:", t.Location())

	// GobEncode implements the gob.GobEncoder interface.
	gobEncoded, err := t.GobEncode()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Gob encoded:", gobEncoded)

	fmt.Println()

	// with GOB encoded(marshalled) data, can save to file(serialize)

	// now, pretend that we loaded the GOB data
	// we want to decode(unmarshal) the data

	var gobTime time.Time

	// GobDecode implements the gob.GobDecoder interface.
	err = gobTime.GobDecode(gobEncoded)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Gob decoded time:", gobTime)
	fmt.Println("Gob decoded location (missing):", gobTime.Location())

	// because location is a pointer type (see http://golang.org/pkg/time/#Time.Location)
	// it won't be encoded by Gob
	// but you can translate it back to local time

	fmt.Println()

	// Local returns t with the location set to local time.
	fmt.Println("Gob decoded local:", gobTime.Local().String())

	// and convert back to target location
	//
	// In returns t with the location information set to loc.
	//
	// In panics if loc is nil.
	NYTime := gobTime.In(loc)

	fmt.Println("Gob decoded back to NY time:", NYTime)
}
