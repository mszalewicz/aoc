package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")

	terrain := [][]int{}
	validPoints := Queue[image.Point]{}

	// input into 2d int array
	for y, line := range strings.Split(string(content), "\n") {
		level := []int{}

		for x, char := range line {
			number, _ := strconv.Atoi(string(char))
			level = append(level, number)

			if number == 0 {
				validPoints.Enqueue(image.Pt(x, y))
			}
		}

		terrain = append(terrain, level)
	}

	validPoints2 := Queue[image.Point]{}

	for _, val := range validPoints.Elements {
		validPoints2.Enqueue(image.Pt(val.X, val.Y))
	}

	// Part 1
	maxY := len(terrain)
	maxX := len(terrain[0])
	result := 0

	for _, start := range validPoints.Elements {
		toCheck := Queue[image.Point]{}
		toCheck.Enqueue(image.Pt(start.X, start.Y))

		visited := []image.Point{}

		for {
			if toCheck.IsEmpty() {
				break
			}

			point, _ := toCheck.Dequeue()

			checkDirectionWithoutRepetition(point.X, point.Y, point.X-1, point.Y, maxX, maxY, &terrain, &toCheck, &visited, &result)
			checkDirectionWithoutRepetition(point.X, point.Y, point.X+1, point.Y, maxX, maxY, &terrain, &toCheck, &visited, &result)
			checkDirectionWithoutRepetition(point.X, point.Y, point.X, point.Y-1, maxX, maxY, &terrain, &toCheck, &visited, &result)
			checkDirectionWithoutRepetition(point.X, point.Y, point.X, point.Y+1, maxX, maxY, &terrain, &toCheck, &visited, &result)
		}
	}

	fmt.Println(result)

	// Part 2
	result2 := 0

	for {
		if validPoints2.IsEmpty() {
			break
		}

		point, _ := validPoints2.Dequeue()

		checkDirection(point.X, point.Y, point.X-1, point.Y, maxX, maxY, &terrain, &validPoints2, &result2)
		checkDirection(point.X, point.Y, point.X+1, point.Y, maxX, maxY, &terrain, &validPoints2, &result2)
		checkDirection(point.X, point.Y, point.X, point.Y-1, maxX, maxY, &terrain, &validPoints2, &result2)
		checkDirection(point.X, point.Y, point.X, point.Y+1, maxX, maxY, &terrain, &validPoints2, &result2)

	}

	fmt.Println(result2)
}

func checkDirectionWithoutRepetition(currentX int, currentY int, x int, y int, maxX int, maxY int, terrain *[][]int, pointsQueue *Queue[image.Point], visited *[]image.Point, result *int) {
	currentValue := (*terrain)[currentY][currentX]

	if isInBounds(x, y, maxX, maxY) {
		newValue := (*terrain)[y][x]

		if currentValue == 8 && newValue == 9 && !slices.Contains(*visited, image.Pt(x, y)) {
			*result++
			*visited = append(*visited, image.Pt(x, y))
			return
		}

		if newValue-currentValue == 1 && !slices.Contains(*visited, image.Pt(x, y)) {
			pointsQueue.Enqueue(image.Pt(x, y))
			*visited = append(*visited, image.Pt(x, y))
			return
		}
	}
}

func checkDirection(currentX int, currentY int, x int, y int, maxX int, maxY int, terrain *[][]int, pointsQueue *Queue[image.Point], result *int) {
	currentValue := (*terrain)[currentY][currentX]

	if isInBounds(x, y, maxX, maxY) {
		newValue := (*terrain)[y][x]

		if currentValue == 8 && newValue == 9 {
			*result++
			return
		}

		if newValue-currentValue == 1 {
			pointsQueue.Enqueue(image.Pt(x, y))
			return
		}
	}
}

func isInBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}
