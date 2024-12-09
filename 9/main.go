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

	mapping2 := make([]string, len(mapping))
	copy(mapping2, mapping)

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

	fmt.Println("Part 1:", result)

	// Idea behind part 2:
	// - first find consecutive numbers of the same time from the end of the mapping
	// - secondly find consecutive empty spaces with the same lenght as in previouse point
	// - swap position between empty spaces and numbers from the end
	// - continue from the smallest index of swapped numbers

	transplantNumberStartPosition := len(mapping2) - 1
	currentPosition := len(mapping2) - 1
	previousType := ""
	currentType := ""

	for {
		if currentPosition <= 0 {
			break
		}

		if previousType == "" || previousType == "." {
			previousType = mapping2[currentPosition]
		}

		if currentType == "" || currentType == "." {
			currentType = mapping2[currentPosition]
		}

		if mapping2[transplantNumberStartPosition] == "." || currentType == "." {
			transplantNumberStartPosition--
			currentPosition--
			continue
		}

		currentPosition--

		if previousType != mapping2[currentPosition] {
			emptySpotStart := 0
			emptySpotLen := 0
			for {
				if mapping2[emptySpotStart+emptySpotLen] == "." {
					emptySpotLen++

					if transplantNumberStartPosition-emptySpotLen < 0 {
						break
					}

					if emptySpotLen == transplantNumberStartPosition-currentPosition {
						for i := range emptySpotLen {
							mapping2[emptySpotStart+i], mapping2[transplantNumberStartPosition-i] = mapping2[transplantNumberStartPosition-i], mapping2[emptySpotStart+i]
						}
						transplantNumberStartPosition = transplantNumberStartPosition - emptySpotLen
						currentPosition = transplantNumberStartPosition
						previousType = mapping2[transplantNumberStartPosition]
						currentType = mapping2[currentPosition]
						break
					}
				} else if emptySpotStart >= transplantNumberStartPosition {
					transplantNumberStartPosition = currentPosition
					previousType = mapping2[currentPosition]
					currentType = mapping2[currentPosition]
					break
				} else {
					emptySpotStart += emptySpotLen + 1
					emptySpotLen = 0
				}
			}
		}
	}

	result2 := 0

	for i, val := range mapping2 {
		number, err := strconv.Atoi(val)

		if err != nil {
			continue
		}

		result2 += number * i
	}

	fmt.Println("Part 2:", result2)
}
