package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
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

FindStartingPoint:
	for y, line := range input {
		for x := 0; x < len(line); x++ {
			if input[y][x] == "^" {
				currentPoint = image.Point{X: x, Y: y}
				break FindStartingPoint
			}
		}
	}

	originalCurrentPoint := currentPoint

	type mark struct {
		place     image.Point
		direction image.Point
	}

	visited := []image.Point{}

	currentDirection := moves.up
	maxX := len(input[0])
	maxY := len(input)

breakWhenOutOfBounds:
	for {
		if !slices.Contains(visited, currentPoint) {
			visited = append(visited, currentPoint)
		}

		currentPoint = currentPoint.Add(currentDirection)

		if currentPoint.X < 0 || currentPoint.X >= maxX || currentPoint.Y < 0 || currentPoint.Y >= maxY {
			break
		}

		for {
			if input[currentPoint.Y][currentPoint.X] == "#" {
				currentPoint = currentPoint.Sub(currentDirection)
				currentDirection = switchDirection(currentDirection, moves)
				currentPoint = currentPoint.Add(currentDirection)

				if currentPoint.X < 0 || currentPoint.X >= maxX || currentPoint.Y < 0 || currentPoint.Y >= maxY {
					break breakWhenOutOfBounds
				}
			} else {
				break
			}

		}

	}

	fmt.Println("Part 1:", len(visited))

	var wg sync.WaitGroup
	result2 := 0

	for _, newObstacle := range visited[1:] {
		wg.Add(1)

		go func(result2 *int, moves Moves, originalCurrentPoint image.Point, newObstacle image.Point) {
			visitedInNewReality := []mark{}
			currentDirection := moves.up
			currentPoint := originalCurrentPoint
			guardStartingPosition := originalCurrentPoint

		breakWhenOutOfBounds:
			for {
				if slices.Contains(visitedInNewReality, mark{currentPoint, currentDirection}) {

					if newObstacle.X != guardStartingPosition.X || newObstacle.Y != guardStartingPosition.Y {
						*result2++
						break
					}
				}

				currentMotion := mark{currentPoint, currentDirection}
				visitedInNewReality = append(visitedInNewReality, currentMotion)

				currentPoint = currentPoint.Add(currentDirection)

				if currentPoint.X < 0 || currentPoint.X >= maxX || currentPoint.Y < 0 || currentPoint.Y >= maxY {
					break
				}

				for {
					if input[currentPoint.Y][currentPoint.X] == "#" || (newObstacle.X == currentPoint.X && newObstacle.Y == currentPoint.Y) {

						currentPoint = currentPoint.Sub(currentDirection)
						currentDirection = switchDirection(currentDirection, moves)
						currentPoint = currentPoint.Add(currentDirection)

						if currentPoint.X < 0 || currentPoint.X >= maxX || currentPoint.Y < 0 || currentPoint.Y >= maxY {
							break breakWhenOutOfBounds
						}
					} else {
						break
					}
				}
			}
			wg.Done()
		}(&result2, moves, originalCurrentPoint, newObstacle)
	}

	wg.Wait()

	fmt.Println("Part 2:", result2)
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
