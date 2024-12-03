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

	text := string(content)

	regexFindNumbers, err := regexp.Compile(`\d+`)

	if err != nil {
		log.Fatal(err)
	}

	findRelevant, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|(do\(\))|(don't\(\))`)

	if err != nil {
		log.Fatal(err)
	}

	relevantEntries := findRelevant.FindAllString(text, -1)
	countFlag := true
	result := 0

	for _, relevantEntry := range relevantEntries {
		switch relevantEntry {
		case "do()":
			countFlag = true
		case "don't()":
			countFlag = false
		default:
			if countFlag {
				numbersText := regexFindNumbers.FindAllString(relevantEntry, -1)

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
		}
	}
	fmt.Println(result)
}
