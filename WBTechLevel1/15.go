package main

import "fmt"

// Поменять местами два числа без создания временной переменной.

func main() {
	var (
		a = 5
		b = 7
	)
	a, b = fifthVersion(a, b)
	fmt.Println(a, b)
}

// With + and -
func firstVersion(a int, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

// With - and +
func secondVersion(a int, b int) (int, int) {
	a = a - b
	b = a + b
	a = b - a
	return a, b
}

// =))
func thirdVersion(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

// With * and /. If b = 0 we have error
func fourthVersion(a int, b int) (int, int) {
	a = a * b
	b = a / b
	a = a / b
	return a, b
}

// A ^ B = C
// C ^ B = A
// A ^ C = B 
func fifthVersion(a int, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
