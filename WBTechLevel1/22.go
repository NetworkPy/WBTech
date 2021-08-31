package main

import (
	"fmt"
	"sort"

	"github.com/google/go-cmp/cmp"
)

// Написать быструю сортировку встроенными методами языка.

func main() {
	arr1 := []int{5, 6, 8, 12, 1, 7, 6, 9, 101000, 12003}
	arr2 := []int{5, 6, 8, 12, 1, 7, 6, 9, 101000, 12003}
	quickSort(arr1, 0, len(arr1))
	sort.Ints(arr2)
	if cmp.Equal(arr1, arr2) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

}

func partition(a []int, l int, r int) int {
	x := a[l]
	j := l
	for i := l + 1; i < r; i++ {
		if a[i] <= x {
			j = j + 1
			a[j], a[i] = a[i], a[j]
		}
	}
	a[l], a[j] = a[j], a[l]
	return j
}

func quickSort(a []int, l int, r int) {
	if l >= r {
		return
	}
	m := partition(a, l, r)
	quickSort(a, l, m-1)
	quickSort(a, m+1, r)
}
