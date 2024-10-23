package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}

	var res = make([]int, len(arr))

	for i, val := range arr {
		res[i] = mapping(sum, val)
	}

	fmt.Println("Plus one>", res)

	for i, val := range arr {
		res[i] = mapping(double, val)
	}

	fmt.Println("Double>", res)
}

type funcType func(int) int

func mapping(name funcType, arr int) int {
	res := name(arr)
	return res
}

func sum(n int) int {
	return n + 1
}

func double(n int) int {
	return n * 2
}
