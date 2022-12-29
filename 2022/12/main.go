package main

import (
	"container/list"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"

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

func (p *Position) GetPositionKey() string {
	return fmt.Sprintf("%v/%v", p.x, p.y)
}

type ByHeuristic []*Position

func (a ByHeuristic) Len() int           { return len(a) }
func (a ByHeuristic) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHeuristic) Less(i, j int) bool { return a[i].h < a[j].h }

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

func buildSvg() {
	width := 500
	height := 500
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	canvas.Circle(width/2, height/2, 100)
	canvas.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas.End()
}

func getColor(height int) (int, int, int) {
	if height < 13 {
		adjustedHeight := (float64(height) / 13) * 255
		return 0, int(255 - adjustedHeight), 0
	}

	adjustedHeight := ((float64(height) - 13) / 13) * 255
	return int(adjustedHeight), 0, 0
}

func heuristic(currentPosition *Position, endPosition *Position) float64 {
	// manhattan distance
	//return math.Abs(float64(currentPosition.x)-float64(currentPosition.y)) + math.Abs(float64(currentPosition.x)-float64(currentPosition.y))
	// a^2 + b^2 = c^2
	posX := float64(currentPosition.x)
	posY := float64(currentPosition.y)
	endX := float64(endPosition.x)
	endY := float64(endPosition.y)
	aSquared := math.Pow(math.Abs(posX-endX), 2)
	bSquared := math.Pow(math.Abs(posY-endY), 2)
	cSquared := aSquared + bSquared
	return math.Sqrt(cSquared)
}

func getHColor(currentPosition *Position, endPosition *Position) (int, int, int) {
	h := heuristic(currentPosition, endPosition)
	adjustedHeuristic := (h / 25) * 255
	return int(adjustedHeuristic), 127, 127
}

func drawLine(endA *Position, endB *Position, canvas *svg.SVG) {
	startX := endA.x*squareWidth + (squareWidth / 2)
	startY := endA.y*squareHeight + (squareHeight / 2)
	endX := endB.x*squareWidth + (squareWidth / 2)
	endY := endB.y*squareHeight + (squareHeight / 2)
	canvas.Line(startX, startY, endX, endY, "stroke: black; outline-color:white;outline-style:solid")
}

func PrintPath(path *list.List) {
	output := []string{}
	for e := path.Front(); e != nil; e = e.Next() {
		ePosition := e.Value.(*Position)
		output = append(output, fmt.Sprintf("%v, %v", ePosition.x, ePosition.y))
	}

	log.Printf(strings.Join(output, " -- "))
}

var fileCount = 0
var canvases []*svg.SVG

func drawPathToFile(path list.List, endPosition *Position, width int, height int, input []string) {
	fileCount++
	fs, err := os.Create(fmt.Sprintf("/Users/rob/tmp/%v.svg", fileCount))
	if err != nil {
		log.Printf("Unable to create file %v", err.Error())
		return
	}
	canvas := svg.New(fs)
	canvas.Start(width, height)

	for il, line := range input {
		chars := strings.Split(line, "")
		for ic, char := range chars {
			var red, green, blue int
			if char == "S" {
				red, green, blue = 0, 0, 255
			} else if char == "E" {
				red, green, blue = 0, 255, 0
			} else {
				red, green, blue = getColor(getNumericHeight(char))
			}

			style := fmt.Sprintf("fill: rgb(%v, %v, %v); xy: '%v--%v';", red, green, blue, ic, il)
			canvas.Rect(0*ic*squareWidth, 0*il*squareHeight, squareWidth, squareHeight, style)
		}
	}

	var prev *list.Element
	for e := path.Front(); e != nil; e = e.Next() {
		if prev == nil {
			prev = e
		} else {
			drawLine(prev.Value.(*Position), e.Value.(*Position), canvas)
			prev = e
		}
	}

	canvas.Circle(
		endPosition.x*squareWidth+(squareWidth/2),
		endPosition.y*squareHeight+(squareHeight/2),
		squareWidth*2,
		"stroke: blue; fill-opacity: 0")

	canvas.End()
	fs.Close()
}

var navigateCount = 0
var unusablePositions = map[string]bool{}

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

func getNeighbors(currentPosition *Position, endPosition *Position, gScore float64, input []string, closedList []*Position) []*Position {
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
		n.h = heuristic(n, endPosition)
		n.g = gScore
		_, isInClosedList := isPositionInSlice(n, closedList)
		if n.height <= currentPosition.height+1 && !isInClosedList {
			returnableNeighbors = append(returnableNeighbors, n)
		}
	}
	return returnableNeighbors
}

func navigateBetter(startPosition *Position, endPosition *Position, input []string) *Position {
	openList := []*Position{}
	closedList := []*Position{}
	openList = append(openList, startPosition)
	startPosition.g = 0

	for len(openList) > 0 {
		sort.Sort(ByFScore(openList))
		currentPosition := openList[0]
		openList = openList[1:]
		closedList = append(closedList, currentPosition)

		log.Printf("Visiting %v, %v", currentPosition.x, currentPosition.y)
		if currentPosition.x == endPosition.x && currentPosition.y == endPosition.y {
			return currentPosition
		}

		neighbors := getNeighbors(currentPosition, endPosition, currentPosition.g+1, input, closedList)
		for _, n := range neighbors {
			if p, ok := isPositionInSlice(n, openList); ok {
				nf := n.g + n.h
				pf := p.g + p.h
				if nf < pf {
					// Replace the position in the open list
					p.g = n.g
					p.h = n.h
					n.parent = currentPosition
				} else {
					// Leave the position in the open list - it's better than the one we just found
				}
			} else {
				// add the neighbor to the open list
				n.parent = currentPosition
				openList = append(openList, n)
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
	startPos := Position{-1, -1, int(math.Inf(-1)), math.Inf(1), 0, nil}
	endPos := Position{-1, -1, int(math.Inf(1)), 0, 0, nil}
	for il, line := range input {
		chars := strings.Split(line, "")
		for ic, char := range chars {
			var red, green, blue int
			if char == "S" {
				startPos = Position{ic, il, getNumericHeight("a"), -1, 0, nil}
				red, green, blue = 0, 0, 255
			} else if char == "E" {
				endPos = Position{ic, il, getNumericHeight("z"), 0, 0, nil}
				red, green, blue = 0, 255, 0
			} else {
				red, green, blue = getColor(getNumericHeight(char))
			}
			style := fmt.Sprintf("fill: rgb(%v, %v, %v); xy: '%v--%v'", red, green, blue, ic, il)
			canvas.Rect(ic*squareWidth, il*squareHeight, squareWidth, squareHeight, style)
		}
	}

	canvas.Circle(startPos.x*squareWidth+(squareWidth/2), startPos.y*squareHeight+(squareHeight/2), squareWidth*2, "stroke: yellow; fill-opacity: 0")
	canvas.Circle(endPos.x*squareWidth+(squareWidth/2), endPos.y*squareHeight+(squareHeight/2), squareWidth*2, "stroke: yellow; fill-opacity: 0")

	currentPosition := Position{
		startPos.x,
		startPos.y,
		getNumericHeight(string(input[startPos.y][startPos.x])),
		-1,
		0,
		nil,
	}
	currentPosition.h = heuristic(&currentPosition, &endPos)

	result := navigateBetter(&currentPosition, &endPos, input)
	fmt.Println(result)

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

	canvas.End()
	log.Printf("Step count: %v", stepCount)
	return

	/*
	   path := list.New()
	   path.PushBack(&currentPosition)
	   log.Printf("Result: %v", navigate(&currentPosition, &endPos, path, input, map[string]bool{}))

	   var prev *list.Element

	   	for e := path.Front(); e != nil; e = e.Next() {
	   		if prev == nil {
	   			prev = e
	   		} else {
	   			drawLine(prev.Value.(*Position), e.Value.(*Position), canvas)
	   			prev = e
	   		}
	   	}

	   canvas.End()
	*/
}
