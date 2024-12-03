package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Times:
//     - part 1: 514µs
//     - part 2: 655µs

func main() {
	content, err := os.ReadFile("input")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	reports := make([][]int, 0)
	numberOfSafeReports := 0

	convertReports(lines, &reports)

	for _, report := range reports {
		if isSafeReport(report) || isSafeWithOneRemoval(report) {
			numberOfSafeReports++
		}
	}

	fmt.Println(numberOfSafeReports)
}

func convertReports(lines []string, reports *[][]int) {
	for _, line := range lines {
		report := make([]int, 0)

		for _, text := range strings.Fields(line) {
			value, err := strconv.Atoi(text)

			if err != nil {
				log.Fatal(err)
			}
			report = append(report, value)
		}
		*reports = append(*reports, report)
	}
}

func isSafeReport(report []int) bool {
	increasing := true
	decreasing := true

	for i := 0; i < len(report)-1; i++ {
		diff := math.Abs(float64(report[i]) - float64(report[i+1]))

		switch {
		case diff < 1 || diff > 3:
			return false
		case report[i] < report[i+1]:
			decreasing = false
		case report[i] > report[i+1]:
			increasing = false
		}
	}

	return increasing || decreasing
}

func isSafeWithOneRemoval(report []int) bool {
	for i := 0; i < len(report); i++ {
		modifiedReport := make([]int, 0)
		modifiedReport = append(modifiedReport, report[:i]...)
		modifiedReport = append(modifiedReport, report[i+1:]...)

		if isSafeReport(modifiedReport) {
			return true
		}
	}
	return false
}
