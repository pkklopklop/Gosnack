package main

import "fmt"

type Student struct {
	name  string
	class string
	age   int64
}

func main() {

	students := map[string]Student{
		"0000000001": {
			name:  "Test1 X",
			class: "Math-Science",
			age:   15,
		},
		"0000000002": {
			name:  "Test2 Y",
			class: "Art-Language",
			age:   16,
		},
		"0000000003": {
			name:  "Test3 Z",
			class: "Art-Language",
			age:   16,
		},
	}

	delete(students, "0000000002")
	// fmt.Printf("%#v\n", students["0000000001"])
	// fmt.Printf("%#v\n", students["0000000002"])
	// fmt.Printf("%#v\n", students["0000000003"])
	fmt.Println(students)

	fmt.Println()

	x := "0000000004"
	if _, exists := students[x]; exists {
		fmt.Printf("Map students has student %s\n", x)
	} else {
		fmt.Printf("Map students has NO student %s\n", x)
	}
}
