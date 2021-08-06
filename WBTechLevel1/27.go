package main

import (
	"fmt"
	"strings"
)

// Написать программу, которая переворачивает слова в строке (snow dog sun - sun dog snow).

func main() {
	str := "snow dog sun"
	fmt.Println(reverseWordsFirst(str))
}

func reverseWordsFirst(words string) string {
	wordsArr := strings.Split(words, " ")

	for i := 0; i < len(wordsArr)/2; i++ {
		wordsArr[i], wordsArr[len(wordsArr)-1-i] = wordsArr[len(wordsArr)-1-i], wordsArr[i]
	}
	return strings.Join(wordsArr, " ")
}
