package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Написать свою структуру счетчик, которая будет инкрементировать и выводить значения в конкурентной среде.

type Counter struct {
	rw    sync.RWMutex
	Count int64
}

func main() {
	wg := &sync.WaitGroup{}
	counter := NewCounter()
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go doJob(counter, wg)
	}
	wg.Wait()
	fmt.Println(counter.Count)
}

func (c *Counter) Add(value int64) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.Count += value
	if c.Count > 500 {
		c.Zero()
	}
	// atomic.AddInt64(&c.Count, value)
}

func (c *Counter) Zero() {
	c.Count = int64(0)
}

func (c *Counter) getCount() int64 {
	return c.Count
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) PrintCount() {
	c.rw.RLock()
	defer c.rw.RUnlock()
	fmt.Println(c.Count)
}

func doJob(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	sleep := rand.Intn(10)
	time.Sleep(time.Duration(sleep) * time.Second) // Do job
	// Next we add value to counter
	c.Add(int64(sleep)) //  Add sleep time
	c.PrintCount()      // Print count
}
