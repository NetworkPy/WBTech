package main

import (
	"sync"
	"sync/atomic"
)

// Написать свою структуру счетчик, которая будет инкрементировать и выводить значения в конкурентной среде.

type Counter struct {
	Wg    sync.WaitGroup
	Atom  atomic.Value
	Rwm   sync.RWMutex
	Count uint64
}

func main() {

}

func (c *Counter) Add() {

}

func (c *Counter) getCount() uint64 {
	return c.Count
}
