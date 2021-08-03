package main

import "fmt"

// Что выведет программа данная программа?

func main() {
	n := 0
	if true {
		n := 1 // Another area of visibility. If we delete := and write = that we can result 2
		n++
	}
	fmt.Println(n)
}
