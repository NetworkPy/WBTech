package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая в конкурентном виде читает элементы из массива в stdout.

type ConArr struct {
	Rwm sync.RWMutex
	Wg  sync.WaitGroup
	Arr []int
}

func main() {
	var arr = []int{8, 2, 4, 5, 7, 8}
	cArr := NewConArr(arr)

	for i := range cArr.Arr {
		cArr.readFromArr(i)
	}
	cArr.Wg.Wait()

}

func (ca *ConArr) readFromArr(idx int) {
	ca.Wg.Add(1)
	defer ca.Wg.Done()
	if idx < len(ca.Arr) {
		ca.Rwm.RLock()
		defer ca.Rwm.RUnlock()
		fmt.Println(ca.Arr[idx])
	}
}

func NewConArr(arr []int) *ConArr {
	return &ConArr{Arr: arr}
}
