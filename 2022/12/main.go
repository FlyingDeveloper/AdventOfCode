package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	svg "github.com/ajstarks/svgo"
)

const (
	squareWidth  = 30
	squareHeight = 30
)

func getNumericHeight(character string) int {
	if len(character) != 1 {
		log.Fatalf("Unable to determine height of %v", character)
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
	adjustedHeight := (float64(height) / 26) * 255
	return int(adjustedHeight), 127, 127
}

func main() {
	input := getSampleInput()
	canvas := svg.New(os.Stdout)
	canvasWidth := len(input[0]) * squareWidth
	canvasHeight := len(input) * squareHeight
	canvas.Start(canvasWidth, canvasHeight)
	startPosX := -1
	startPosY := -1
	endPosX := -1
	endPosY := -1
	for il, line := range input {
		chars := strings.Split(line, "")
		for ic, char := range chars {
			var red, green, blue int
			if char == "S" {
				startPosX = ic
				startPosY = il
				red, green, blue = 0, 0, 255
			} else if char == "E" {
				endPosX = ic
				endPosY = il
				red, green, blue = 0, 255, 0
			} else {
				red, green, blue = getColor(getNumericHeight(char))
			}
			style := fmt.Sprintf("fill: rgb(%v, %v, %v); alt-text: %v", red, green, blue, char)
			canvas.Rect(ic*squareWidth, il*squareHeight, squareWidth, squareHeight, style)
		}
	}
	canvas.Circle(startPosX*squareWidth+(squareWidth/2), startPosY*squareHeight+(squareHeight/2), squareWidth*2, "stroke: yellow; fill-opacity: 0")
	canvas.Circle(endPosX*squareWidth+(squareWidth/2), endPosY*squareHeight+(squareHeight/2), squareWidth*2, "stroke: yellow; fill-opacity: 0")

	canvas.End()
}
