package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	for i := range 10 {
		fmt.Printf("Factorial of %d is %d\n", i, factorial(i))
	}
}