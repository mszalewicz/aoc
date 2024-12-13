package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")
	lines := strings.Split(string(content), "\n")

	result := 0

mainLoop:
	for i := 0; i < len(lines); i += 4 {
		ax, _ := strconv.Atoi(lines[i][12:14])
		ay, _ := strconv.Atoi(lines[i][18:20])

		bx, _ := strconv.Atoi(lines[i+1][12:14])
		by, _ := strconv.Atoi(lines[i+1][18:20])

		prize := strings.Split(lines[i+2], ",")

		leftPart := strings.Split(prize[0], "=")
		rightPart := strings.Split(prize[1], "=")

		prizex, _ := strconv.Atoi(leftPart[1])
		prizey, _ := strconv.Atoi(rightPart[1])

		// fmt.Println(ax, ay, bx, by, prizex, prizey)

		countA := 1
		countB := 1

		for {
			for {
				if ax*countA+bx*countB == prizex && ay*countA+by*countB == prizey {
					result += countA*3 + countB
					continue mainLoop
				}

				if ax*countA+bx*countB > prizex || ay*countA+by*countB > prizey {
					break
				}

				countB++
			}

			if ax*countA >= prizex || ay*countA >= prizey {
				break
			}

			countA++
			countB = 1
		}
	}

	fmt.Println(result)
}
