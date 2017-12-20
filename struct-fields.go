package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v Vertex
	v = Vertex{1, 2}
	fmt.Println(v.X)
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
