package main

import "fmt"

func main() {
	sum := 0
	i := 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)

	names := []string{"game", "bank", "samui", "dog", "jiew"}

	for j, names := range names {
		fmt.Println(names)
		fmt.Println(j)
	}

	i = 0
	for {
		fmt.Println("Forever loop...")
		i++
		if i == 10 {
			break
		}
	}
}
