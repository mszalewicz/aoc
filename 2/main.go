package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Times:
//     - part 1: 514µs
//     - part 2:

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	numberOfSafeReports := 0

	start := time.Now()

	for _, line := range lines {
		numbersText := strings.Fields(line)

		increasing := true
		decreasing := true
		safe := true

		for i := 0; i < len(numbersText)-1; i++ {
			var number1, number2 int
			var err error

			if number1, err = strconv.Atoi(numbersText[i]); err != nil {
				log.Fatal(err)
			}

			if number2, err = strconv.Atoi(numbersText[i+1]); err != nil {
				log.Fatal(err)
			}

			switch {
			case number1 < number2:
				if number2-number1 > 3 {
					safe = false
					break
				}
				decreasing = false
			case number1 > number2:
				if number1-number2 > 3 {
					safe = false
					break
				}
				increasing = false
			default:
				safe = false
				break
			}
		}

		if !increasing && !decreasing {
			safe = false
		}

		if safe {
			numberOfSafeReports++
		}

	}

	end := time.Since(start)
	fmt.Println(end)

	fmt.Println("Part 1:", numberOfSafeReports)
}
