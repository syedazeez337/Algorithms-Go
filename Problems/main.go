package main

import "fmt"

// Find even numbers from a list

// 1. array of numbers generator
func numGen(n int) []int {
	var res []int

	for i := range n + 1 {
		res = append(res, i)
	}

	return res
}

// filter function
func filter(fn func(int) bool, ar []int) []int {
	var res []int
	for _, elem := range ar {
		if fn(elem) {
			res = append(res, elem)
		}
	}
	return res
}

func evenNumber(n int) bool {
	if n == 0 {
		return false
	}
	return n % 2 == 0
}

func oddNumber(n int) bool {
	if n == 0 {
		return false
	}
	return !evenNumber(n)
}

func main() {
	ar := numGen(100)
	fmt.Println(filter(evenNumber, ar))
	fmt.Println(filter(oddNumber, ar))
}
