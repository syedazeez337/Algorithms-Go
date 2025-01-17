package main

import "fmt"

func firstNonRepeating(str string) string {
	charCount := make(map[rune]int)
	for _, elem := range str {
		charCount[elem]++
	}
	// fmt.Println(charCount)

	for _, char := range str {
		if charCount[char] == 1 {
			return string(char)
		}
	}

	return "None"
}



func main() {
	fmt.Println(firstNonRepeating("swiss"))
}