package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов значений взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}

	for i := range arr {
		wg.Add(1)
		go worker(arr[i], wg) // Запуск ворекра
	}
	wg.Wait()
}

func worker(digit int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(digit * digit)
	// runtime.Gosched()
}
