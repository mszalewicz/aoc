package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")
	lines := strings.Split(string(content), "\n")

	result := 0

	for i := 0; i < len(lines); i += 4 {
		var aX, aY, bX, bY, prizeX, prizeY int

		fmt.Sscanf(lines[i+0], "Button A: X+%d, Y+%d", &aX, &aY)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bX, &bY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &prizeX, &prizeY)

		prizeX += 10000000000000
		prizeY += 10000000000000

		determinant := aX*bY - bX*aY
		determinantX := prizeX*bY - bX*prizeY
		determinantY := aX*prizeY - prizeX*aY

		determinantIsNonzero := (determinant != 0)
		xSolutionIsInteger := (determinantX % determinant == 0)
		ySolutionIsInteger := (determinantY % determinant == 0)

		if determinantIsNonzero && xSolutionIsInteger && ySolutionIsInteger {
			result += (determinantX/determinant)*3 + (determinantY / determinant)
		}
	}

	fmt.Println(result)
}
