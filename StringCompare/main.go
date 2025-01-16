package main

import "fmt"

func strcmp(wrd1, wrd2 string) bool {
	if len(wrd1) != len(wrd2) {
		return false
	}

	n := len(wrd1)

	i := 0
	for i < n && wrd1[i] == wrd2[i] {
		i = i + 1
	}

	return i == n
}

func main() {
	wrd1 := "Hello world"
	wrd2 := "Hello friend"
	wrd3 := "Hello world"

	fmt.Printf("%v == %v -> %v\n", wrd1, wrd2, strcmp(wrd1, wrd2))
	fmt.Printf("%v == %v -> %v\n", wrd1, wrd3, strcmp(wrd1, wrd3))
}
