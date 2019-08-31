package main

import "fmt"

type Person struct {
	name  string
	shoes []string
}

func (p Person) Walk() {
	fmt.Printf("%s is walking.\n", p.name)
}

func (p Person) Eat() {
	fmt.Printf("%s is eating.\n", p.name)
}

func (p Person) Greeting() {
	fmt.Printf("Hello, %s.\n", p.name)
}

func (p Person) Getter() string {
	return p.name + " " + p.shoes[0] + " " + p.shoes[1]
}

func (p *Person) Setter(s string) {
	p.name = s
	p.shoes = append(p.shoes, fmt.Sprintf("%s-Left shoe", p.name))
	p.shoes = append(p.shoes, fmt.Sprintf("%s-Right shoe", p.name))
	return
}

func main() {
	personA := Person{}

	personA.Setter("Boy")
	fmt.Println(personA.Getter())

	personA.Greeting()
	personA.Walk()
	personA.Eat()
}
