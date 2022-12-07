package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

const totalDiskSize = 70000000
const sizeNeeded = 30000000

type Node struct {
	size     int
	name     string
	children []*Node
	parent   *Node
}

func CreateNode(name string) *Node {
	return &Node{
		size:     0,
		name:     name,
		children: []*Node{},
	}
}

func (n *Node) GetSize(path string) (int, map[string]int) {
	runningSum := 0
	sizeMap := map[string]int{}
	path = path + n.name
	for _, child := range n.children {
		if len(child.children) > 0 {
			childSum, childSizeMap := child.GetSize(path)
			runningSum += childSum
			for k, v := range childSizeMap {
				sizeMap[path+k] = v
			}
		} else {
			runningSum += child.size
		}
	}

	sizeMap[path] = runningSum
	return runningSum, sizeMap
}

func ProcessInput(input []string) (rootDirectory *Node) {
	cdCommandRegex := regexp.MustCompile("^\\$ cd (.+)")
	lsCommandRegex := regexp.MustCompile("^\\$ ls")
	lsOutputDirectory := regexp.MustCompile("^dir (.+)")
	lsOutputFile := regexp.MustCompile("^(\\d+) (.+)")
	var currentDirectory *Node
	for _, line := range input {
		if cdCommandRegex.MatchString(line) {
			submatches := cdCommandRegex.FindStringSubmatch(line)
			if len(submatches) != 2 {
				log.Fatalf("Unable to parse cd line %v", line)
			}

			dirName := submatches[1]
			if dirName == ".." {
				currentDirectory = currentDirectory.parent
				continue
			}

			newDirectory := CreateNode(dirName)
			if currentDirectory != nil {
				currentDirectory.children = append(currentDirectory.children, newDirectory)
				newDirectory.parent = currentDirectory
			} else {
				rootDirectory = newDirectory
			}
			currentDirectory = newDirectory
		} else if lsCommandRegex.MatchString(line) {
			// No need to do anything here
			continue
		} else if lsOutputDirectory.MatchString(line) {
			// I don't think we need to do anything here since the
			// cd command will build out the child directory node
			continue
		} else if lsOutputFile.MatchString(line) {
			submatches := lsOutputFile.FindStringSubmatch(line)
			if len(submatches) != 3 {
				log.Fatalf("Unable to parse file line %v\n", line)
			}
			size, _ := strconv.Atoi(submatches[1])
			name := submatches[2]
			node := CreateNode(name)
			node.size = size
			currentDirectory.children = append(currentDirectory.children, node)
		}
	}

	return
}

func PrintTree(rootNode *Node, level int) {
	name := rootNode.name
	if len(rootNode.children) > 0 {
		name = name + "/"
	}
	fmt.Printf("%"+strconv.Itoa(level)+"s\n", name)
	for _, child := range rootNode.children {
		PrintTree(child, level+2)
	}
}

func main() {
	tree := ProcessInput(getInputLines())
	totalUsedSize, sizeMap := tree.GetSize("")
	availableSpace := totalDiskSize - totalUsedSize
	sizeThatNeedsToBeDeleted := sizeNeeded - availableSpace

	bestDirectorySizeToDelete := int(math.Inf(1))
	for _, v := range sizeMap {
		if v >= sizeThatNeedsToBeDeleted {
			if v < bestDirectorySizeToDelete {
				bestDirectorySizeToDelete = v
			}
		}
	}
	fmt.Println(bestDirectorySizeToDelete)
}
