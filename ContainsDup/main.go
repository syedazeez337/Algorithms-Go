package main

import "fmt"

func conDup(ar []int) bool {
	var res bool

	if len(ar) <= 1 {
		res = false
	} else {
		for i, elem := range ar {
			if elem == ar[i+1] {
				res = true
				return res
			} else {
				res = false
				return res
			}
		}
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 6, 4}
	fmt.Println(conDup(arr))
}

/*
217. Contains Duplicate

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]

Output: true

Explanation:

The element 1 occurs at the indices 0 and 3.

Example 2:

Input: nums = [1,2,3,4]

Output: false

Explanation:

All elements are distinct.

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]

Output: true
*/
