package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

const startOfPacketMarkerWidth = 14

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
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
}

func main() {
	input := getInputLines()
	for _, line := range input {
		for i := startOfPacketMarkerWidth; i < len(line); i++ {
			possibleMarker := line[i-startOfPacketMarkerWidth : i]
			foundCharacters := map[byte]interface{}{}
			hasRepeats := false
			for j := 0; j < len(possibleMarker); j++ {
				if _, ok := foundCharacters[possibleMarker[j]]; ok {
					// This is a repeat
					hasRepeats = true
					break
				}
				foundCharacters[possibleMarker[j]] = true
			}
			if !hasRepeats {
				// We found it!
				fmt.Println(i)
				break
			}
		}
	}
}
