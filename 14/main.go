package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

// type PointInfo struct {
// 	point Point
// 	vx    int
// 	vy    int
// }

func main() {
	content, _ := os.ReadFile("input")

	lines := strings.Split(string(content), "\n")

	topLeft     := 0
	topRight    := 0
	bottomLeft  := 0
	bottomRight := 0
	wide        := 101
	tall        := 103

	halfWide := wide / 2
	halfTall := tall / 2

	for _, line := range lines {
		var x, y, vx, vy int

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)

		for range 100 {
			x = (x + vx)
			y = (y + vy)

			switch {
			case x >= wide:
				x = x % wide
			case x < 0:
				x = wide + x
			}

			switch {
			case y >= tall:
				y = y % tall
			case y < 0:
				y = tall + y
			}
		}

		switch {
		case x < halfWide && y < halfTall:
			topLeft++
		case x > halfWide && y < halfTall:
			topRight++
		case x < halfWide && y > halfTall:
			bottomLeft++
		case x > halfWide && y > halfTall:
			bottomRight++
		}
	}

	fmt.Println("Part 1:", topLeft * topRight * bottomLeft * bottomRight)
}
