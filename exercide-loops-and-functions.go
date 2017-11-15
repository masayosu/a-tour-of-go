package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	s := 0.0
	for {
		z = z - (z*z-x)/(2*z)
		if z-s < 1e-10 {
			break
		}
		s = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
