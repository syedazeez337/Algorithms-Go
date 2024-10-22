package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 8, 10}
	item := 7

	res := binarySearch(arr, item)

	fmt.Println(res)
}

func binarySearch(arr []int, item int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		if guess == item {
			return mid
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
