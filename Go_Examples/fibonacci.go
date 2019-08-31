package main

import "fmt"

func fibonacci() func() int {
	prev, curr := 0, 1
	// newVal := 0
	// times := 1

	// return func() int {
	// 	if times == 1 {
	// 		times++
	// 		return 0
	// 	} else if times == 2 {
	// 		times++
	// 		return 1
	// 	} else {
	// 		times++
	// 		newVal = prev + curr
	// 		prev = curr
	// 		curr = newVal
	// 		return newVal
	// 	}
	// }

	return func() int {
		prev, curr = curr, prev+curr
		return prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i <= 10; i++ {
		fmt.Println(f())
	}
}
