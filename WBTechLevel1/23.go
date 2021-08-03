package main

import (
	"fmt"
)

// Написать бинарный поиск встроенными методами языка.
// FOR TEST IN STDIN
// 5 1 5 8 12 13 - 5 is number, array with values is [1 5 8 12 13]
// 5 8 1 23 1 11 - 5 is number, array with values for search in array is [8 1 23 1 11]
func main() {
	// const N = 10000
	// arr := make([]int, N)
	// for i := 0; i < N; i++ {
	// 	arr[i] = rand.Intn(100)
	// }
	// sort.Ints(arr)
	// fmt.Println(binarySearch(arr, 94))
	strArr := make([]int, 0)
	valArr := make([]int, 0)
	var strArrNum int
	var strValNum int
	var value int
	fmt.Scan(&strArrNum)
	for i := 0; i < strArrNum; i++ {
		fmt.Scan(&value)
		strArr = append(strArr, value)
	}

	fmt.Scan(&strValNum)
	for i := 0; i < strValNum; i++ {
		fmt.Scan(&value)
		valArr = append(valArr, value)
	}
	fmt.Println(strArr, valArr)
	for _, val := range valArr {
		// fmt.Println(val)
		binarySearch(strArr, val)
	}
}

// O(logN)
func binarySearch(a []int, k int) {
	l := 0          // Left point in array
	r := len(a) - 1 // Right point in array
	var m int       // Middle point in array

	for l <= r {
		m = l + ((r - l) / 2) // Calculate middle point
		if a[m] == k {
			fmt.Printf("%v ", m+1)
			return
		} else if a[m] > k {
			r = m - 1 // Change right point
		} else {
			l = m + 1 // Change left point
		}
	}
	fmt.Printf("%v ", -1)
	return
}
