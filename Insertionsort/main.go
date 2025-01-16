package main

import (
	"fmt"
)

func insertionSort(ar []int) []int {
	// var insort []int

	n := len(ar)

	for i := 1; i < n; i++ {
		current := ar[i]

		j := i - 1
		for ; j >= 0 && ar[j] > current; j-- {
			ar[j+1] = ar[j]
		}
		ar[j+1] = current
		fmt.Printf("i = %v, sort: %v\n", i, ar)
	}

	return ar
}

func main() {
	ar := []int{61, 82, 67, 4, 98, 20, 37, 85}

	fmt.Printf("Sorted %v\n", insertionSort(ar))
}
