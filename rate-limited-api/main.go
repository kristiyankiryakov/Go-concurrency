package main

import (
	"context"
	"fmt"
	"time"
)

/*
 Task 3: Rate Limiting an API Using Goroutines and Channels 💡
 Scenario: You are building an API that allows only a certain number of requests per second to prevent overload.

Minimal Guidance Now – Just Hints!
Implement a token bucket algorithm using a buffered channel.
The channel holds tokens (representing available requests).
A Goroutine refills the channel at a fixed rate (e.g., 5 tokens per second).
Clients must take a token from the channel before making a request.
If no token is available, the request must wait or be rejected.
🔹 Bonus Challenge:

Allow burst capacity (e.g., up to 10 tokens can be used instantly, but only refilled at 5 per second).
Simulate multiple clients making requests at random intervals.
*/

type tokens chan struct{}

type TokenBucket struct {
	count  int64
	tokens tokens
	ticker *time.Ticker
}

func main() {

	// create a new TokenBucket with a bucket size of 2 and rate of 5 token per second
	bucket := NewTokenBucket(2, 5)

	// Start the token bucket background process.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bucket.Start(ctx)

	//Simulate request being made at a constant rate
	for i := 0; i < 10; i++ {
		requestedAt := time.Now()
		bucket.Wait()
		processedAt := time.Now()
		fmt.Printf("Processing request [%v] at [%v] processed at [%v]\n", i+1, requestedAt, processedAt)
	}
}

func NewTokenBucket(count int64, rate int64) *TokenBucket {

	tokens := make(chan struct{}, count)

	// we start with full bucket
	c := int(count)
	for i := 0; i < c; i++ {
		tokens <- struct{}{}
	}

	fmt.Printf("Initialized Tokenbucket with count %v and token rateMS %v\n", len(tokens), rate)

	everyMs := 1 / float64(rate) * 1000
	return &TokenBucket{
		count:  count,
		tokens: tokens,
		ticker: time.NewTicker(time.Duration(int64(everyMs) * int64(time.Millisecond))),
	}

}

func (tb *TokenBucket) Start(ctx context.Context) {
	// spinning up a goroutine for the ticker to fill the bucket if any slot is available

	go func() {
		for {
			select {
			case <-tb.ticker.C:
				select {
				case tb.tokens <- struct{}{}:
					fmt.Printf("added token, current token count %v\n", len(tb.tokens))
				default:
				}
			case <-ctx.Done():
				fmt.Println("context cancelled")
				return
			}
		}
	}()
}

func (tb *TokenBucket) Wait() {
	// reading the tokens from the bucket, it blocks if empty until a new token is available
	<-tb.tokens
	fmt.Printf("consumed token, current token count %v\n", len(tb.tokens))
}
