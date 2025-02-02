package main

import (
	"fmt"
	"sync"
)

/*
Task 1: Simulating a Ticket Booking System 💡
 Scenario: You are building a concert ticket booking system where multiple users try to book tickets at the same time.

Guidance:
Start with a total number of available tickets (e.g., 10).
Create multiple Goroutines representing users trying to book a ticket.
Use a channel to handle booking requests.
Process bookings safely by making sure two users don’t book the same ticket.
Print confirmation messages for successful bookings and rejections if tickets run out.
🔹 Bonus Challenge:

Use a sync.Mutex or sync/atomic package to ensure ticket count is updated safely.
Add a waiting list feature using a second channel for unfulfilled requests

*/

type Ticket struct {
	id       int
	boughtBy int
}

var tickets = make(chan *Ticket, 10)

var wg = &sync.WaitGroup{}

func main() {
	for i := 0; i < 10; i++ {
		ticket := &Ticket{id: i}
		tickets <- ticket
	}
	close(tickets)

	wg.Add(20)
	for i := 0; i < 20; i++ {
		go customer(i)
	}

	wg.Wait()
}

func customer(id int) {
	defer wg.Done()

	ticket, ok := <-tickets
	if !ok {
		fmt.Printf("customer %v could not get ticket - no tickets left'\n", id)
		return
	}

	fmt.Printf("customer %v is trying to book %v\n", id, ticket.id)

	if ticket.boughtBy == 0 {
		ticket.boughtBy = id
		fmt.Printf("cst %v, successfully bought ticket %v\n", id, ticket.id)
	}
}
