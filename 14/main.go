package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Info struct {
	VX int
	VY int
}

func main() {
	content, _ := os.ReadFile("input")

	lines := strings.Split(string(content), "\n")

	topLeft := 0
	topRight := 0
	bottomLeft := 0
	bottomRight := 0
	wide := 101
	tall := 103

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

	fmt.Println("Part 1:", topLeft*topRight*bottomLeft*bottomRight)

	robots := []Point{}
	robotsInfo := []Info{}

	for _, line := range lines {
		var x, y, vx, vy int

		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)

		robot := Point{X: x, Y: y}
		robots = append(robots, robot)
		robotInfo := Info{VX: vx, VY: vy}
		robotsInfo = append(robotsInfo, robotInfo)
	}

	counter := 1

mainLoop:
	for {
		for i := range robots {
			x := (robots[i].X + robotsInfo[i].VX)
			y := (robots[i].Y + robotsInfo[i].VY)

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

			robots[i].X = x
			robots[i].Y = y
		}

		visited := []Point{}
		for _, robot := range robots {
			toCheck := []Point{}
			toCheck = append(toCheck, Point{robot.X, robot.Y})

			if !slices.Contains(visited, Point{X: robot.X, Y: robot.Y}) {
				neighbours := 0
				visited = append(visited, Point{X: robot.X, Y: robot.Y})

				for len(toCheck) != 0 {

					left := Point{X: toCheck[0].X - 1, Y: toCheck[0].Y}
					right := Point{X: toCheck[0].X + 1, Y: toCheck[0].Y}
					up := Point{X: toCheck[0].X, Y: toCheck[0].Y - 1}
					down := Point{X: toCheck[0].X, Y: toCheck[0].Y + 1}

					newChecks := []Point{}

					for _, val := range toCheck[1:] {
						newChecks = append(newChecks, val)
					}

					if slices.Contains(robots, Point{left.X, left.Y}) && !slices.Contains(visited, Point{left.X, left.Y}) {
						neighbours++
						visited = append(visited, Point{left.X, left.Y})
						newChecks = append(newChecks, Point{left.X, left.Y})
					}

					if slices.Contains(robots, Point{right.X, right.Y}) && !slices.Contains(visited, Point{right.X, right.Y}) {
						neighbours++
						visited = append(visited, Point{right.X, right.Y})
						newChecks = append(newChecks, Point{right.X, right.Y})
					}

					if slices.Contains(robots, Point{up.X, up.Y}) && !slices.Contains(visited, Point{up.X, up.Y}) {
						neighbours++
						visited = append(visited, Point{up.X, up.Y})
						newChecks = append(newChecks, Point{up.X, up.Y})
					}

					if slices.Contains(robots, Point{down.X, down.Y}) && !slices.Contains(visited, Point{down.X, down.Y}) {
						neighbours++
						visited = append(visited, Point{down.X, down.Y})
						newChecks = append(newChecks, Point{down.X, down.Y})
					}

					toCheck = newChecks
				}

				if neighbours > 20 {
					fmt.Println("Current:", counter)
					printCurrentState(robots, wide, tall)
					continue mainLoop
				}
			}
		}
		counter++
	}
}

func printCurrentState(robots []Point, wide int, tall int) {
	for y := range tall {
		for x := range wide {
			if slices.Contains(robots, Point{x, y}) {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Print("\n")
	}
}
