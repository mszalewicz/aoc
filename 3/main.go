package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	regexFindMul, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)

	if err != nil {
		log.Fatal(err)
	}

	regexFindNumbers, err := regexp.Compile(`\d+`)

	if err != nil {
		log.Fatal(err)
	}

	validStrings := regexFindMul.FindAllString(string(content), -1)

	result := 0

	for _, validString := range validStrings {
		numbersText := regexFindNumbers.FindAllString(validString, -1)

		number1, err := strconv.Atoi(numbersText[0])

		if err != nil {
			log.Fatal(err)
		}

		number2, err := strconv.Atoi(numbersText[1])

		if err != nil {
			log.Fatal(err)
		}

		result += number1 * number2
	}

	fmt.Println(result)
}
