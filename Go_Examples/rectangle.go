package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

func (rec rectangle) area() int {
	fmt.Println("Method area of rectangle struct.")
	return rec.width * rec.length
}

func area(rec rectangle) int {
	fmt.Println("function area")
	return rec.width * rec.length
}

func main() {
	rec := rectangle{3, 4}
	a := area(rec)
	fmt.Println(a)
	b := rec.area()
	fmt.Println(b)
}
