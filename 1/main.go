package main

import (
	"fmt"
	"log"
	"math"
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

	values := strings.Fields(string(content))

	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for i := 0; i < len(values)-1; i += 2 {
		leftNumber, err := strconv.Atoi(values[i])

		if err != nil {
			log.Fatal(err)
		}

		rightNumber, err := strconv.Atoi(values[i+1])

		if err != nil {
			log.Fatal(err)
		}

		leftList = append(leftList, leftNumber)
		rightList = append(rightList, rightNumber)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	distance := 0

	for i, _ := range leftList {
		distance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	fmt.Println("Part 1:", distance)
}
