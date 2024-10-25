package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array Before function>", a)

	arrytModification(a)
	fmt.Println("Array After Function>", a)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice before function>", s)

	sliceModification(s)
	fmt.Println("Slice After Function>", s)
}

func arrytModification(s [5]int) {
	s[1] = 23
	fmt.Println("Array In Function> ", s)
}

func sliceModification(s []int) {
	s[1] = 23
	fmt.Println("Slice In Function>", s)
	s[2] = 24
}