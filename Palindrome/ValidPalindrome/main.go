package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "race a car"

	s = strings.ToLower(s)

	fmt.Println(s)

	result := same(s)
	fmt.Println(result)

	palin := palindrome(result)
	fmt.Println(palin)
}

func same(s string) string {
	var res []rune
	for _, val := range s {
		// fmt.Println(i, val)
		if val >= 97 && val <= 122 {
			res = append(res, val)
		}
	}
	return string(res)
}

func palindrome(s string) bool {
	return false
}