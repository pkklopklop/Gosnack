package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func WordCount(s string) map[string]int {

	s = strings.ToLower(s)

	// Make a Regex to say we only want letters, numbers and spaces
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(s, "")

	splitString := strings.Split(processedString, " ")
	counts := make(map[string]int)

	for i := 0; i < len(splitString); i++ {
		counts[splitString[i]]++
	}

	return counts
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)

	fmt.Println(w)
}
