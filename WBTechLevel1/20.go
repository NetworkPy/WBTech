package main

import "fmt"

// Какой результат выполнения данного кода и почему?
// We have result [b b a] and next [a a]

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a") // Create new array in memory because we have cap = 2 but now we need 3 (Compiler do cap = 2 * 2 = 4)
		slice[0] = "b"             // Add to new memory
		slice[1] = "b"             // Add to new memory
		fmt.Print(slice)           // Print slice with new link to memory
	}(slice)
	fmt.Print(slice) // Print slice with old link to memory
}
