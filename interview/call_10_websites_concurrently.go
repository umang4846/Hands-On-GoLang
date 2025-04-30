package main

/*
Problem statement : Call 10 different URLs concurrently
*/

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// fetchURL calls the given URL and prints the status code
func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this Goroutine as done
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Fetched %s - Status: %d - Time: %v\n", url, resp.StatusCode, time.Since(start))
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
		"https://www.reddit.com",
		"https://www.medium.com",
		"https://www.nytimes.com",
		"https://www.bbc.com",
		"https://www.cnn.com",
		"https://www.wikipedia.org",
		"https://www.apple.com",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait() // Wait for all Goroutines to complete
	fmt.Println("All URLs fetched")
}
