package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	var result [][]uint8

	for i := 0; i < dx; i++ {
		result = append(result, []uint8{})
		for j := 0; j < dy; j++ {

			if i%2 == 0 {
				result[i] = append(result[i],
					uint8((i+j)^3+(i+j)^2+(i+j)))
			} else {
				result[i] = append(result[i],
					uint8(i*j))
			}
		}
	}

	return result
}

func main() {
	pic.Show(Pic)

}
