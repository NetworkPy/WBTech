package main

import "fmt"

// Дана переменная int64. Написать программу которая устанавливает i-й бит в 1 или 0.

func main() {
	fmt.Println(setBit(5, 2))
	fmt.Println(clearBit(5, 2))
}

func setBit(b int, i int) int {
	return (b | (1 << (i - 1)))
}

func clearBit(b int, i int) int {
	return b & (^(1 << (i - 1)))
}
