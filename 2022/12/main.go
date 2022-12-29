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
}

func (p *Position) GetPositionKey() string {
	return fmt.Sprintf("%v/%v", p.x, p.y)
}

type ByHeuristic []*Position

func (a ByHeuristic) Len() int           { return len(a) }
func (a ByHeuristic) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByHeuristic) Less(i, j int) bool { return a[i].h < a[j].h }

func getNumericHeight(character string) int {
	if len(character) != 1 {
		log.Fatalf("Unable to determine height of %v", character)
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
	adjustedHeight := (float64(height) / 6) * 255
	return int(adjustedHeight), 127, 127
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
	canvas.Line(startX, startY, endX, endY, "stroke: brown")
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
			style := fmt.Sprintf("fill: rgb(%v, %v, %v); alt-text: %v", red, green, blue, char)
			canvas.Rect(ic*squareWidth, il*squareHeight, squareWidth, squareHeight, style)
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

func navigate(currentPosition *Position, endPosition *Position, path *list.List, input []string, currentPath map[string]bool) bool {
	// Find candidates to move to
	navigateCount++
	if math.Mod(float64(navigateCount), 10000) == 0 {
		drawPathToFile(*path, endPosition, len(input[0])*squareWidth, len(input)*squareHeight, input)
	}
	if path.Len() > 4000 {
		log.Print("Bailing out at 4000 runs")
		return false
	}
	//PrintPath(path)

	candidates := []*Position{}
	if currentPosition.x > 1 {
		// left
		leftPos := Position{currentPosition.x - 1, currentPosition.y, -1, -1}
		leftPos.height = getNumericHeight(string(input[leftPos.y][leftPos.x]))
		leftPos.h = heuristic(&leftPos, endPosition)
		_, isInPath := currentPath[leftPos.GetPositionKey()]
		_, isUnusuable := unusablePositions[leftPos.GetPositionKey()]
		if leftPos.height <= currentPosition.height+1 && !isInPath && !isUnusuable {
			candidates = append(candidates, &leftPos)
		}
	}
	if currentPosition.y > 0 {
		// up
		upPos := Position{currentPosition.x, currentPosition.y - 1, -1, -1}
		upPos.height = getNumericHeight(string(input[upPos.y][upPos.x]))
		upPos.h = heuristic(&upPos, endPosition)
		_, isInPath := currentPath[upPos.GetPositionKey()]
		_, isUnusable := unusablePositions[upPos.GetPositionKey()]
		if upPos.height <= currentPosition.height+1 && !isInPath && !isUnusable {
			candidates = append(candidates, &upPos)
		}
	}
	if currentPosition.x < len(input[0])-1 {
		// right
		rightPos := Position{currentPosition.x + 1, currentPosition.y, -1, -1}
		rightPos.height = getNumericHeight(string(input[rightPos.y][rightPos.x]))
		rightPos.h = heuristic(&rightPos, endPosition)
		isInPath := currentPath[rightPos.GetPositionKey()]
		_, isUnusable := unusablePositions[rightPos.GetPositionKey()]
		if rightPos.height <= currentPosition.height+1 && !isInPath && !isUnusable {
			candidates = append(candidates, &rightPos)
		}
	}
	if currentPosition.y < len(input)-2 {
		// down
		downPos := Position{currentPosition.x, currentPosition.y + 1, -1, -1}
		downPos.height = getNumericHeight(string(input[downPos.y][downPos.x]))
		downPos.h = heuristic(&downPos, endPosition)
		_, isInPath := currentPath[downPos.GetPositionKey()]
		_, isUnusable := unusablePositions[downPos.GetPositionKey()]
		if downPos.height <= currentPosition.height+1 && !isInPath && !isUnusable {
			candidates = append(candidates, &downPos)
		}
	}

	sort.Sort(ByHeuristic(candidates))

	for _, currentCandidate := range candidates {
		_, alreadyInPath := currentPath[currentCandidate.GetPositionKey()]
		/*		for e := path.Back(); e != nil; e = e.Prev() {
				//for e := path.Front(); e != nil; e = e.Next() {
				ePosition := e.Value.(*Position)
				if ePosition.x == currentCandidate.x && ePosition.y == currentCandidate.y {
					alreadyInPath = true
					break
				}
			}*/
		if alreadyInPath {
			//log.Printf("Not visiting %v, %v because it's already in the current path", currentCandidate.x, currentCandidate.y)
			continue
		}

		//log.Printf("current position: %v, %v, Visiting %v, %v", currentPosition.x, currentPosition.y, currentCandidate.x, currentCandidate.y)
		path.PushBack(currentCandidate)
		currentPath[currentCandidate.GetPositionKey()] = true
		if input[currentCandidate.y][currentCandidate.x] == 'E' {
			return true
		}

		result := navigate(currentCandidate, endPosition, path, input, currentPath)
		if result == true {
			return true
		}
		path.Remove(path.Back())
		delete(currentPath, currentCandidate.GetPositionKey())
	}

	unusablePositions[currentPosition.GetPositionKey()] = true
	return false
}

func main() {
	input := getInputLines()
	canvas := svg.New(os.Stdout)
	canvasWidth := len(input[0]) * squareWidth
	canvasHeight := len(input) * squareHeight
	canvas.Start(canvasWidth, canvasHeight)
	startPos := Position{-1, -1, int(math.Inf(-1)), math.Inf(1)}
	endPos := Position{-1, -1, int(math.Inf(1)), 0}
	for il, line := range input {
		chars := strings.Split(line, "")
		for ic, char := range chars {
			var red, green, blue int
			if char == "S" {
				startPos = Position{ic, il, getNumericHeight("a"), -1}
				red, green, blue = 0, 0, 255
			} else if char == "E" {
				endPos = Position{ic, il, getNumericHeight("z"), 0}
				red, green, blue = 0, 255, 0
			} else {
				red, green, blue = getColor(getNumericHeight(char))
			}
			style := fmt.Sprintf("fill: rgb(%v, %v, %v); alt-text: %v", red, green, blue, char)
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
	}
	currentPosition.h = heuristic(&currentPosition, &endPos)

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
}
