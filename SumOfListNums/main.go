package main

import (
	"fmt"
)

func Sum(arr []int) int {
	total := 0
	for _, num := range arr {
		// fmt.Printf("Num -> %d\n", num)
		total += num
	}
	return total
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))              // 15
	fmt.Println(Sum([]int{3, 3}))                       // 6
	fmt.Println(Sum([]int{3, 5, 3, 5, 3}))              // 19
	fmt.Println(Sum([]int{4, 2, 22, -10, 8}))           // 26

	// let's just return 0 for empty lists
	fmt.Println(Sum(nil))     // 0
	fmt.Println(Sum([]int{})) // 0
}
