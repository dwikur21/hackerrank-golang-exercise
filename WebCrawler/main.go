package main

import (
	"fmt"
	"sync"
)

// Fetcher an interface
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, c chan fetchResult)
}

type fetchedURL struct {
	urlFetched map[string]string
	mux        sync.Mutex
}

type fetchResult struct {
	body string
	urls []string
	err  error
}

func (f *fetchedURL) assignValue(key string, val string) {
	f.mux.Lock()
	f.urlFetched[key] = val
	f.mux.Unlock()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(f *fetchedURL, url string, depth int, fetcher Fetcher) {
	// DONE: Fetch URLs in parallel.
	// DONE: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	if depth <= 0 {
		return
	}

	if _, ok := f.urlFetched[url]; ok {
		return
	}
	c := make(chan fetchResult)
	go fetcher.Fetch(url, c)
	for i := range c {
		body, urls, err := i.body, i.urls, i.err
		if err != nil {
			fmt.Println(err)

			f.assignValue(url, err.Error())
			return
		}

		f.assignValue(url, body)

		fmt.Printf("found: %s %q\n", url, body)

		for _, u := range urls {
			Crawl(f, u, depth-1, fetcher)
		}
	}

	return
}

func main() {
	f := fetchedURL{urlFetched: make(map[string]string)}
	Crawl(&f, "https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, c chan fetchResult) {
	cRes := new(fetchResult)

	if res, ok := f[url]; ok {
		cRes.body = res.body
		cRes.urls = res.urls
		cRes.err = nil
		c <- *cRes
	} else {
		cRes.body = ""
		cRes.urls = nil
		cRes.err = fmt.Errorf("not found: %s", url)
		c <- *cRes
	}

	close(c)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
