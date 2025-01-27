package main

import (
	"fmt"
	"sync"
	"time"
)

// 5 forks, 5 philosophers, 2 besides each plate, each philosopher must eat with 2 forks

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of the philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3 // how many times does a person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var orderMutex sync.Mutex
var diningOrder = []string{}

func main() {
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("Table is empty.")

	time.Sleep(sleepTime)

	// start the meal
	dine()

	fmt.Println("The table is empty.")
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()

	fmt.Println(diningOrder)
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
	fmt.Printf("%s is seated at the table.\n", philosopher.name)
	seated.Done()

	seated.Wait()

	// eat number of hunger times
	for i := hunger; i > 0; i-- {

		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork.\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork.\n", philosopher.name)
		}

		fmt.Printf("%s has both forks and is eating.\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking.\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		fmt.Printf("\t%s put down the forks.\n", philosopher.name)
	}

	fmt.Println(philosopher.name, " is satisfied.")
	fmt.Println(philosopher.name, " left the table.")

	orderMutex.Lock()
	diningOrder = append(diningOrder, philosopher.name)
	orderMutex.Unlock()
}
