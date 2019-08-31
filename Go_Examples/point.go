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

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println("Distance function : ", Distance(p, q))
	fmt.Println("Distance method : ", p.Distance(q))
}
