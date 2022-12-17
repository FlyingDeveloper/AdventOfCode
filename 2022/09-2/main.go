package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type Knot struct {
	positionX int
	positionY int
	follower  *Knot
	name      string
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

func moveFollowerTowardLeader(leader *Knot, follower *Knot) {
	// Already touching (do nothing)
	if leader.IsTouching(follower) {
		return
	}

	// Same Row
	if leader.positionY == follower.positionY {
		if leader.positionX > follower.positionX {
			for !leader.IsTouching(follower) {
				log.Printf("Moving tail to the right to %v, %v\n", follower.positionX+1, follower.positionY)
				follower.positionX = follower.positionX + 1
			}
			return
		}
		for !leader.IsTouching(follower) {
			log.Printf("Moving tail to the left to %v, %v\n", follower.positionX-1, follower.positionY)
			follower.positionX = follower.positionX - 1
		}

		return
	}

	// Same Col
	if leader.positionX == follower.positionX {
		if leader.positionY > follower.positionY {
			for !leader.IsTouching(follower) {
				log.Printf("Moving tail up to %v, %v\n", follower.positionX, follower.positionY+1)
				follower.positionY = follower.positionY + 1
			}
			return
		}
		for !leader.IsTouching(follower) {
			log.Printf("Moving tail down to %v, %v\n", follower.positionX, follower.positionY-1)
			follower.positionY = follower.positionY - 1
		}

		return
	}

	// Gonna have to move diagonally
	if leader.positionX > follower.positionX {
		log.Printf("Diagonal: moving tail right to %v, %v\n", follower.positionX+1, follower.positionY)
		follower.positionX = follower.positionX + 1
	} else if leader.positionX < follower.positionX {
		log.Printf("Diagonal: moving tail left to %v, %v\n", follower.positionX-1, follower.positionY)
		follower.positionX = follower.positionX - 1
	}

	if leader.positionY > follower.positionY {
		log.Printf("Diagnoal: moving tail up to %v, %v\n", follower.positionX, follower.positionY+1)
		follower.positionY = follower.positionY + 1
	} else if leader.positionY < follower.positionY {
		log.Printf("Diagonal: moving tail down to %v, %v\n", follower.positionX, follower.positionY-1)
		follower.positionY = follower.positionY - 1
	}
}

func buildTrain(length int) (head *Knot) {
	head = &Knot{
		positionX: 0,
		positionY: 0,
	}
	current := head
	for i := 1; i < length; i++ {
		current.follower = &Knot{
			positionX: 0,
			positionY: 0,
		}
		current = current.follower
	}

	return
}

func visualize(head *Knot) {
	maxY := math.Inf(-1)
	minY := math.Inf(1)
	maxX := math.Inf(-1)
	minX := math.Inf(1)

	rows := map[int]map[int]string{}

	for current := head; current != nil; {
		if current.positionX < int(minX) {
			minX = float64(current.positionX)
		}
		if current.positionX > int(maxX) {
			maxX = float64(current.positionX)
		}
		if current.positionY < int(minY) {
			minY = float64(current.positionY)
		}
		if current.positionY > int(maxY) {
			maxY = float64(current.positionY)
		}

		if _, ok := rows[current.positionY]; !ok {
			rows[current.positionY] = map[int]string{}
		}
		rows[current.positionY][current.positionX] = current.name

		current = current.follower
	}
	minX = -30
	maxX = 30
	minY = -30
	maxY = 30

	for r := maxY; r >= minY; r-- {
		row, okRow := rows[int(r)]
		if !okRow {
			fmt.Println(". . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . . ")
		} else {
			rowOutput := []string{}
			for c := minX; c <= maxX; c++ {
				currentVal, okCol := row[int(c)]
				if !okCol {
					rowOutput = append(rowOutput, ".")
				} else {
					rowOutput = append(rowOutput, fmt.Sprintf("%v", currentVal))
				}
			}
			fmt.Println(strings.Join(rowOutput, " "))
		}
	}
	fmt.Println("################################")
}

func main() {
	z, _ := zap.NewProduction()
	defer z.Sync()
	undo, _ := zap.RedirectStdLogAt(z, zap.DebugLevel)
	defer undo()

	visited := map[string]interface{}{}

	var head *Knot
	var previous *Knot
	for i := 0; i < 10; i++ {
		current := Knot{positionX: 0, positionY: 0, name: strconv.Itoa(i)}
		if previous != nil {
			previous.follower = &current
		} else {
			head = &current
		}
		previous = &current
	}

	visited[getVisitedKey(head)] = true

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

			for current := head; current != nil; {
				log.Printf("----%v\n", current.name)
				if current.follower != nil {
					moveFollowerTowardLeader(current, current.follower)
				} else {
					visited[getVisitedKey(current)] = true
				}

				current = current.follower
			}
		}
	}

	z.Sugar().Infow("Number of visited locations", "locations", len(visited))
}
