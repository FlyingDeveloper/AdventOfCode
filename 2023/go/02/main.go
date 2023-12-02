package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"rob-hamilton.com/AdventOfCode/internal"
)

const pullRegexString string = `(\d+) (\w+)`
const gameRegexString string = `^Game (\d+): (.+)$`
const sampleInput1 string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

var pullRegex *regexp.Regexp
var gameRegex *regexp.Regexp

type BagPull struct {
	Red   int
	Green int
	Blue  int
}

func NewBagPull(red int, green int, blue int) *BagPull {
	return &BagPull{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
}

func ParseBagPull(input string) *BagPull {
	matches := pullRegex.FindAllStringSubmatch(input, -1)
	bagPull := &BagPull{}
	for _, match := range matches {
		number, _ := strconv.Atoi(match[1])
		color := match[2]
		switch color {
		case "red":
			bagPull.Red = number
		case "green":
			bagPull.Green = number
		case "blue":
			bagPull.Blue = number
		}
	}
	return bagPull
}

type Game struct {
	GameId   int
	BagPulls []*BagPull
}

func NewGame(bagPulls []*BagPull) *Game {
	return &Game{BagPulls: bagPulls}
}

func ParseGame(input string) *Game {
	matches := gameRegex.FindStringSubmatch(input)
	gameId, _ := strconv.Atoi(matches[1])
	bagPullsString := matches[2]
	bagPulls := []*BagPull{}
	for _, bagPullString := range strings.Split(bagPullsString, ";") {
		bagPull := ParseBagPull(bagPullString)
		bagPulls = append(bagPulls, bagPull)
	}

	return &Game{
		GameId:   gameId,
		BagPulls: bagPulls,
	}
}

func (bp *BagPull) AsString() string {
	return fmt.Sprintf("Red: %v, Green: %v, Blue: %v\n", bp.Red, bp.Green, bp.Blue)
}

func (bp *BagPull) Power() int {
	return bp.Red * bp.Green * bp.Blue
}

func (g *Game) GetMinCubes() (red int, green int, blue int) {
	red = 0
	green = 0
	blue = 0

	for _, bp := range g.BagPulls {
		if bp.Red != 0 && bp.Red > red {
			red = bp.Red
		}
		if bp.Green != 0 && bp.Green > green {
			green = bp.Green
		}
		if bp.Blue != 0 && bp.Blue > blue {
			blue = bp.Blue
		}
	}

	return
}

func (g *Game) IsPossible(red int, green int, blue int) bool {
	for _, bp := range g.BagPulls {
		if red < bp.Red || green < bp.Green || blue < bp.Blue {
			return false
		}
	}

	return true
}

func part1() {
	//input := strings.Split(sampleInput1, "\n")
	input, _ := internal.GetInputAsStringArray("./02/input")
	games := []*Game{}
	for _, line := range input {
		if line == "" {
			continue
		}

		game := ParseGame(line)
		games = append(games, game)
	}

	possibleGames := []*Game{}
	for _, g := range games {
		if g.IsPossible(12, 13, 14) {
			possibleGames = append(possibleGames, g)
		}
	}

	runningTotal := 0
	for _, g := range possibleGames {
		runningTotal += g.GameId
	}

	fmt.Println(runningTotal)
}

func part2() {
	//input := strings.Split(sampleInput1, "\n")
	input, _ := internal.GetInputAsStringArray("./02/input")

	runningTotal := 0
	for _, line := range input {
		if line == "" {
			continue
		}
		g := ParseGame(line)
		red, green, blue := g.GetMinCubes()
		power := red * green * blue
		runningTotal += power
	}

	fmt.Println(runningTotal)
}

func main() {
	pullRegex = regexp.MustCompile(pullRegexString)
	gameRegex = regexp.MustCompile(gameRegexString)

	part1()
	part2()
}
