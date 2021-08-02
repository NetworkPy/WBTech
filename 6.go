package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Какие существуют способы остановить выполнения горутины? Написать примеры использования.

func main() {
	fmt.Println("Start")
	// firstExample()
	// secondExample()
	// thirdExample()
	// fourthExample()
	fifthExample()
}

// Example with channel
func firstExample() {
	channel := make(chan struct{})
	go func(c chan struct{}) {
	LOOP:
		for {
			select {
			case <-c:
				fmt.Println("Stop")
				break LOOP
			default:
				fmt.Println("Hi")
			}
		}
	}(channel)
	time.Sleep(10 * time.Millisecond)
	channel <- struct{}{}
}

// Close channel
func secondExample() {
	channel := make(chan struct{})
	go func(c chan struct{}) {
		for range c {
			fmt.Println("Hi")
		}
	}(channel)

	for i := 0; i < 100; i++ {
		if i == 15 {
			close(channel)
			break
		}
		channel <- struct{}{}
	}
}

// Context. We can use timeout, cancel or deadline
// timeout use duration
// cancel use cancel function
// deadline use time

// Cancel
func thirdExample() {
	gen := func(ctx context.Context) <-chan int {
		channel := make(chan int)
		num := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case channel <- num:
					num++
				}
			}
		}()
		return channel
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for value := range gen(ctx) {
		fmt.Println(value)
		if value == 5 {
			break
		}
	}
}

// Deadline
func fourthExample() {
	const shortDuration = 3 * time.Second
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel() // Do it always!!!!

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, wg *sync.WaitGroup, ctx context.Context) {
			defer wg.Done()
		LOOP:
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Stop after %v\n", shortDuration)
					break LOOP
				default:
					fmt.Printf("Worker number %v say %s\n", num, "Hi")
					runtime.Gosched()
				}
			}
		}(i, wg, ctx)
	}
	wg.Wait()
}

// Timeout
func fifthExample() {
	const shortDuration = 1 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	wg := &sync.WaitGroup{}

	defer cancel()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int, ctx context.Context) {
			defer wg.Done()
			n := rand.Intn(5)
			timer := time.NewTimer(time.Duration(n) * time.Second)
		LOOP:
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Not finish %v with time work %v\n", i, n)
					break LOOP
				case <-timer.C:
					timer.Stop()
					fmt.Printf("Finish %v\n", i)
					break LOOP
				}
			}
		}(wg, i, ctx)
	}
	wg.Wait()
}
