package main

import "fmt"

func towerOfHanoi(n int) int {
	if n == 1 {
		return 1
	} else {
		return 2*towerOfHanoi(n-1) + 1
	}
}

func main() {
	for i := range 3 {
		fmt.Printf("Steps %d\n", towerOfHanoi(i))
	}
}
