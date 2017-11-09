package main

import "fmt"

func main() {
	var sum int = 0

	for sum < 1000 {
		sum += 1
	}

	fmt.Println(sum)

}
