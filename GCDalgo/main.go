package main

import "fmt"

// Writing GCD Euclid's algorithm

func gcd(m, n int) int {
	if n == 0 {
		return m
	}
	if m > 0 && n > 0 {
		r := m % n
		m = n
		n = r
	}
	return gcd(m, n)
}

func main() {
	ans := gcd(60, 24)
	fmt.Println(ans)
}
