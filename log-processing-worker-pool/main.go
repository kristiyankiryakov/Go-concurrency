package main

import (
	"fmt"
	"sync"
)

/*
Task 2: Log Processing with Worker Pool
💡 Scenario: You have a stream of log messages that need to be processed efficiently using multiple workers.

Hints (Less Guidance Now):
Read log lines from a slice (simulate a continuous log stream).
Create a channel to send log messages to workers.
Start N workers that process logs concurrently.
Each worker should print the log message with its worker ID.
Ensure all workers finish before exiting (sync.WaitGroup).
🔹 Bonus Challenge:
Instead of a fixed number of logs, use an infinite loop to continuously generate logs.
Use a quit channel to stop processing gracefully.
*/

var logsToProcess = []string{
	"INFO: app started successfully",
	"ERROR: Something went wrong... couldn't fetch users",
	"WARN: process runs slower than expected",
}

type Job struct {
	Message string
	ID      int
}

const numWorkers = 3

var wg = &sync.WaitGroup{}

func main() {

	messages := make(chan Job)

	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go worker(i, messages)
	}

	for workerID, message := range logsToProcess {
		loadMessageQueue(workerID, message, messages)
	}
	close(messages)

	wg.Wait()
}

func worker(id int, jobs <-chan Job) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processed: %s\n", id, job.Message)
	}
}

func loadMessageQueue(workerID int, message string, ch chan Job) {
	job := Job{ID: workerID, Message: message}
	ch <- job
}
