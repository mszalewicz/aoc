package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)
	mapping := []string{}

	for i := 0; i < len(input); i += 2 {
		fileLen, _ := strconv.Atoi(string(input[i]))

		for range fileLen {
			mapping = append(mapping, strconv.Itoa(i/2))
		}

		if i+1 < len(input) {
			emptyLen, _ := strconv.Atoi(string(input[i+1]))

			for range emptyLen {
				mapping = append(mapping, ".")
			}
		}
	}

	fromStart := 0
	fromEnd := len(mapping) - 1

	for {
		if fromStart == fromEnd {
			break
		}

		if _, err := strconv.Atoi(mapping[fromEnd]); err == nil && mapping[fromStart] == "." {
			mapping[fromStart], mapping[fromEnd] = mapping[fromEnd], mapping[fromStart]

			fromStart++
			fromEnd--

			continue
		}

		if mapping[fromStart] != "." {
			fromStart++
		}

		if _, err := strconv.Atoi(mapping[fromEnd]); err != nil {
			fromEnd--
		}
	}

	position := 0
	result := 0

	for {
		number, err := strconv.Atoi(mapping[position])

		if err != nil {
			break
		}

		result += number * position
		position++
	}

	fmt.Println(result)
}
