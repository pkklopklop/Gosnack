package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Printf("s = %T, value = %#v\n", s, s)

	s = s[1:4]
	fmt.Println(s)
	fmt.Println(cap(s))

	s = s[:2]
	fmt.Println(s)
	fmt.Println(cap(s))

	s = s[1:]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = s[:4]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
