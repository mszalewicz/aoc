package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	maxY := len(lines) - 1
	maxX := len(lines[0]) - 1

	stations := make(map[rune][]image.Point)

	for y, line := range lines {
		for x, char := range line {
			if unicode.IsLetter(char) || unicode.IsDigit(char) {
				stations[char] = append(stations[char], image.Pt(x, y))
			}
		}
	}

	result1 := 0
	result2 := 0
	visited := []image.Point{}
	visited2 := []image.Point{}

	for _, stationType := range stations {
		for i, coordinate := range stationType {

			coordinatesToCompare := slices.Concat(stationType[:i], stationType[i+1:])

			for _, toCompare := range coordinatesToCompare {
				xd := -(coordinate.X - toCompare.X)
				yd := -(coordinate.Y - toCompare.Y)

				{ // Part 1
					antinodeX := coordinate.X + 2*xd
					antinodeY := coordinate.Y + 2*yd

					if isInBounds(antinodeX, antinodeY, maxX, maxY) && !slices.Contains(visited, image.Pt(antinodeX, antinodeY)) {
						result1++
						visited = append(visited, image.Pt(antinodeX, antinodeY))
					}
				}

				{ // Part 2
					firstIteration := true
					antinodeX := 0
					antinodeY := 0

					for {
						if firstIteration {
							antinodeX = coordinate.X + 2*xd
							antinodeY = coordinate.Y + 2*yd
							firstIteration = false
						} else {
							antinodeX = antinodeX + xd
							antinodeY = antinodeY + yd
						}

						if isInBounds(antinodeX, antinodeY, maxX, maxY) {

							if !slices.Contains(visited2, image.Pt(antinodeX, antinodeY)) {
								visited2 = append(visited2, image.Pt(antinodeX, antinodeY))
								result2++
							}

							if !slices.Contains(visited2, image.Pt(coordinate.X, coordinate.Y)) {
								visited2 = append(visited2, image.Pt(coordinate.X, coordinate.Y))
								result2++
							}

							if !slices.Contains(visited2, image.Pt(toCompare.X, toCompare.Y)) {
								visited2 = append(visited2, image.Pt(toCompare.X, toCompare.Y))
								result2++
							}
						} else {
							break
						}
					}
				}
			}
		}
	}

	fmt.Println(result1)
	fmt.Println(result2)
}

func isInBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x <= maxX && y >= 0 && y <= maxY
}
