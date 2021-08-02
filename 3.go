package main

import (
	"fmt"
)

// Дана последовательность чисел  (2,4,6,8,10) найти их сумму квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	var arr = []int{2, 4, 6, 8, 10}
	firstSqr(arr)

}

func firstSqr(a []int) {
	channel := make(chan int, len(a)) // Create channel

	for _, value := range a {
		go func(v int, c chan<- int) { // Start worker
			vSqr := v * v
			c <- vSqr // Push in channel
		}(value, channel)
	}

	sum := 0
	for i := 0; i < len(a); i++ {
		sum += <-channel // Get from channel
	}
	fmt.Println(sum)
}
