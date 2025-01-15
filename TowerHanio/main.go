package main

import "fmt"

func towerOfHanoi(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2*towerOfHanoi(n-1) + 1
	}
}

func M(n int) int {
    return (1 << n) - 1 // Efficient calculation using bitwise shift
}


func main() {
	for i := range 10 {
		fmt.Printf("Steps %d\n", M(i))
	}
}
