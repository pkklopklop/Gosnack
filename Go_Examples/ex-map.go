package main

import "fmt"

type Student struct {
	name  string
	class string
	age   int64
}

func main() {
	var students map[string]Student

	students = make(map[string]Student)

	students["0000000001"] = Student{
		name:  "Test1 X",
		class: "Math-Science",
		age:   15,
	}

	students["0000000002"] = Student{
		name:  "Test2 Y",
		class: "Art-Language",
		age:   16,
	}

	students["0000000003"] = Student{
		name:  "Test3 Z",
		class: "Art-Language",
		age:   16,
	}

	fmt.Printf("%#v\n", students["0000000001"])
	fmt.Printf("%#v\n", students["0000000002"])
	fmt.Printf("%#v\n", students["0000000003"])
}
