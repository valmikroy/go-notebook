package main

import (
	"fmt"
	"sync"
	"time"
)

// refer https://radhifadlillah.com/blog/2018-10-08-limiting-number-of-goroutine/

func scrapWebPage(url string) {
	time.Sleep(time.Second * 5)
	fmt.Println(url)
}

func main() {
	// urls is list of URL that will be downloaded
	urls := []string{
		"http://example-site.com/about",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
		"http://example-site.com/contact",
	}

	// waitGroup is used to make sure app doesn't finish prematurely
	// until all goroutines finished
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(urls))

	// guard is a channel that used to make sure that only N=10
	// goroutines will run at a time
	guard := make(chan struct{}, 3)

	for _, url := range urls {
		// as we loop through URLs, first we put an empty struct to channel guard.
		// If the channel is still empty, the process will continue to the next line.
		// Else, the process will be blocked until there are rooms in the channel to put the empty struct.
		guard <- struct{}{}

		go func(url string) {
			// when this goroutine finished, make sure to :
			// - mark the waitGroup for this goroutine as finished; and
			// - release the guard, so the next goroutine can be run.
			defer func() {
				waitGroup.Done()
				<-guard
			}()

			// download and process the URL
			scrapWebPage(url)
		}(url)
	}

	// wait until all goroutine finished
	waitGroup.Wait()
}
