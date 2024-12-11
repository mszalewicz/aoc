package main

import (
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("input")

	strings.Fields(string(content))

}
