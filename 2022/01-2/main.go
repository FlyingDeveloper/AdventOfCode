package main

import "io/fs"
import "os"
import "log"
import "strings"
import "fmt"
import "strconv"

func main() {
	workingDirectory, _ := os.Getwd()
	fileSystem := os.DirFS(workingDirectory)
	data, err := fs.ReadFile(fileSystem, "input")
	if err != nil {
		log.Fatal(err)
	}
	stringData := string(data)
	lines := strings.Split(stringData, "\n")

	currentElf := 0
	firstElf := -1
	secondElf := -1
	thirdElf := -1
	for _, line := range lines {
		if line == "" {
			if currentElf > thirdElf {
				thirdElf = currentElf
			}
			if thirdElf > secondElf {
				temp := secondElf
				secondElf = thirdElf
				thirdElf = temp
			}
			if secondElf > firstElf {
				temp := firstElf
				firstElf = secondElf
				secondElf = temp
			}
			currentElf = 0
		}

		asInt, _ := strconv.Atoi(line)
		currentElf += asInt
	}

	fmt.Println(firstElf + secondElf + thirdElf)
}
