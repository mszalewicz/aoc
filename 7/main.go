package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	entries := make(map[int][]int, 0)

	for _, line := range lines {
		parts := strings.Split(line, ":")

		result, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatal(err)
		}

		var numbers []int

		for _, numberText := range strings.Fields(parts[1]) {
			number, err := strconv.Atoi(numberText)

			if err != nil {
				log.Fatal(err)
			}

			numbers = append(numbers, number)
		}

		entries[result] = numbers
	}

	start := time.Now()
	result := 0

	for key, values := range entries {
		currentResults := []int{0}
		for _, value := range values {
			newResults := []int{}
			for _, currentResult := range currentResults {
				r1 := currentResult + value
				r2 := currentResult * value
				r3Text := strconv.Itoa(currentResult) + strconv.Itoa(value)
				r3, _ := strconv.Atoi(r3Text)

				if r1 <= key {
					newResults = append(newResults, r1)
				}

				if r2 <= key {
					newResults = append(newResults, r2)
				}

				if r3 <= key {
					newResults = append(newResults, r3)
				}
			}
			currentResults = newResults
		}

		if slices.Contains(currentResults, key) {
			result += key
		}
	}

	end := time.Since(start)
	fmt.Println(end)
	fmt.Println(result)
}

func isInBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}
