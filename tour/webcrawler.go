package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	v   map[string]int
	mux sync.Mutex
}

func NewSafeMap() SafeMap {
	return SafeMap{map[string]int{}, sync.Mutex{}}
}

func (s *SafeMap) Inc(key string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.v[key]++
}

func (s *SafeMap) Value(key string) (v int, ok bool) {
	//	s.mux.Lock()
	//	defer s.mux.Unlock()
	v, ok = s.v[key]
	return
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeMap, quit chan int) {
	defer func(quit *chan int) {
		*quit <- 1
	}(&quit)
	if depth <= 0 {
		return
	}
	if _, ok := cache.Value(url); ok {
		cache.Inc(url)
		return
	} else {
		cache.Inc(url)
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	childQuit := make(chan int)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, cache, childQuit)
	}
	for _ = range urls {
		<-childQuit
	}
	return
}

func main() {
	crawlerCache := NewSafeMap()
	quit := make(chan int)
	go Crawl("http://golang.org/", 4, fetcher, &crawlerCache, quit)
	<-quit
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
