package main

import "fmt"

func uniqueElem(ar []int) bool {
	n := len(ar)

	for i := 0; i <= n-2; i++ {
		for j := i + 1; j <= n-1; j++ {
			if ar[i] == ar[j] {
				return false
			}
		}
	}

	return true
}

func main() {
	ar := []int{1, 2, 3, 4, 5}
	ar2 := []int{1, 2, 3, 1, 3, 4}

	res := uniqueElem(ar)
	fmt.Printf("Array %v is unique %v\n", ar, res)

	res = uniqueElem(ar2)
	fmt.Printf("Array %v is unique %v\n", ar2, res)
}