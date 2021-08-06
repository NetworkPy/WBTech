package main

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// Написать программу, которая переворачивает строку. Символы могут быть unicode.

func main() {
	br := testing.Benchmark(BenchmarkFunction)
	fmt.Println(br)
}

// Append is slow!
func reverseStrOne(str string) string {
	l := len([]rune(str))
	newStr := make([]rune, l)

	for i := l - 1; i >= 0; i-- {
		newStr = append(newStr, []rune(str)[i])
	}

	return string(newStr)
}

// Very fast!
func reverseStrTwo(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	fmt.Println(len(rune))
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	return output
}

// So so
func reverseStrThird(str string) string {
	l := len([]rune(str))
	newStr := make([]rune, l)

	for i := l - 1; i >= 0; i-- {
		newStr[i] = []rune(str)[i]
	}

	return string(newStr)
}

// Champion!!! Best result
func reverseStrFourth(str string) string {
	newStr := []rune(str)
	for i := 0; i < len(newStr)/2; i++ {
		newStr[i], newStr[len(newStr)-1-i] = newStr[len(newStr)-1-i], newStr[i]
	}
	return string(newStr)
}

// Benchmark
func BenchmarkFunction(b *testing.B) {
	str := "The quick brown 狐 jumped over the lazy 犬"
	// str := generateRandString(100)
	// reverseStrOne(str)
	// reverseStrThird(str)
	// reverseStrTwo(str)
	result := reverseStrFourth(str)
	fmt.Println(result)
}

// Generate random string
func generateRandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}
