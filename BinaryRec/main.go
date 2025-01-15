package main

import "fmt"

/*
func binRec(n int) int {
	if n == 1 {
		return 0
	} else {
		return binRec(n/2) + 1
	}
}
*/

func binRec(n int) int {
	if n <= 0 {
		return -1 // Handle invalid input (e.g., negative or zero)
	}
	count := 0
	for n > 1 {
		n /= 2
		count++
	}
	return count
}

func main() {
	for i := range 10 {
		fmt.Printf("Binary %v\n", binRec(i))
	}
}
