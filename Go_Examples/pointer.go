package main

import (
	"fmt"
)

func main() {
	i := 42
	fmt.Println("i: ", i)
	fmt.Printf("address i: %p\n", &i)

	var p *int
	fmt.Println("p: ", p)

	p = &i
	fmt.Println("p: ", p)

	i = 55
	fmt.Println("*p: ", *p)

	*p = 33
	fmt.Println("new i: ", i)

}
