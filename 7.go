package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись в map.

type MyMap struct {
	rwm sync.RWMutex
	Map map[int]string
}

// Write to map
func (mm *MyMap) WriteMap(key int, value string) {
	mm.rwm.Lock()
	defer mm.rwm.Unlock()
	mm.Map[key] = value
}

// Read from map
func (mm *MyMap) ReadMap(key int) (string, bool) {
	mm.rwm.RLock()
	defer mm.rwm.RUnlock()
	value, ok := mm.Map[key]
	return value, ok
}

// Create new struct for work with map
func newMyMap() *MyMap {
	Map := make(map[int]string)
	return &MyMap{
		Map: Map,
	}
}

func main() {
	m := newMyMap()
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(num int, wg *sync.WaitGroup, m *MyMap) {
			defer wg.Done()
			m.WriteMap(num, "Hello")
		}(i, wg, m)
	}

	wg.Wait()

	fmt.Println(m.Map)
}
