package main

import (
	"log"
	"strconv"
	"strings"
)

func buildForest(lines []string) [][]int {
	output := [][]int{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		lineSlice := make([]int, len(line))
		output = append(output, lineSlice)
		splitLine := strings.Split(line, "")
		for j, tree := range splitLine {
			treeHeight, err := strconv.Atoi(tree)
			if err != nil {
				log.Fatalf("Unable to handle tree height %v: %v\n", tree, err)
			}

			lineSlice[j] = treeHeight
		}
	}

	return output
}

func isTreeVisible(row int, col int, forest [][]int) bool {
	if len(forest) == 0 {
		log.Fatalf("Forest is empty")
	}

	// Check edge
	forestWidth := len(forest[0])
	forestHeight := len(forest)
	if row == 0 || row == forestHeight-1 || col == 0 || col == forestWidth-1 {
		return true
	}

	currentTreeHeight := forest[row][col]
	// Check left
	leftVisible := true
	for c := col - 1; c >= 0; c-- {
		if currentTreeHeight <= forest[row][c] {
			leftVisible = false
			break
		}
	}
	if leftVisible {
		return true
	}

	// Check right
	rightVisible := true
	for c := col + 1; c < forestWidth; c++ {
		if currentTreeHeight <= forest[row][c] {
			rightVisible = false
			break
		}
	}
	if rightVisible {
		return true
	}

	// Check up
	upVisible := true
	for r := row - 1; r >= 0; r-- {
		if currentTreeHeight <= forest[r][col] {
			upVisible = false
			break
		}
	}
	if upVisible {
		return true
	}

	// Check down
	downVisible := true
	for r := row + 1; r < forestHeight; r++ {
		if currentTreeHeight <= forest[r][col] {
			downVisible = false
			break
		}
	}
	if downVisible {
		return true
	}

	return false
}

func main() {
	lines := getInputLines()
	forest := buildForest(lines)

	visibleTrees := 0
	for r := range forest {
		for c := range forest[r] {
			if ok := isTreeVisible(r, c, forest); ok {
				visibleTrees++
			}
		}
	}

	log.Printf("Nuber of visible trees: %v\n", visibleTrees)
}
