package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	
	res := mapping(sum, arr)

	fmt.Println("Plus one>", res)

	res = mapping(double, arr)

	fmt.Println("Double>", res)
}

type funcType func(int) int

func mapping(name funcType, arr []int) []int {
	var res = make([]int, len(arr))
	for i, val := range arr {
		res[i] = name(val)
	}
	return res
}

func sum(n int) int {
	return n + 1
}

func double(n int) int {
	return n * 2
}
