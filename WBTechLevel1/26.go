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
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)
	return output
}

// Benchmark append vs swap
func BenchmarkFunction(b *testing.B) {
	// str := "The quick brown 狐 jumped over the lazy 犬"
	str := generateRandString(10000)
	// reverseStrOne(str)
	reverseStrTwo(str)
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
