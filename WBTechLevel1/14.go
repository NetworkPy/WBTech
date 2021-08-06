package main

import "fmt"

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	strMap := make(map[string]int)
	var strArr = []string{"cat", "cat", "dog", "cat", "tree", "duck"}

	for _, str := range strArr {
		if _, ok := strMap[str]; !ok {
			strMap[str] = 1
		} else {
			continue
		}
	}

	result := make([]string, len(strMap))
	i := 0
	for key, _ := range strMap {
		result[i] = key
		i++
	}
	fmt.Println(result)
}
