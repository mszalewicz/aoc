package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")
	input := strings.Fields(string(content))

	list := LinkedList[string]{}

	for _, value := range input {
		list.Add(value)
	}

	numberOfBlinks := 25

	for range numberOfBlinks {
		current := list.Head
		for current != nil {
			switch {
			case current.Value == "0":
				current.Value = "1"
			case len(current.Value)%2 == 0:
				leftPart := current.Value[:len(current.Value)/2]
				rightPart := current.Value[(len(current.Value) / 2):]

				countZeros := 0
				for _, char := range rightPart {
					if string(char) == "0" {
						countZeros++
					} else {
						break
					}
				}

				if countZeros == len(rightPart) {
					rightPart = "0"
				} else {
					rightPart = rightPart[countZeros:]
				}

				newNode := &Node[string]{Value: rightPart}

				nextNode := current.Next
				if nextNode != nil {
					nextNode.Previous = newNode
				}

				newNode.Previous = current
				newNode.Next = nextNode

				current.Value = leftPart
				current.Next = newNode

				current = nextNode
				list.Size++
				continue

			default:
				number, _ := strconv.Atoi(current.Value)
				current.Value = strconv.Itoa(number * 2024)
			}
			current = current.Next
		}
	}

	fmt.Println(list.Size)
}
