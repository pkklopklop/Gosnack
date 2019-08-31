package main

import "fmt"

func main() {
	defer fmt.Println("Defer print...")
	defer fmt.Println("Defer print2...")

	fmt.Println("After defer...")

	fmt.Println("counting...")

	for i := 0; i < 10; i++ {
		// defer fmt.Println(i)
		defer func(newI int) {
			fmt.Println(newI)
		}(i)
	}
}
