package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type State struct {
	x         int
	y         int
	direction int
	cost      int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Point struct {
	x int
	y int
}

type Move struct {
	point     Point
	direction int
}

type Direction struct {
	dx int
	dy int
}

func main() {
	var (
		start Point
		end   Point
	)

	directions := []Direction{
		Direction{ 1,  0}, // right
		Direction{ 0,  1}, // down
		Direction{-1,  0}, // left
		Direction{ 0, -1}, // up
	}

	content, _ := os.ReadFile("input")
	terrain := [][]string{}

	for y, line := range strings.Split(string(content), "\n") {
		row := []string{}
		for x, char := range line {
			row = append(row, string(char))

			if char == 'S' { start = Point{x, y} }
			if char == 'E' { end   = Point{x, y} }

		}
		terrain = append(terrain, row)
	}

	result := 0
	visited := map[Move]bool{}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{start.x, start.y, 0, 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)

		if current.x == end.x && current.y == end.y {
			result = current.cost
			break
		}

		move := Move{Point{current.x, current.y}, current.direction}

		if visited[move] {
			continue
		}
		visited[move] = true


		// continue in directioin
		newX := current.x + directions[current.direction].dx
		newY := current.y + directions[current.direction].dy
		if terrain[newY][newX] != "#" {
			heap.Push(pq, State{newX, newY, current.direction, current.cost +1})
		}

		// rotate clockwise
		newDir := (current.direction + 1) % 4
		heap.Push(pq, State{current.x, current.y, newDir, current.cost + 1000})

		// rotate counterclockwise
		newDir = (current.direction + 3) % 4
		heap.Push(pq, State{current.x, current.y, newDir, current.cost + 1000})
	}

	fmt.Println(result)
}
