package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"

	svg "github.com/ajstarks/svgo"
)

const (
	squareWidth  = 20
	squareHeight = 20
)

type Position struct {
	x      int
	y      int
	height int
	h      float64
	g      float64
	parent *Position
}

func (p *Position) getKey() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

type ByFScore []*Position

func (a ByFScore) Len() int           { return len(a) }
func (a ByFScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFScore) Less(i, j int) bool { return (a[i].h + a[i].g) < (a[j].h + a[j].g) }

func getNumericHeight(character string) int {
	if len(character) != 1 {
		log.Fatalf("Unable to determine height of %v", character)
	}

	if character == "S" {
		return getNumericHeight("a")
	}
	if character == "E" {
		return getNumericHeight("z")
	}

	b := []byte(character)[0]
	height := int(b - 'a')
	return height
}

func getSingleDigitHeight(height int) int {
	return int(math.Mod(float64(height), 10))
}

func getColor(height int) (int, int, int) {
	if height < 13 {
		adjustedHeight := (float64(height) / 13) * 255
		return 0, int(255 - adjustedHeight), 0
	}

	adjustedHeight := ((float64(height) - 13) / 13) * 255
	return int(adjustedHeight), 0, 0
}

func getHScore(positionToScore *Position, endPosition *Position) float64 {
	// a^2 + b^2 = c^2
	// pythagorean theorem
	posX := float64(positionToScore.x)
	posY := float64(positionToScore.y)
	endX := float64(endPosition.x)
	endY := float64(endPosition.y)
	aSquared := math.Pow(math.Abs(posX-endX), 2)
	bSquared := math.Pow(math.Abs(posY-endY), 2)
	cSquared := aSquared + bSquared
	return math.Sqrt(cSquared)
}

func isPositionInSlice(p *Position, slice []*Position) (*Position, bool) {
	x := p.x
	y := p.y
	for i := 0; i < len(slice); i++ {
		if slice[i].x == x && slice[i].y == y {
			return slice[i], true
		}
	}
	return nil, false
}

func getNeighbors(currentPosition *Position, endPosition *Position, gScore float64, input []string, closedList map[string]*Position) []*Position {
	returnableNeighbors := []*Position{}
	neighbors := []*Position{}
	if currentPosition.y > 0 {
		// up
		neighbors = append(neighbors, &Position{currentPosition.x, currentPosition.y - 1, -1, -1, -1, nil})
	}
	if currentPosition.x > 0 {
		// left
		neighbors = append(neighbors, &Position{currentPosition.x - 1, currentPosition.y, -1, -1, -1, nil})
	}
	if currentPosition.y < len(input)-1 {
		// down
		neighbors = append(neighbors, &Position{currentPosition.x, currentPosition.y + 1, -1, -1, -1, nil})
	}
	if currentPosition.x < len(input[0])-1 {
		// right
		neighbors = append(neighbors, &Position{currentPosition.x + 1, currentPosition.y, -1, -1, -1, nil})
	}

	for _, n := range neighbors {
		n.height = getNumericHeight(string(input[n.y][n.x]))
		n.h = getHScore(n, endPosition)
		n.g = gScore
		// Don't include a neighbor if it's in the closed list - it's already been visited
		_, isInClosedList := closedList[n.getKey()]
		if n.height <= currentPosition.height+1 && !isInClosedList {
			returnableNeighbors = append(returnableNeighbors, n)
		}
	}

	return returnableNeighbors
}

func AStar(startPosition *Position, endPosition *Position, input []string) *Position {
	openList := []*Position{} // The positions that need to be visited
	closedList := map[string]*Position{}
	openList = append(openList, startPosition) // Add the start to the open list
	startPosition.g = 0                        // Set the g score for the start to zero

	for len(openList) > 0 {
		sort.Sort(ByFScore(openList))
		currentPosition := openList[0]
		openList = openList[1:]
		closedList[currentPosition.getKey()] = currentPosition

		log.Printf("Visiting %v, %v h: %v, g: %v", currentPosition.x, currentPosition.y, currentPosition.h, currentPosition.g)
		if currentPosition.x == endPosition.x && currentPosition.y == endPosition.y {
			return currentPosition
		}

		neighbors := getNeighbors(currentPosition, endPosition, currentPosition.g+1, input, closedList)
		for _, neighbor := range neighbors {
			if openListPosition, ok := isPositionInSlice(neighbor, openList); ok {
				// This neighbor is already in the open list
				// This means that we already plan to visit it
				// While sitting in the open list, it already has a g score
				// If the new g score is better, update it and give it a new parent
				// Otherwise, leave it alone and do nothing
				neighborFScore := neighbor.g + neighbor.h
				olpFScore := openListPosition.g + openListPosition.h
				if neighborFScore < olpFScore {
					// Replace the position in the open list
					openListPosition.g = neighbor.g
					openListPosition.h = neighbor.h
					openListPosition.parent = currentPosition
				} else {
					// Leave the position in the open list - it's better than the one we just found
				}
			} else {
				// add the neighbor to the open list
				neighbor.parent = currentPosition
				openList = append(openList, neighbor)
			}
		}
	}

	return nil
}

func main() {
	input := getInputLines()
	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	canvas := svg.New(os.Stdout)
	canvasWidth := len(input[0]) * squareWidth
	canvasHeight := len(input) * squareHeight
	canvas.Start(canvasWidth, canvasHeight)
	_, endPos := drawLandscapeMap(canvas, input)

	minStepCount := int(math.Inf(1))
	for i, line := range input {
		for j, char := range line {
			if char == 'a' || char == 'S' {
				p := &Position{
					x:      j,
					y:      i,
					height: 0,
				}
				result := AStar(p, endPos, input)
				stepCount := drawPathFromEndToStart(result, canvas)
				if stepCount < minStepCount && stepCount != 0 {
					minStepCount = stepCount
				}
			}
		}
	}

	canvas.End()
	log.Printf("Step count: %v", minStepCount)
}

func drawPathFromEndToStart(result *Position, canvas *svg.SVG) int {
	var p *Position
	p = result
	stepCount := 0
	for p != nil {
		if p.parent != nil {
			drawLine(p, p.parent, canvas)
			stepCount++
		}

		p = p.parent
	}
	return stepCount
}
