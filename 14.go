package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	strMap := make(map[string]int)
	var strArr = []string{"cat", "cat", "dog", "cat", "tree", "duck"}
	result := make([]string, 0)
	for _, str := range strArr {
		if _, ok := strMap[str]; !ok {
			strMap[str] = 1
		} else {
			continue
		}
	}

	for key, _ := range strMap {
		result = append(result, key)
	}
	fmt.Println(result)
}
