package main

import (
	"fmt"
	"math"
)

func sieve(n int) []int {
	var res []int
	var out []int

	if n > 1 {
		for i := 0; i <= n; i++ {
			res = append(res, i)
		}

		for p := 2; p <= int(math.Sqrt(float64(n))); p++ {
			if res[p] != 0 {
				for j := p * p; j <= n; j = j + p {
					res[j] = 0
				}
			}
		}

		for p := 2; p <= n; p++ {
			if res[p] != 0 {
				out = append(out, res[p])
			}
		}
	}

	return out
}

func main() {
	fmt.Println(sieve(25))
}
