package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	var arr1 = []int{-3, 1, 2, 3, 6, 8}
	var arr2 = []int{-3, 1, 2, 3, 6, 8}
	arr1 = deleteFromSliceFirst(arr1, 2)
	fmt.Println(arr1)
	arr2 = deleteFromSliceSecond(arr2, 0)
	fmt.Println(arr2)
}

// Changes the order of elements
func deleteFromSliceFirst(slice []int, num int) []int {
	slice[num] = slice[len(slice)-1]
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice
}

// Slow then first version. Not changes the order of elements
func deleteFromSliceSecond(slice []int, num int) []int {
	copy(slice[num:], slice[num+1:])
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice
}
