package main

import "fmt"

func main() {
	ar := []int{1, 2, 3, 4}
	ar2 := []int{5, 6, 7, 8, 9}

	str1 := []string{"A", "B", "C", "E"}
	str2 := []string{"F", "G", "H", "I"}

	fmt.Println(concat(ar, ar2))
	fmt.Println(concat(str1, str2))
}

func concat[T any] (ar []T, arr []T) []T {
	if len(ar) == 0 {
		return arr
	}

	if len(arr) == 0 {
		return ar
	}

	for _, elem := range arr {
		ar = append(ar, elem)
	}

	return ar
}