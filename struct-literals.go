package main

import "fmt"

type Vertex struct {
	X, Y int
}

var v1 = Vertex{1, 2}
var v2 = Vertex{X: 1}
var v3 = Vertex{}
var p = &Vertex{1, 2}

func main() {
	fmt.Println(v1, v2, v3, p)
}
