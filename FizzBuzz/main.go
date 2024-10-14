package main

import (
	"fmt"
)

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%v ", "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Printf("%v ", "Fizz")
		} else if i%5 == 0 {
			fmt.Printf("%v ", "Buzz")
		} else {
			fmt.Printf("%v ", i)
		}
	}
	fmt.Println()
}

func main() {
	FizzBuzz(30)
	FizzBuzz(1)
	// 1
	FizzBuzz(5)
	// 1, 2, Fizz, 4, Buzz
	FizzBuzz(15)
	// 1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz
}
