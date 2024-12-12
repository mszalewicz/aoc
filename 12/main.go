package main

import (
	"fmt"
	"image"
	"os"
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

	result := 0

	for notSeen.Size() != 0 {
		start, _ := notSeen.Pop()
		area, perimeter := 0, 0

		currentType := land[start.Y][start.X]

		group := Stack[image.Point]{}
		group.Add(start)

		for group.Size() != 0 {
			node, _ := group.Pop()
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
				}
			}
		}
		result += area * perimeter
	}

	fmt.Println(result)
}

func isInBounds(x, y, maxX, maxY int) bool {
	return x >= 0 && x < maxX && y >= 0 && y < maxY
}
