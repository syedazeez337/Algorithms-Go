package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	num := 4
	val := search(arr, num)

	fmt.Println(arr, val)
}

func search(arr []int, num int) int {
	for i, elem := range arr {
		if elem == num {
			return i
		}
	}
	return -1
}
