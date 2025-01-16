package main

import (
	"fmt"
	"time"
)

func main() {
	arr := sortedArray(50)
	num := 105
	start := time.Now()
	val := linearSearch(arr, num)

	fmt.Printf("Array: %v, Index: %v, time taken: %v\n", arr, val, time.Since(start))

	start = time.Now()
	val = binarySearch(arr, num)
	fmt.Printf("Array: %v, Index: %v, time taken: %v\n", arr, val, time.Since(start))
}

func linearSearch(arr []int, num int) int {
	for i, elem := range arr {
		if elem == num {
			return i
		}
	}
	return -1
}

// binary search
func binarySearch(ar []int, n int) int {
	indexHigh := len(ar) - 1
	indexLow := 0

	for indexLow <= indexHigh {
		indexMid := int((indexHigh + indexLow) / 2)

		if ar[indexMid] == n {
			return indexMid
		}

		if ar[indexMid] < n {
			indexLow = indexMid + 1
		} else {
			indexHigh = indexMid - 1
		}
	}
	return -1
}

// generate a sorted array
func sortedArray(n int) []int {
	var arr []int

	for i := 0; i <= n; i++ {
		arr = append(arr, i)
	}

	return arr
}
