package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)

	for i := range a {
		a[i] = make([]uint8, dx)
	}

	for y, row := range a {
		for x := range row {
			row[x] = uint8(x * y)
		}
	}

	return a
}

func main() {
	pic.Show(Pic)
}
