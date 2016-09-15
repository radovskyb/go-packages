package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"math"
)

type Pythagoras interface {
	Hypotenuse() float64
}

type Point struct {
	X, Y int
}

func (p Point) Hypotenuse() float64 {
	return math.Hypot(float64(p.X), float64(p.Y))
}

func myHypot(x, y int) {
	fmt.Println(x + (y / 2))
}

// This example shows how to encode an interface value. The key
// distinction from regular types is to register the concrete type that
// implements the interface.
func main() {
	network := new(bytes.Buffer) // Stand-in for the network.

	// We must register the concrete type for the encoder
	// and decoder (which would normally be on a separate machine
	// from the encoder). On each end, this tells the engine which
	// concrete type is being sent that implements the interface.
	//
	// Register records a type, identified by a value for that type,
	// under its internal type name. That name will identify the
	// concrete type of a value sent or received as an interface variable.
	// Only types that will be transferred as implementations of
	// interface values need to be registered. Expecting to be used
	// only during initialization, it panics if the mapping between
	// types and names is not a bijection.
	gob.Register(Point{})

	// Create a new encoder and send some values
	enc := gob.NewEncoder(network)
	for i := 1; i <= 3; i++ {
		interfaceEncode(enc, Point{3 * i, 4 * i})
		myHypot(3*i, 4*i)
	}

	// Create a new decoder and receive some values
	dec := gob.NewDecoder(network)
	for i := 1; i <= 3; i++ {
		result := interfaceDecode(dec)
		fmt.Printf("%d, %d: %v\n",
			result.(Point).X,
			result.(Point).Y,
			result.Hypotenuse(),
		)
	}
}

// interfaceEncode encodes the interface value into the encoder.
func interfaceEncode(enc *gob.Encoder, p Pythagoras) {
	// The encode will fail unless the concrete type has been
	// registered. We registered it in the calling function.

	// Pass pointer to interface so Encode sees (and hence sends)
	// a value of interface type. If we passed p directly it
	// would see the concrete type instead.
	// See the blog post, "The Laws of Reflection" for background.
	err := enc.Encode(&p)
	if err != nil {
		log.Fatalln("Encode error:", err)
	}
}

// interfaceDecode decodes the next interface value from
// the stream and returns it.
func interfaceDecode(dec *gob.Decoder) Pythagoras {
	var p Pythagoras
	// The decode will fail unless the concrete type on the wire has been
	// registered. We registered it in the calling function.
	err := dec.Decode(&p)
	if err != nil {
		log.Fatalln("Decode error:", err)
	}
	return p
}
