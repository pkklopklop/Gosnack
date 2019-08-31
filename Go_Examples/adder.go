package main

import "fmt"

func adder() (func() int, func() int) {
	sum := 1
	return func() int {
			sum++
			return sum
		},
		func() int {
			return sum
		}
}

func main() {
	inc, curr := adder()

	fmt.Println(curr())

	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(curr())
}
