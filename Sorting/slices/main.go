package main

import "fmt"

func main() {
	var empSlice []int
	fmt.Printf("Type %T, Value %[1]v\n", empSlice)

	empSlice = append(empSlice, 2)
	empSlice = append(empSlice, 3)
	fmt.Println("Slice 1>", empSlice)
	fmt.Printf("length %d, capacity %d\n", len(empSlice), cap(empSlice))

	anotherSlice := make([]int, 5, 8)
	fmt.Printf("Type %T, Value %[1]v\n", anotherSlice)
	fmt.Printf("length %d, capacity %d\n", len(anotherSlice), cap(anotherSlice))

	anotherSlice[1] = 2
	anotherSlice[2] = 3
	fmt.Println("Slice 2>", anotherSlice)
}
