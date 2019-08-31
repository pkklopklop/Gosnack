package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("i: %#v, f: %#v u: %#v", i, f, u)
}
