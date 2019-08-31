package main

import "fmt"

func main() {
	a := [...]int{2, 3, 5, 7, 11, 13}
	b := []int{2, 3, 5, 7, 11, 13}
	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Printf("a = %T, value = %#v\n", a, a)
	fmt.Printf("b = %T, value = %#v\n", b, b)
	fmt.Printf("primes = %T, value = %#v\n", primes, primes)

	var s = primes[1:4]

	fmt.Printf("s = %T, value = %#v\n", s, s)

	s[0] = 11

	fmt.Printf("primes = %T, value = %#v\n", primes, primes)
	fmt.Printf("s = %T, value = %#v\n", s, s)
}
