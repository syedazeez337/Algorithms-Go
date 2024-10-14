package main

import (
	"fmt"
)

func NumInList(arr []int, n int) bool {
	for i:=0; i < len(arr); i++ {
		if arr[i] == n {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(NumInList([]int{1, 2, 3, 4, 5}, 5))      // true
	fmt.Println(NumInList([]int{3, 3, 3, 3, 3}, 5))      // false
	fmt.Println(NumInList([]int{3, 5, 3, 5, 3}, 5))      // true
	fmt.Println(NumInList([]int{4, 2, 22, -10, 8}, -10)) // true

	// empty lists!
	fmt.Println(NumInList(nil, 5))     // false
	fmt.Println(NumInList([]int{}, 5)) // false
}
