package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	var slices [][]uint8

	for y := 0; y < dy; y++ {

		var row = make([]uint8, dx, dx)

		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
		}

		slices = append(slices, row)
	}

	return slices
}

func main() {
	pic.Show(Pic)
}
