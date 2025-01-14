package main

import "fmt"

func maxElem(ar []int) int {
	maxVal := ar[0]

	for i := 1; i < len(ar); i++ {
		if ar[i] > maxVal {
			maxVal = ar[i]
		}
	}

	return maxVal
}

func main() {
	arr := []int{23, 843, 234, 454}
	fmt.Println(maxElem(arr))
}