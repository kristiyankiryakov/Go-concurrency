package main

import (
	"fmt"
	"strings"
)

// channels - a way routines can communicate with each other
//  main functions is a go routine
// deadlock means - there is no routines listening

func shout(ping <-chan string, pong chan<- string) {
	for {
		s := <-ping

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)

	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("type something and press Enter (enter Q to quit)")

	for {
		fmt.Print("-> ")

		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for response
		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("All done, closing channels")
	close(ping)
	close(pong)
}
