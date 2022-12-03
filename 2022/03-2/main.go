package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

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
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
}

func getCompartments(rucksack string) (string, string) {
	splitPoint := (len(rucksack) / 2)
	runes := []rune(rucksack)
	return string(runes[:splitPoint]), string(runes[splitPoint:])
}

func findTriplePackedItem(compartmentA string, compartmentB string, compartmentC string) rune {
	for _, r := range []rune(compartmentA) {
		if strings.ContainsRune(compartmentB, r) {
			if strings.ContainsRune(compartmentC, r) {
				return r
			}
		}
	}

	log.Fatal("Unable to find a duplicated item")
	return -1
}

func getPriority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return int(r) - 96
	} else if r >= 'A' && r <= 'Z' {
		return int(r) - 38
	}

	return -1
}

func main() {
	input := getInputLines()

	sumOfPriorities := 0
	for i := 0; i < len(input)-1; i++ {
		compartmentA := input[i]
		i++
		compartmentB := input[i]
		i++
		compartmentC := input[i]
		duplicatedItem := findTriplePackedItem(compartmentA, compartmentB, compartmentC)
		sumOfPriorities += getPriority(duplicatedItem)
	}
	fmt.Println(sumOfPriorities)
}
