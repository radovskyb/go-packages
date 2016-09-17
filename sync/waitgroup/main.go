package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	// Create a new WaitGroup `wg`.
	//
	// A WaitGroup waits for a collection of goroutines to finish.
	// The main goroutine calls Add to set the number of
	// goroutines to wait for. Then each of the goroutines
	// runs and calls Done when finished. At the same time,
	// Wait can be used to block until all goroutines have finished.
	var wg sync.WaitGroup

	var urls = []string{
		"http://www.golang.org",
		"http://www.google.com",
		"http://www.betsee.com.au",
	}

	for _, url := range urls {
		// Increment the WaitGroup counter.
		//
		// Add adds delta, which may be negative, to the WaitGroup counter.
		// If the counter becomes zero, all goroutines blocked on Wait are released.
		// If the counter goes negative, Add panics.
		//
		// Note that calls with a positive delta that occur when the counter is zero
		// must happen before a Wait. Calls with a negative delta, or calls with a
		// positive delta that start when the counter is greater than zero, may happen
		// at any time.
		// Typically this means the calls to Add should execute before the statement
		// creating the goroutine or other event to be waited for.
		// If a WaitGroup is reused to wait for several independent sets of events,
		// new Add calls must happen after all previous Wait calls have returned.
		// See the WaitGroup example.
		wg.Add(1)

		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Done decrements the WaitGroup counter.
			defer wg.Done()

			// Fetch the URL.
			_, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}
		}(url)
	}

	// Wait for all HTTP fetches to complete.
	//
	// Wait blocks until the WaitGroup counter is zero.
	wg.Wait()

	fmt.Println("Finished fetching all URLs.")
}
