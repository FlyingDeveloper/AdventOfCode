package main

import (
	"container/list"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	quantity    int
	source      string
	destination string
}

func DeserializeInstruction(input string) (Instruction, error) {
	r := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	submatches := r.FindStringSubmatch(input)
	if len(submatches) != 4 {
		return Instruction{}, errors.New(fmt.Sprintf("Unable to parse input %v", input))
	}
	quantity, _ := strconv.Atoi(submatches[1])
	return Instruction{
		quantity:    quantity,
		source:      submatches[2],
		destination: submatches[3],
	}, nil
}

func buildStacks(input []string) (map[string]*list.List, error) {
	r := regexp.MustCompile("\\s+(\\d+)\\s+")
	lastLine := input[len(input)-1]
	isLastLineValid := r.MatchString(lastLine)
	if !isLastLineValid {
		return nil, errors.New("Last line of input should be a series of digits separated by whitespace")
	}

	splitLastLine := strings.Split(lastLine, " ")
	stackNames := []string{}
	for _, s := range splitLastLine {
		if s == "" {
			continue
		}
		stackNames = append(stackNames, s)
	}

	stacks := map[string]*list.List{}
	for _, stackName := range stackNames {
		stacks[stackName] = list.New()
	}

	numStacks := len(stacks)
	stackWidth := 4
	lineWidth := numStacks & stackWidth
	for _, containerLine := range input {
		hasContainers, _ := regexp.MatchString("\\[\\w+\\]", containerLine)
		if !hasContainers {
			continue
		}
		paddedContainerLine := fmt.Sprintf("%-"+strconv.Itoa(lineWidth)+"s", containerLine)
		for i := 0; i < numStacks; i++ {
			containerBegin := i * stackWidth
			containerEnd := containerBegin + stackWidth
			var currentContainer string
			if containerBegin > len(paddedContainerLine)-1 {
				continue
			}
			if containerEnd > len(paddedContainerLine)-1 {
				currentContainer = paddedContainerLine[containerBegin:]
			} else {
				currentContainer = paddedContainerLine[containerBegin:containerEnd]
			}
			trimmedCurrentContainer := strings.TrimSpace(currentContainer)
			if trimmedCurrentContainer != "" {
				stacks[strconv.Itoa(i+1)].PushBack(trimmedCurrentContainer)
			}
		}
	}

	return stacks, nil
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
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		"1   2   3 ",
		"",
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}
}

func ExecuteInstruction(i Instruction, stacks map[string]*list.List) {
	source := stacks[i.source]
	dest := stacks[i.destination]
	for j := 0; j < i.quantity; j++ {
		elementToMove := source.Front()
		source.Remove(elementToMove)
		dest.PushFront(elementToMove.Value)
	}
}

func main() {
	lines := getInputLines()
	stackLines := []string{}
	buildingStacks := true
	instructions := []Instruction{}
	for _, line := range lines {
		if line == "" {
			buildingStacks = false
			continue
		}
		if buildingStacks {
			stackLines = append(stackLines, line)
			continue
		}

		i, err := DeserializeInstruction(line)
		if err != nil {
			log.Fatal(err.Error())
		}
		instructions = append(instructions, i)
	}

	stacks, err := buildStacks(stackLines)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, instruction := range instructions {
		ExecuteInstruction(instruction, stacks)
	}

	output := ""
	for i := 1; i < len(stacks)+1; i++ {
		stringIndex := strconv.Itoa(i)
		value := stacks[stringIndex].Front().Value
		output = output + value.(string)
	}
	fmt.Println(output)
}
