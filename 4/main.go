package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type lettersInWord struct {
	x int
	m int
	a int
	s int
}

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	lettersArray := make([][]string, 0)

	for _, line := range lines {

		letters := make([]string, 0)

		for _, letter := range line {
			letters = append(letters, string(letter))
		}

		lettersArray = append(lettersArray, letters)
	}

	result1 := 0

	rows := lettersArray
	rowsReversed := reverse(rows)

	diagonals := getAllDiagonals(lettersArray)
	diagonalReversed := reverse(diagonals)

	columns := getColumns(lettersArray)
	columnsReversed := reverse(columns)

	result1 += score(rows)
	result1 += score(rowsReversed)
	result1 += score(diagonals)
	result1 += score(diagonalReversed)
	result1 += score(columns)
	result1 += score(columnsReversed)

	fmt.Println(result1)

}

func isInBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}

func reverse(lines [][]string) [][]string {
	linesReversed := make([][]string, 0)

	for _, line := range lines {

		antiline := make([]string, len(line))

		for i := 0; i < len(line); i++ {
			antiline[len(antiline)-1-i] = line[i]
		}

		linesReversed = append(linesReversed, antiline)
	}

	return linesReversed
}

func getAllDiagonals(grid [][]string) [][]string {
	rows := len(grid)
	cols := len(grid[0])

	var diagonals [][]string

	for start := 0; start < rows; start++ {
		var diagonal []string
		r, c := start, 0
		for r < rows && c < cols {
			diagonal = append(diagonal, grid[r][c])
			r++
			c++
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	for start := 1; start < cols; start++ {
		var diagonal []string
		r, c := 0, start
		for r < rows && c < cols {
			diagonal = append(diagonal, grid[r][c])
			r++
			c++
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	for start := 0; start < rows; start++ {
		var diagonal []string
		r, c := start, cols-1
		for r < rows && c >= 0 {
			diagonal = append(diagonal, grid[r][c])
			r++
			c--
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	for start := cols - 2; start >= 0; start-- {
		var diagonal []string
		r, c := 0, start
		for r < rows && c >= 0 {
			diagonal = append(diagonal, grid[r][c])
			r++
			c--
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	return diagonals
}

func getColumns(grid [][]string) [][]string {
	rows := len(grid)
	cols := len(grid[0])

	var colArrays [][]string

	for col := 0; col < cols; col++ {
		var column []string
		for row := 0; row < rows; row++ {
			column = append(column, grid[row][col])
		}
		colArrays = append(colArrays, column)
	}

	return colArrays
}

func score(lines [][]string) int {
	result := 0

	re, err := regexp.Compile("XMAS")

	if err != nil {
		log.Fatal()
	}

	for _, letters := range lines {
		var line string

		for _, letter := range letters {
			line += letter
		}

		validStrings := re.FindAllString(line, -1)
		result += len(validStrings)
	}
	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
