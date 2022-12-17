package main

import (
	"log"
	"regexp"
	"strconv"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	undo, _ := zap.RedirectStdLogAt(logger, zap.DebugLevel)
	defer undo()

	lines := getInputLines()
	cycleValueMap := map[int]int{}
	cycle := 1
	noopRegex := regexp.MustCompile("^noop$")
	addRegex := regexp.MustCompile("^addx (-?\\d+)$")
	XValue := 1

	for _, instruction := range lines {
		if instruction == "" {
			continue
		}

		log.Printf("Cycle %v, XValue: %v, Instruction: %v", cycle, XValue, instruction)
		cycleValueMap[cycle] = XValue
		if !noopRegex.MatchString(instruction) {
			submatches := addRegex.FindStringSubmatch(instruction)
			if len(submatches) != 2 {
				log.Fatalf("Unable to parse instruction: %v", instruction)
			}
			cycle++
			cycleValueMap[cycle] = XValue

			value, _ := strconv.Atoi(submatches[1])
			XValue = XValue + value
		}

		cycle++
	}

	log.Printf("Final XValue: %v", XValue)
	//20th, 60th, 100th, 140th, 180th, and 220th
	interestingCycles := []int{20, 60, 100, 140, 180, 220}
	runningSum := 0
	for _, ic := range interestingCycles {
		signalStrength := ic * cycleValueMap[ic]
		log.Printf("Cycle %v, XValue: %v, strength: %v", ic, cycleValueMap[ic], signalStrength)
		runningSum += signalStrength
	}

	logger.Sugar().Infow("Sum on interesting signal strengths", "value", runningSum)
}
