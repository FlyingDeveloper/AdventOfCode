package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Assignment struct {
	min int
	max int
}

type Pair struct {
	left  Assignment
	right Assignment
}

func getInputLines() []string {
	workingDirectory, _ := os.Getwd()
	fileSystem := os.DirFS(workingDirectory)
	data, err := fs.ReadFile(fileSystem, "input")
	if err != nil {
		log.Fatal(err)
	}
	stringData := string(data)
	return strings.Split(stringData, "\n")
}

func getSampleInput() []string {
	return []string{
		"2-4,6-8",
		"2-3,4-5",
		"5-7,7-9",
		"2-8,3-7",
		"6-6,4-6",
		"2-6,4-8",
	}
}

func (a *Assignment) Contains(oa *Assignment) bool {
	if a.min <= oa.min && a.max >= oa.max {
		return true
	}

	return false
}

func DeserializeAssignment(input string) (Assignment, error) {
	r, _ := regexp.Compile("(\\d+)-(\\d+)")
	submatches := r.FindStringSubmatch(input)
	if len(submatches) != 3 {
		return Assignment{}, errors.New(fmt.Sprintf("Unable to parse input string %v", input))
	}

	min, _ := strconv.Atoi(submatches[1])
	max, _ := strconv.Atoi(submatches[2])
	if min > max {
		return Assignment{}, errors.New(fmt.Sprintf("Input has larger min than max %v", input))
	}

	return Assignment{
		min,
		max,
	}, nil
}

func DeserializePair(input string) (Pair, error) {
	r := regexp.MustCompile("([\\d-]+),([\\d-]+)")
	output := r.FindStringSubmatch(input)
	if len(output) != 3 {
		return Pair{}, errors.New(fmt.Sprintf("Unable to parse input %v", input))
	}

	left, err := DeserializeAssignment(output[1])
	if err != nil {
		return Pair{}, errors.New(fmt.Sprintf("Unable to deserialize left: %v", err.Error()))
	}
	right, err := DeserializeAssignment(output[2])
	if err != nil {
		return Pair{}, errors.New(fmt.Sprintf("Unable to deserialize right: %v", err.Error()))
	}

	return Pair{left, right}, nil
}

func (p *Pair) ContainsFullOverlap() bool {
	return p.left.Contains(&p.right) || p.right.Contains(&p.left)
}

func main() {
	input := getInputLines()
	fullyOverlappedPairs := 0
	for _, serializedPair := range input {
		if serializedPair == "" {
			continue
		}

		pair, err := DeserializePair(serializedPair)
		if err != nil {
			log.Fatalf("Unable to parse pair %v", err)
		}

		if pair.ContainsFullOverlap() {
			fullyOverlappedPairs++
		}
	}

	fmt.Println(fullyOverlappedPairs)
}
