package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"slices"
	"strings"
)

type Moves struct {
	up    image.Point
	down  image.Point
	left  image.Point
	right image.Point
}

func main() {
	moves := Moves{
		up:    image.Pt(0, -1),
		down:  image.Pt(0, 1),
		left:  image.Pt(-1, 0),
		right: image.Pt(1, 0),
	}

	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	input := intoArray(strings.Split(string(content), "\n"))

	var currentPoint image.Point

	// find starting point
FindStartingPoint:
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			if input[y][x] == "^" {
				currentPoint = image.Point{X: x, Y: y}
				break FindStartingPoint
			}
		}
	}

	//originalCurrentPoint := currentPoint

	type mark struct {
		place     image.Point
		direction image.Point
	}

	visited := []image.Point{}

	currentDirection := moves.up
	maxX := len(input[0])
	maxY := len(input)

	for {
		if !slices.Contains(visited, currentPoint) {
			visited = append(visited, currentPoint)
		}

		currentPoint = currentPoint.Add(currentDirection)

		if currentPoint.X < 0 || currentPoint.X >= maxX || currentPoint.Y < 0 || currentPoint.Y >= maxY {
			break
		}

		if input[currentPoint.Y][currentPoint.X] == "#" {
			currentPoint = currentPoint.Sub(currentDirection)
			currentDirection = switchDirection(currentDirection, moves)
			currentPoint = currentPoint.Add(currentDirection)

			if currentPoint.X < 0 || currentPoint.X >= maxX || currentPoint.Y < 0 || currentPoint.Y >= maxY {
				break
			}
		}

	}

	fmt.Println("Part 1:", len(visited))

}

func switchDirection(currentDirection image.Point, moves Moves) image.Point {
	switch currentDirection {
	case moves.up:
		currentDirection = moves.right
	case moves.right:
		currentDirection = moves.down
	case moves.down:
		currentDirection = moves.left
	case moves.left:
		currentDirection = moves.up
	}

	return currentDirection
}

func intoArray(input []string) [][]string {
	lettersArray := [][]string{}

	for _, line := range input {
		letters := make([]string, 0)

		for _, letter := range line {
			letters = append(letters, string(letter))
		}

		lettersArray = append(lettersArray, letters)
	}

	return lettersArray
}
