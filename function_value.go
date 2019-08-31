package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(9, 5)
}

// func hypot(x, y float64) float64 {
// 	return math.Sqrt(x*x + y*y)
// }

// func triangleArea(x, y float64) float64 {
// 	return (x*x + y*y)
// }

// var hypot = func(x, y float64) float64 {
// 	return math.Sqrt(x*x + y*y)
// }

// var triangleArea = func(x, y float64) float64 {
// 	return (x*x + y*y)
// }

func main() {

	z := 99.99

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y + z)
	}

	triangleArea := func(x, y float64) float64 {
		return (x*x + y*y + z)
	}

	fmt.Println(hypot(3, 4))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Println(compute(triangleArea))
}
