package main

import "fmt"

func main() {
	var m map[string]string
	m = make(map[string]string)

	fmt.Println(m)

	m["Boy"] = "Peerayut Kiatpakdee"

	fmt.Println(m["Boy"])

	m["เทส"] = "ภาษาไทย"

	fmt.Println(m["เทส"])
}
