package main

import "fmt"

// Создать слайс с предварительно выделенными 100 элементами.

func main() {
	slice := make([]int, 0, 100)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = append(slice, 1)
	fmt.Println(len(slice))
}
