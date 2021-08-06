package main

import (
	"fmt"
	"math/rand"
)

// Даны 2 канала - в первый пишутся рандомные числа после чего они проверяются
// на четность и отправляются во второй канал. Результаты работы из второго канала пишутся в stdout.

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go generateRandomDigit(channel1)
	go checkEven(channel1, channel2)

	for val := range channel2 {
		fmt.Println(val)
	}
}

func checkEven(cIn <-chan int, cOut chan<- int) {
	defer close(cOut)
	for val := range cIn {
		if val%2 == 0 {
			cOut <- val
		}
	}
}

func generateRandomDigit(c chan<- int) {
	defer close(c)
	for i := 0; i < 100; i++ {
		c <- rand.Intn(10)
	}
}
