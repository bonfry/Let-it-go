package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var image [][]uint8

	for y := 0; y < dy; y++ {
		var imageRow []uint8
		for x := 0; x < dx; x++ {
			var cell uint8
			switch {
			case (x + y) > 128:
				cell = 122
			case (x + y) > 64:
				cell = uint8((x + y) / 2)
			case (x + y) < 32:
				cell = uint8(x * y)
			default:
				cell = 0
			}

			imageRow = append(imageRow, cell)
		}

		image = append(image, imageRow)
	}

	return image
}

func main() {
	pic.Show(Pic)
}
