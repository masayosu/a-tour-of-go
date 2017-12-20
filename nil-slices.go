package main

import "fmt"

func main() {
	//	var s []int = []int{1, 2}
	var s []int = []int{}
	//	s[0] = 1
	//	s[1] = 2

	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}
