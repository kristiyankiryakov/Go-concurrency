package main

// 1.Create a Goroutine that prints "Hello from Goroutine!" and ensure the main function waits for it.
// 2.Goal: Create multiple Goroutines and observe concurrency.
// 3.Goal: Use a channel to synchronize Goroutines with main().
// 4.Goal: Send multiple values through a channel and receive them.

// Goal: Chain multiple Goroutines where each stage modifies the data.
// Create three stages:
// Generate numbers
// Double the numbers
// Print the results
// Use separate channels to pass data between stages.

// func main() {

// 	wg := &sync.WaitGroup{}
// 	numbers := make(chan int)
// 	doubled := make(chan int)

// 	wg.Add(2)

// 	go func() {
// 		defer wg.Done()
// 		loadNumbers(numbers)
// 		close(numbers)
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for val := range numbers {
// 			doubled <- val * 2
// 		}
// 		close(doubled)
// 	}()

// 	for val := range doubled {
// 		fmt.Println(val)
// 	}

// 	wg.Wait()

// }

// func loadNumbers(numbers chan<- int) {
// 	for i := 1; i <= 10; i++ {
// 		numbers <- i
// 	}
// }

// 7.Goal: Merge multiple channels into one.

// func main() {

// 	even := make(chan int)
// 	odd := make(chan int)
// 	merged := make(chan int)
// 	wg := &sync.WaitGroup{}

// 	go func() {
// 		for i := 1; i <= 10; i++ {
// 			loadChannels(even, odd, i)
// 		}
// 		close(even)
// 		close(odd)
// 	}()

// 	go func() {
// 		for {
// 			select {
// 			case v, ok := <-even:
// 				if !ok {
// 					return
// 				}
// 				merged <- v
// 			case v, ok := <-odd:
// 				if !ok {
// 					return
// 				}
// 				merged <- v
// 			}
// 		}
// 	}()

// 	wg.Add(10)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-merged)
// 			wg.Done()
// 		}
// 		close(merged)
// 	}()

// 	wg.Wait()

// }

// func loadChannels(even chan int, odd chan int, num int) {
// 	if num%2 == 0 {
// 		even <- num
// 	} else {
// 		odd <- num
// 	}
// }

// // 6.Goal: Distribute work across multiple Goroutines using a single input channel.
// func main() {

// 	ch := make(chan int, 10)
// 	wg := &sync.WaitGroup{}

// 	for i := 1; i <= 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)

// 	wg.Add(5)
// 	for i := 1; i <= 5; i++ {
// 		go consume(wg, ch)
// 	}

// 	wg.Wait()
// }

// func consume(wg *sync.WaitGroup, ch <-chan int) {
// 	defer wg.Done()

// 	first, ok := <-ch
// 	if !ok {
// 		return
// 	}

// 	second, ok := <-ch
// 	if !ok {
// 		return
// 	}
// 	fmt.Println(first + second)
// }

// 5.Goal: Use a buffered channel to hold messages.

// func main() {

// 	ch := make(chan string, 3)

// 	ch <- "Hello"
// 	ch <- "from"
// 	ch <- "Go"

// 	res := []string{}

// 	for i := 1; i <= 3; i++ {
// 		collect(ch, &res)
// 	}

// 	fmt.Println(res)
// 	close(ch)
// }

// func collect(ch chan string, res *[]string) {
// 	*res = append(*res, <-ch)
// }
