package main

import "fmt"

func conDup(ar []int) bool {

	n := len(ar)

	if len(ar) <= 1 {
		return false
	} else {
		for i := 0; i <= n-2; i++ {
			for j := i + 1; j <= n-1; j++ {
				if ar[i] == ar[j] {
					return true
				}
			}
		}

	}
	return false
}

func main() {
	ar := []int{1, 2, 3, 4, 5}
	ar2 := []int{1,1,1,3,3,4,3,2,4,2}

	res := conDup(ar)
	fmt.Printf("Array %v has duplicates %v\n", ar, res)

	res = conDup(ar2)
	fmt.Printf("Array %v has duplicates %v\n", ar2, res)
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
