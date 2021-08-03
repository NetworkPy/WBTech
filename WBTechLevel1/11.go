package main

import "fmt"

// Написать пересечение двух неупорядоченных массивов.

func main() {
	var arr1 = []int{3, 4, 7, 8, 10, 25275, 285, -1}
	var arr2 = []int{1, 3, 7, 9, 12, 18, 10, 285, 8, -1}
	res := SerachCross(arr1, arr2)
	fmt.Println(res)
}

func SerachCross(arr1 []int, arr2 []int) []int {
	result := make([]int, 0)
	hash := make(map[int]int)
	maxLen := 0

	if len(arr1) > len(arr2) {
		maxLen += len(arr1)
	} else {
		maxLen += len(arr2)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(arr1) {
			if _, ok := hash[arr1[i]]; !ok {
				hash[arr1[i]] = 1
			} else {
				hash[arr1[i]]++
			}
		}
		if i < len(arr2) {
			if _, ok := hash[arr2[i]]; !ok {
				hash[arr2[i]] = 1
			} else {
				hash[arr2[i]]++
			}
		}
	}

	for key, val := range hash {
		if val > 1 {
			result = append(result, key)
		}
	}
	return result
}
