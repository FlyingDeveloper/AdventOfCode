package main

import (
	"log/slog"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const sampleInput string = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func readInput() string {
	bytes, err := os.ReadFile("input")
	if err != nil {
		slog.Error("Unable to read file", "error", err)
		panic("Unable to read file")
	}

	return string(bytes)
}

func main() {
	input := readInput()
	slog.With("result", part1(input)).Info("Part 1")
	slog.With("result", part2(input)).Info("Part 2")
}

func part1(input string) int {
	left, right := loadLeftAndRight(input)

	sort.Float64s(left)
	sort.Float64s(right)
	var runningSum float64

	for i, _ := range left {
		runningSum = runningSum + math.Abs(left[i]-right[i])
	}

	return int(runningSum)
}

func part2(input string) int {
	left, right := loadLeftAndRight(input)

	var runningSum float64
	sort.Float64s(right)

	for _, i := range left {
		var rightCount float64
		for _, j := range right {
			if j > i {
				break
			}
			if j == i {
				rightCount++
			}
		}

		runningSum += rightCount * i
	}

	return int(runningSum)
}

func loadLeftAndRight(input string) (left []float64, right []float64) {
	lineRegex := regexp.MustCompile("(\\d+)\\s+(\\d+)")

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		regexResult := lineRegex.FindAllStringSubmatch(line, -1)
		if len(regexResult) == 0 {
			continue
		}

		lineLeft, _ := strconv.ParseFloat(regexResult[0][1], 64)
		lineRight, _ := strconv.ParseFloat(regexResult[0][2], 64)
		left = append(left, lineLeft)
		right = append(right, lineRight)
	}
	return left, right
}
