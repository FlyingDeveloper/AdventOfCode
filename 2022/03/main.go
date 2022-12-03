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

func findDoublePackedItem(compartmentA string, compartmentB string) rune {
	for _, r := range []rune(compartmentA) {
		if strings.ContainsRune(compartmentB, r) {
			return r
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
	for _, rucksack := range input {
		if rucksack == "" {
			continue
		}
		compartmentA, compartmentB := getCompartments(rucksack)
		duplicatedItem := findDoublePackedItem(compartmentA, compartmentB)
		sumOfPriorities += getPriority(duplicatedItem)
	}
	fmt.Println(sumOfPriorities)
}
