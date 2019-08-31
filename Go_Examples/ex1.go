package main

import "fmt"

const (
	zero = (iota + 1) * 2
	one
	two
	three
	four
	five
)

func main() {
	fmt.Println(zero, one, two, three, four, five)
}
