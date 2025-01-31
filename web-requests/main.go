package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Task 1: Concurrent Web Requests (Simulating API Calls)
// 💡 Scenario: You need to fetch data from multiple APIs concurrently and process the results.

// Guidance:
// Create a list of fake URLs (you won’t actually make HTTP calls, just simulate delays).
// Use Goroutines to "fetch" data from each URL (use time.Sleep to simulate network latency).
// Use a channel to collect results from each Goroutine.
// Print the results as they arrive.
// Ensure the main function waits for all Goroutines to finish using sync.WaitGroup.
// 🔹 Bonus Challenge: Introduce a timeout using select so if an API call takes too long, it gets skipped.

var urls = []string{
	"http://google.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://x.com",
	"http://linkedin.com",
}

func main() {
	wg := &sync.WaitGroup{}
	res := make(chan string)
	timeout := time.After(5 * time.Second)

	wg.Add(len(urls))

	for _, url := range urls {
		go fetch(url, res, wg)
	}

	for i := 0; i < len(urls); i++ {
		select {
		case result := <-res:
			fmt.Println(result)
		case <-timeout:
			fmt.Println("Timeout: Some requests took too long.")
			return
		}
	}

	close(res)

	wg.Wait()
}

func fetch(url string, res chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("something went wrong... %s", err)
		return
	}
	res <- resp.Status
}
