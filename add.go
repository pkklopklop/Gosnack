package main

import "fmt"

func addfunc(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(addfunc(4, 5))
	fmt.Println(addfunc(99, 98))
}
