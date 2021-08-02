package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Написать программу, которая будет последовательно писать значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершиться.

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fiveQuestFirst(wg)
	wg.Wait()
}

func fiveQuestFirst(wg *sync.WaitGroup) {
	defer wg.Done()
	workTime := 5 * time.Second                                   // Time work our porogramm
	ctx, _ := context.WithTimeout(context.Background(), workTime) // Create context
	channel := make(chan string)                                  // Create channel for push and get
	defer close(channel)                                          // Optional

	go func(c chan<- string, ctx context.Context) {
		ticker := time.NewTicker(2 * time.Second) // Create ticker for push in channel in terms of time
	LOOP:
		for {
			select {
			case <-ctx.Done():
				ticker.Stop() // We need to stop ticker else it call cause leak
				break LOOP    // If ctx is Done next break goroutine
			case <-ticker.C: // Do while ctx not Done
				c <- "Hello"
			}
		}
	}(channel, ctx)

	go func(c <-chan string, ctx context.Context) {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case val := <-c:
				fmt.Println(val)
			}
		}
	}(channel, ctx)
	<-ctx.Done()
}
