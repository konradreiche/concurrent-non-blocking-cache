package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

var urls = []string{
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"http://gopl.io",
	"https://golang.org",
	"https://godoc.org",
	"https://play.golang.org",
	"http://gopl.io",
}

func httpGetBody(url string) ([]byte, error) {
	//fmt.Printf("Fetching %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range urls {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

func TestSequential(t *testing.T) {
	cache := NewCache(httpGetBody)
	for url := range incomingURLs() {
		func(url string) {
			start := time.Now()
			value, err := cache.Get(url)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value))
		}(url)
	}
}

func TestConcurrent(t *testing.T) {
	cache := NewCache(httpGetBody)
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := cache.Get(url)
			if err != nil {
				t.Error(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value))
			n.Done()
		}(url)
	}
	n.Wait()
}
