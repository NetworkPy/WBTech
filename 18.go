package main

import "fmt"

// Что выведет данная программа и почему?

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}

// Compile create copy slice from main whoes have link to memory
// Next, we change zero element in copy and change this in memory
// Next, we add element to copy and have new link for it to new memory,
// but slice in main have old link.
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}
