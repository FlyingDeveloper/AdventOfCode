package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	ROCK = iota
	PAPER
	SCISSOR
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
		"A Y",
		"B X",
		"C Z",
	}
}

func normalizePlay(position rune) int {
	if position == 'A' || position == 'X' {
		return ROCK
	}
	if position == 'B' || position == 'Y' {
		return PAPER
	}
	return SCISSOR
}

func getGameResult(input string) (score int) {
	r, err := regexp.Compile("[ABC] [XYZ]")
	if err != nil {
		log.Fatalf("Unable to compile regex: %v", err)
	}

	if !r.MatchString(input) {
		log.Fatalf("Unable to parse input %v", input)
	}

	runes := []rune(input)
	opponent := normalizePlay(runes[0])
	ours := normalizePlay(runes[2])
	switch ours {
	case ROCK:
		score += 1
		switch opponent {
		case ROCK:
			score += 3
		case PAPER:
			score += 0
		case SCISSOR:
			score += 6
		}
	case PAPER:
		score += 2
		switch opponent {
		case ROCK:
			score += 6
		case PAPER:
			score += 3
		case SCISSOR:
			score += 0
		}
	case SCISSOR:
		score += 3
		switch opponent {
		case ROCK:
			score += 0
		case PAPER:
			score += 6
		case SCISSOR:
			score += 3
		}
	}

	return
}

func main() {
	score := 0
	for _, line := range getInputLines() {
		if line == "" {
			continue
		}

		score += getGameResult(line)
	}

	fmt.Println(score)
}
