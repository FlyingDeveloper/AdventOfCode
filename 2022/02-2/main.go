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

const (
	WIN = iota
	LOSE
	DRAW
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

func normalizeOutcome(token rune) int {
	if token == 'X' {
		return LOSE
	}
	if token == 'Y' {
		return DRAW
	}
	return WIN
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
	desiredOutcome := normalizeOutcome(runes[2])
	switch desiredOutcome {
	case LOSE:
		score += 0
		switch opponent {
		case ROCK:
			score += 3
		case PAPER:
			score += 1
		case SCISSOR:
			score += 2
		}
	case DRAW:
		score += 3
		switch opponent {
		case ROCK:
			score += 1
		case PAPER:
			score += 2
		case SCISSOR:
			score += 3
		}
	case WIN:
		score += 6
		switch opponent {
		case ROCK:
			score += 2
		case PAPER:
			score += 3
		case SCISSOR:
			score += 1
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
