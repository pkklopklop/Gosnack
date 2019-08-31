package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{
		Y: 2,
		X: 1,
	}
	// v := Vertex{
	// 	X: 2,
	// }
	fmt.Printf("v = %#v\n", v)

	v.X = 5
	fmt.Printf("v2 = %#v\n", v)

	var p *Vertex
	p = &v
	p.Y = 16 //or (*p).Y = 16
	fmt.Printf("v3 = %#v\n", v)
}
