package main

import (
	"fmt"
)

func Reverse(s string) string {
	word := ""
	for i:=len(s)-1; i >= 0; i-- {
		word += string(s[i])
	}
	return word
}

func main() {
	fmt.Println(Reverse("cat")) // "tac"
	fmt.Println(Reverse("ca t")) // "t ac"
	fmt.Println(Reverse("alphabet")) // "tebahpla"
}