package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance() {
	fmt.Println("Print point.")
}

//Embedding Point to struct ColorPoint.
type ColorPoint struct {
	Point
	Color string
}

func main() {

	p := Point{1, 2}
	q := Point{4, 6}
	c1 := ColorPoint{
		Point: p,
		Color: "Red",
	}
	c2 := ColorPoint{
		Point: q,
		Color: "Blue",
	}

	fmt.Println("p = ", c1)
	fmt.Println("q = ", c2)
	fmt.Println("c1.X = ", c1.X)
	fmt.Println("c1.Y = ", c1.Y)
	c1.Distance()
}
