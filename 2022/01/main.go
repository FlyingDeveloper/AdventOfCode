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
	maxElf := -1
	for _, line := range lines {
		if line == "" {
			if currentElf > maxElf {
				maxElf = currentElf
			}
			currentElf = 0
		}

		asInt, _ := strconv.Atoi(line)
		currentElf += asInt
		if asInt > maxElf {
			maxElf = asInt
		}
	}

	fmt.Println(maxElf)
}
