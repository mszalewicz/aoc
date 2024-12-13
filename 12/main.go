package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")

	land := [][]string{}
	notSeen := Stack[image.Point]{}

	for y, line := range strings.Split(string(content), "\n") {
		row := []string{}
		for x, char := range line {
			row = append(row, string(char))
			notSeen.Add(image.Pt(x, y))
		}
		land = append(land, row)
	}

	maxX := len(land[0])
	maxY := len(land)

	directions := []image.Point{
		image.Pt(-1, 0),
		image.Pt(+1, 0),
		image.Pt(0, -1),
		image.Pt(0, +1),
	}

	result1 := 0

	sets := [][]image.Point{}

	// Part 1
	for notSeen.Size() != 0 {
		start, _ := notSeen.Pop()
		area := 0
		perimeter := 0

		currentType := land[start.Y][start.X]

		group := Stack[image.Point]{}
		group.Add(start)

		set := []image.Point{}
		for group.Size() != 0 {
			node, _ := group.Pop()
			set = append(set, image.Pt(node.X, node.Y))
			area++
			perimeter += 4

			for _, direction := range directions {
				neighbour := image.Pt(node.X+direction.X, node.Y+direction.Y)
				if isInBounds(neighbour.X, neighbour.Y, maxX, maxY) && land[neighbour.Y][neighbour.X] == currentType {

					perimeter--

					if !notSeen.Contains(neighbour) {
						continue
					}

					notSeen.Remove(neighbour)
					group.Add(neighbour)
					if !slices.Contains(set, image.Pt(neighbour.X, neighbour.Y)) {
						set = append(set, image.Pt(neighbour.X, neighbour.Y))
					}
				}
			}
		}
		sets = append(sets, set)
		result1 += area * perimeter
	}

	fmt.Println(result1)

	// Part 2
	result2 := 0

	for _, set := range sets {
		uniques := uniquePoints(set)

		area := len(uniques)
		sides := 0

		for _, point := range uniques {

			// outer corners
			if !slices.Contains(uniques, image.Pt(point.X - 1, point.Y)) && !slices.Contains(uniques, image.Pt(point.X, point.Y - 1)) {
				sides++
			}

			if !slices.Contains(uniques, image.Pt(point.X + 1, point.Y)) && !slices.Contains(uniques, image.Pt(point.X, point.Y - 1)) {
				sides++
			}

			if !slices.Contains(uniques, image.Pt(point.X - 1, point.Y)) && !slices.Contains(uniques, image.Pt(point.X, point.Y + 1)) {
				sides++
			}

			if !slices.Contains(uniques, image.Pt(point.X + 1, point.Y)) && !slices.Contains(uniques, image.Pt(point.X, point.Y + 1)) {
				sides++
			}

			// inner corners
			if slices.Contains(uniques, image.Pt(point.X - 1, point.Y)) && slices.Contains(uniques, image.Pt(point.X, point.Y - 1)) && !slices.Contains(uniques, image.Pt(point.X - 1, point.Y - 1)) {
				sides++
			}

			if slices.Contains(uniques, image.Pt(point.X + 1, point.Y)) && slices.Contains(uniques, image.Pt(point.X, point.Y - 1)) && !slices.Contains(uniques, image.Pt(point.X + 1, point.Y - 1)) {
				sides++
			}

			if slices.Contains(uniques, image.Pt(point.X - 1, point.Y)) && slices.Contains(uniques, image.Pt(point.X, point.Y + 1)) && !slices.Contains(uniques, image.Pt(point.X - 1, point.Y + 1)) {
				sides++
			}

			if slices.Contains(uniques, image.Pt(point.X + 1, point.Y)) && slices.Contains(uniques, image.Pt(point.X, point.Y + 1)) && !slices.Contains(uniques, image.Pt(point.X + 1, point.Y + 1)) {
				sides++
			}
		}
		result2 += area * sides
	}

	fmt.Println(result2)
}

func isInBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}

type Point struct {
	X int
	Y int
}

func uniquePoints(points []image.Point) []image.Point {
	visited := map[Point]bool{}
	uniques := []image.Point{}

	for _, point := range points {
		if _, present := visited[Point{X: point.X, Y: point.Y}]; !present {
			visited[Point{X: point.X, Y: point.Y}] = true
			uniques = append(uniques, image.Pt(point.X, point.Y))
		}
	}

	return uniques
}