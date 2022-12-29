package main

import (
	"fmt"
	"math"
	"strings"

	svg "github.com/ajstarks/svgo"
)

func drawLine(endA *Position, endB *Position, canvas *svg.SVG) {
	startX := endA.x*squareWidth + (squareWidth / 2)
	startY := endA.y*squareHeight + (squareHeight / 2)
	endX := endB.x*squareWidth + (squareWidth / 2)
	endY := endB.y*squareHeight + (squareHeight / 2)
	canvas.Line(startX, startY, endX, endY, "stroke: black; outline-color:white;outline-style:solid")
}

func drawLandscapeMap(canvas *svg.SVG, input []string) (startPos *Position, endPos *Position) {
	for il, line := range input {
		chars := strings.Split(line, "")
		for ic, char := range chars {
			var red, green, blue int
			if char == "S" {
				startPos = &Position{
					x:      ic,
					y:      il,
					height: getNumericHeight("a"),
				}

				red, green, blue = 0, 0, 255
			} else if char == "E" {
				endPos = &Position{
					x:      ic,
					y:      il,
					height: getNumericHeight("z"),
				}
				red, green, blue = 0, 255, 0
			} else {
				red, green, blue = getColor(getNumericHeight(char))
			}
			style := fmt.Sprintf("fill: rgb(%v, %v, %v); xy: '%v--%v'", red, green, blue, ic, il)
			canvas.Rect(ic*squareWidth, il*squareHeight, squareWidth, squareHeight, style)
		}
	}

	canvas.Circle(
		startPos.x*squareWidth+(squareWidth/2),
		startPos.y*squareHeight+(squareHeight/2),
		squareWidth*2, "stroke: yellow; fill-opacity: 0",
	)
	canvas.Circle(
		endPos.x*squareWidth+(squareWidth/2),
		endPos.y*squareHeight+(squareHeight/2),
		squareWidth*2, "stroke: yellow; fill-opacity: 0",
	)

	return
}

func drawOpenClosedMap(openList []*Position, closedList map[string]*Position, canvas *svg.SVG, input []string) {
	drawLandscapeMap(canvas, input)

	for _, op := range openList {
		canvas.Circle(
			op.x*squareWidth+(squareWidth/2),
			op.y*squareHeight+(squareHeight/2),
			int(math.Min(squareWidth, squareHeight)/2),
			"stroke: white; stroke-width: 2px; fill-opacity: 0",
		)
	}

	for _, cp := range closedList {
		canvas.Circle(
			cp.x*squareWidth+(squareWidth/2),
			cp.y*squareHeight+(squareHeight/2),
			int(math.Min(squareWidth, squareHeight)/2),
			"stroke: white; stroke-width: 2px; fill: black; fill-opacity: 100",
		)
	}
}
