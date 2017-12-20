package main

import "fmt"

func main() {
	s := []int{}
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1, 2, 3, 4, 5, 6, 7)
	printSlice(s)

	s = append(s, 8, 9, 10)
	printSlice(s)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
