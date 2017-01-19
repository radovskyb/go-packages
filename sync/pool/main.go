package main

import (
	"fmt"
	"sync"
)

type Client struct {
	Name string
}

func (c Client) Talk() {
	fmt.Printf("Hi, my name is %s.\n", c.Name)
}

func NewClient(name string) *Client {
	return &Client{Name: name}
}

func main() {
	// Create a new sync.Pool called pool.
	//
	// A Pool is a set of temporary objects that may be individually saved and
	// retrieved.
	//
	// Any item stored in the Pool may be removed automatically at any time without
	// notification. If the Pool holds the only reference when this happens, the
	// item might be deallocated.
	//
	// A Pool is safe for use by multiple goroutines simultaneously.
	//
	// Pool's purpose is to cache allocated but unused items for later reuse,
	// relieving pressure on the garbage collector. That is, it makes it easy to
	// build efficient, thread-safe free lists. However, it is not suitable for all
	// free lists.
	//
	// An appropriate use of a Pool is to manage a group of temporary items
	// silently shared among and potentially reused by concurrent independent
	// clients of a package. Pool provides a way to amortize allocation overhead
	// across many clients.
	//
	// An example of good use of a Pool is in the fmt package, which maintains a
	// dynamically-sized store of temporary output buffers. The store scales under
	// load (when many goroutines are actively printing) and shrinks when
	// quiescent.
	//
	// On the other hand, a free list maintained as part of a short-lived object is
	// not a suitable use for a Pool, since the overhead does not amortize well in
	// that scenario. It is more efficient to have such objects implement their own
	// free list.
	var pool = &sync.Pool{
		// Set the pool's new function to return a value
		// from calling NewClient("Pool").
		New: func() interface{} {
			return NewClient("Pool")
		},
	}

	// Create a new client and call it's Talk method.
	c1 := NewClient("Client One")

	// Calling c1 talk says that it's name is Client One.
	c1.Talk()

	// Now that we are done with it, put the client c1 in to the pool.
	pool.Put(c1)

	// Now `later` another client is needed, so now instead
	// of creating a new one, use pool.Get().(*Client) to try
	// to take one from the pool and if none are already on there,
	// then create a new one using the pool's `new` method
	c2 := pool.Get().(*Client)

	// Call c2's Talk method, which will actually be calling
	// c1's Talk method since c2 is c1 taken from the pool.
	// c2 also says that it's name is Client One.
	//
	// However, since godoc states that `Callers should not assume
	// any relation between values passed to Put and the values
	// returned by Get.`, It's possible that this will actually
	// say that it's name is Pool (go run -race main.go).
	c2.Talk()

	// Now another client is needed, but c2 is still being used, so
	// now calling pool.Get().(*Client) will return a new client which
	// is created from using the pool.New() method.
	c3 := pool.Get().(*Client)

	// Call c3's Talk method which says that it's name is Pool.
	c3.Talk()
}
