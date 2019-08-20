package main

import (
	"fmt"
	"os"
	"strconv"
)

const input = 987

var romanStr string
var romanMaps map[string]string

func main() {
	romanMaps = make(map[string]string)

	romanMaps["1"] = "I"
	romanMaps["5"] = "V"
	romanMaps["10"] = "X"
	romanMaps["50"] = "L"
	romanMaps["100"] = "C"
	romanMaps["500"] = "D"
	romanMaps["1000"] = "M"

	fmt.Println(romanMaps)

	inputStr := strconv.Itoa(input)

	if input > 1000 || input < 1 {
		fmt.Println("Error 1: input > 1000 OR input < 1.")
		os.Exit(3)
	} else if input == 1000 {
		fmt.Println("Result : M")
		os.Exit(0)
	}

	fmt.Println("Input String: " + inputStr)

	reverseInputStr := reverse(inputStr)

	for i := (len(reverseInputStr) - 1); i >= 0; i-- {

		fmt.Println("Process... " + string(reverseInputStr[i]))
		processRomanNumber(i, reverseInputStr[i])

		fmt.Println("Result of " + string(reverseInputStr[i]) + " at position " + strconv.Itoa(i) + "= " + romanStr)
	}

	fmt.Println("-----------------------")
	fmt.Println("-----------------------")
	fmt.Println("Result : " + romanStr)
	fmt.Println("-----------------------")
	fmt.Println("-----------------------")
	fmt.Println("Done!!!!!")

}

//Process Roman number for each position.
func processRomanNumber(pos int, numByte byte) {
	firstIter, secondIter, thirdIter := "", "", ""

	switch {
	case pos == 0:
		firstIter, secondIter, thirdIter = "1", "5", "10"
	case pos == 1:
		firstIter, secondIter, thirdIter = "10", "50", "100"
	case pos == 2:
		firstIter, secondIter, thirdIter = "100", "500", "1000"
	}

	numInt, err := strconv.Atoi(string(numByte))
	if err != nil {
		fmt.Println("Error 2: input has non-numeric characters - " + string(numByte))
		os.Exit(3)
	}

	switch {
	case numInt < 4:
		for numInt > 0 {
			romanStr += romanMaps[firstIter]
			numInt--
		}
	case numInt == 4:
		romanStr += romanMaps[firstIter] + romanMaps[secondIter]
	case numInt < 9:
		remainingFrom5 := numInt - 5
		romanStr += romanMaps[secondIter]

		for remainingFrom5 > 0 {
			romanStr += romanMaps[firstIter]
			remainingFrom5--
		}
	case numInt == 9:
		romanStr += romanMaps[firstIter] + romanMaps[thirdIter]
	}

	return
}

// Copy Util from other website to reverse string in GO.
// Credit: https://codelike.pro/reverse-a-string-in-go/
// Reverse string in GO.
func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
