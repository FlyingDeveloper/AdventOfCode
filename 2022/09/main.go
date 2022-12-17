package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"

	"go.uber.org/zap"
)

type Knot struct {
	positionX int
	positionY int
}

type Instruction struct {
	direction string
	distance  int
}

func ParseInstruction(input string) Instruction {
	r := regexp.MustCompile("^([LUDR]{1})\\s(\\d+)$")
	submatches := r.FindStringSubmatch(input)
	if len(submatches) != 3 {
		log.Fatalf("Unable to parse input string %v\n", input)
	}
	d, _ := strconv.Atoi(submatches[2])
	return Instruction{
		direction: submatches[1],
		distance:  d,
	}
}

func (k *Knot) IsTouching(otherKnot *Knot) bool {
	if k.positionX == otherKnot.positionX {
		// Same column
		if math.Abs(float64(k.positionY-otherKnot.positionY)) <= 1 {
			return true
		}
		return false
	}
	if k.positionY == otherKnot.positionY {
		if math.Abs(float64(k.positionX-otherKnot.positionX)) <= 1 {
			return true
		}
		return false
	}

	//Check diagonals
	if math.Abs(float64(k.positionX-otherKnot.positionX)) <= 1 &&
		math.Abs(float64(k.positionY-otherKnot.positionY)) <= 1 {
		return true
	}
	return false
}

func getVisitedKey(knot *Knot) string {
	return fmt.Sprintf("%v,%v", knot.positionX, knot.positionY)
}

func moveTailTowardHead(head *Knot, tail *Knot, visited map[string]interface{}) {
	// Already touching (do nothing)
	if head.IsTouching(tail) {
		return
	}

	// Same Row
	if head.positionY == tail.positionY {
		if head.positionX > tail.positionX {
			for !head.IsTouching(tail) {
				log.Printf("Moving tail to the right to %v, %v\n", tail.positionX+1, tail.positionY)
				tail.positionX = tail.positionX + 1
			}
			visited[getVisitedKey(tail)] = true
			return
		}
		for !head.IsTouching(tail) {
			log.Printf("Moving tail to the left to %v, %v\n", tail.positionX-1, tail.positionY)
			tail.positionX = tail.positionX - 1
		}
		visited[getVisitedKey(tail)] = true

		return
	}

	// Same Col
	if head.positionX == tail.positionX {
		if head.positionY > tail.positionY {
			for !head.IsTouching(tail) {
				log.Printf("Moving tail up to %v, %v\n", tail.positionX, tail.positionY+1)
				tail.positionY = tail.positionY + 1
			}
			visited[getVisitedKey(tail)] = true
			return
		}
		for !head.IsTouching(tail) {
			log.Printf("Moving tail down to %v, %v\n", tail.positionX, tail.positionY-1)
			tail.positionY = tail.positionY - 1
		}
		visited[getVisitedKey(tail)] = true

		return
	}

	// Gonna have to move diagonally
	if head.positionX > tail.positionX {
		log.Printf("Diagonal: moving tail right to %v, %v\n", tail.positionX+1, tail.positionY)
		tail.positionX = tail.positionX + 1
	} else if head.positionX < tail.positionX {
		log.Printf("Diagonal: moving tail left to %v, %v\n", tail.positionX-1, tail.positionY)
		tail.positionX = tail.positionX - 1
	}

	if head.positionY > tail.positionY {
		log.Printf("Diagnoal: moving tail up to %v, %v\n", tail.positionX, tail.positionY+1)
		tail.positionY = tail.positionY + 1
	} else if head.positionY < tail.positionY {
		log.Printf("Diagonal: moving tail down to %v, %v\n", tail.positionX, tail.positionY-1)
		tail.positionY = tail.positionY - 1
	}
	visited[getVisitedKey(tail)] = true

}

func main() {
	z, _ := zap.NewProduction()
	defer z.Sync()
	undo, _ := zap.RedirectStdLogAt(z, zap.DebugLevel)
	defer undo()

	visited := map[string]interface{}{}
	head := Knot{positionX: 0, positionY: 0}
	tail := Knot{positionX: 0, positionY: 0}
	visited[getVisitedKey(&tail)] = true

	instructions := getInputLines()

	for _, i := range instructions {
		if i == "" {
			continue
		}

		parsed := ParseInstruction(i)
		log.Print(parsed)
		for j := 0; j < parsed.distance; j++ {
			switch parsed.direction {
			case "U":
				head.positionY = head.positionY + 1
			case "R":
				head.positionX = head.positionX + 1
			case "D":
				head.positionY = head.positionY - 1
			case "L":
				head.positionX = head.positionX - 1
			}
			log.Print(head)
			log.Print(head.IsTouching(&tail))
			if !head.IsTouching(&tail) {
				moveTailTowardHead(&head, &tail, visited)
			}
		}
	}

	z.Sugar().Infow("Number of visited locations", "locations", len(visited))
}
