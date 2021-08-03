package main

import "fmt"

// Написать конвейер чисел. Даны 2 канала - в первый пишутся числа из массива, во второй пишется результат операции 2*x, после чего данные выводятся в stdout.

func main() {
	channelOne := make(chan int)
	channelTwo := make(chan int)
	go input(channelOne)
	go x2(channelOne, channelTwo)
	go stdoutX(channelTwo)
	fmt.Scanln()
}

func x2(cIn chan int, cOut chan int) {
	defer close(cOut)
	for x := range cIn {
		cOut <- x * 2
	}
}

func stdoutX(cIn chan int) {
	for x := range cIn {
		fmt.Println(x)
	}
}

func input(cOut chan int) {
	defer close(cOut)
	for x := 0; x < 10; x++ {
		cOut <- x
	}
}
