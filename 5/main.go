package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	rules := make(map[int][]int, 0)
	entries := make([][]int, 0)

	// reorganize data from input
	for _, line := range lines {
		switch {
		case strings.Contains(line, "|"):
			numbersText := strings.Split(line, "|")

			key, err := strconv.Atoi(numbersText[0])

			if err != nil {
				log.Fatal(err)
			}

			value, err := strconv.Atoi(numbersText[1])

			if err != nil {
				log.Fatal(err)
			}

			_, keyExists := rules[key]

			if keyExists {
				rules[key] = append(rules[key], value)
			} else {
				rules[key] = []int{value}
			}

		case strings.Contains(line, ","):
			numbersText := strings.Split(line, ",")

			entry := []int{}

			for _, rune := range numbersText {
				number, err := strconv.Atoi(rune)

				if err != nil {
					log.Fatal(err)
				}

				entry = append(entry, number)
			}

			entries = append(entries, entry)
		}
	}

	result1 := part1(rules, entries)
	result2 := part2(rules, entries)

	fmt.Println("Part1:", result1)
	fmt.Println("Part2:", result2)
}

func part1(rules map[int][]int, entries [][]int) int {
	result := 0

EntriesLoop:
	for _, entry := range entries {
		visited := []int{}
		for _, value := range entry {

			for _, rule := range rules[value] {
				if slices.Contains(visited, rule) {
					continue EntriesLoop
				}
			}
			visited = append(visited, value)

		}

		result += entry[len(entry)/2]
	}

	return result
}

func part2(rules map[int][]int, entries [][]int) int {
	result := 0

EntriesLoop:
	for _, entry := range entries {
		visited := []int{}
		for _, value := range entry {
			for _, rule := range rules[value] {

				// if entry has incorrect order:
				if slices.Contains(visited, rule) {
					// reorder to correct form and add middle value to result:
					reorderEntry(&entry, rules)
					result += entry[len(entry)/2]

					continue EntriesLoop
				}

			}
			visited = append(visited, value)
		}
	}

	return result
}

func reorderEntry(entry *[]int, rules map[int][]int) {
	for {
		swapped := false
		for i := 0; i < len(*entry)-1; i++ {
			if shouldComeBefore((*entry)[i], (*entry)[i+1], rules) {
				(*entry)[i], (*entry)[i+1] = (*entry)[i+1], (*entry)[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func shouldComeBefore(currentValue int, nextValue int, rulesMap map[int][]int) bool {
	rules := rulesMap[currentValue]

	for _, rule := range rules {
		if rule == nextValue {
			return true
		}
	}
	return false
}
