package main

import (
	"fmt"
	"sync"
)

func printSth(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"epsilon",
		"delta",
		"zeta",
	}

	wg.Add(len(words))

	for i, v := range words {
		go printSth(fmt.Sprintf("%d: %s", i, v), &wg)
	}

	wg.Wait()

	wg.Add(1)

	printSth("Second one", &wg)
}
