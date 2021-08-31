package main

import "fmt"

// Написать программу, которая проверяет, что все символы в строке уникальные.

func main() {
	str := "abcdaefg"
	fmt.Println(uniqueString(str))
}

func uniqueString(str string) bool {
	myMap := make(map[string]int)
	for _, s := range []byte(str) {
		if _, ok := myMap[string(s)]; ok {
			return false
		} else {
			myMap[string(s)] = 1
		}
	}
	return true
}