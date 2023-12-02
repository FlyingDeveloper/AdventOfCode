package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"rob-hamilton.com/AdventOfCode/internal"
)

const part1SampleInput string = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

const part2SampleInput string = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func main() {
	//part1()
	part2()
}

func part1() {
	input, err := internal.GetInputAsStringArray("./01/input")
	if err != nil {
		log.Fatalf("Failure reading inputer: %v", err)
		return
	}
	//input := strings.Split(sampleInput, "\n")

	runningTotal := sumLines(input)

	fmt.Printf("Part 1: %v\n", runningTotal)
}

func sumLines(input []string) int {
	firstDigitRegex := regexp.MustCompile(`^[^\d]*(\d)`)
	lastDigitRegex := regexp.MustCompile(`(\d)[^\d]*$`)
	runningTotal := 0
	for _, line := range input {
		result := firstDigitRegex.FindStringSubmatch(line)
		digits := ""
		if len(result) > 0 {
			digits += result[1]
		}

		result = lastDigitRegex.FindStringSubmatch(line)
		if len(result) > 1 {
			digits += result[1]
		}

		value, _ := strconv.Atoi(digits)
		runningTotal += value
	}
	return runningTotal
}

func part2() {
	input, _ := internal.GetInputAsStringArray("./01/input")
	//input := strings.Split(part2SampleInput, "\n")
	substitutions := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
	}

	replacedLines := []string{}
	for _, line := range input {
		fmt.Printf("--%s--\n", line)
		lowestReplacementIndex := math.MaxInt
		highestReplacementIndex := -1
		firstStringToSubstituteIn := ""
		lastStringToSubstituteIn := ""
		lengthOfFirstStringBeingReplaced := 0
		lengthOfLastStringBeingReplaced := 0
		for k, v := range substitutions {
			index := strings.Index(line, k)
			if index != -1 && index < lowestReplacementIndex {
				lowestReplacementIndex = index
				firstStringToSubstituteIn = v
				lengthOfFirstStringBeingReplaced = len(k)
			}

			if index != -1 && index > highestReplacementIndex {
				highestReplacementIndex = index
				lastStringToSubstituteIn = v
				lengthOfLastStringBeingReplaced = len(k)
			}
		}

		lastReplaced := ""
		if lowestReplacementIndex != math.MaxInt {
			firstReplaced := fmt.Sprintf("%s%s%s", line[0:lowestReplacementIndex], firstStringToSubstituteIn, line[lowestReplacementIndex+lengthOfFirstStringBeingReplaced:])
			if highestReplacementIndex < lowestReplacementIndex+lengthOfFirstStringBeingReplaced {
				lastReplaced = firstReplaced
			} else {
				highestReplacementIndex = highestReplacementIndex - lengthOfFirstStringBeingReplaced + len(firstStringToSubstituteIn)
				lastReplaced = fmt.Sprintf("%s%s%s", firstReplaced[0:highestReplacementIndex], lastStringToSubstituteIn, firstReplaced[highestReplacementIndex+lengthOfLastStringBeingReplaced:])
			}
		} else {
			lastReplaced = line
		}
		fmt.Printf("++%s++\n", lastReplaced)
		replacedLines = append(replacedLines, lastReplaced)
	}

	fmt.Println(sumLines(replacedLines))
}
