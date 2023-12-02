package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2023/day-02-cube-conundrum/game"
	"github.com/russellslater/advent-of-code/internal/util"
)

const (
	blueMax  = 14
	redMax   = 12
	greenMax = 13
)

func main() {
	filename := "./2023/day-02-cube-conundrum/input.txt"
	lines := util.LoadInput(filename)

	games := parseGames(lines)

	sum := 0
	sumPower := 0

	for _, game := range games {
		if game.IsValid(blueMax, redMax, greenMax) {
			sum += game.ID
		}
		sumPower += game.PowerMinCubeSets()
	}

	fmt.Printf("Part One Answer: %d\n", sum)
	fmt.Printf("Part Two Answer: %d\n", sumPower)
}

func parseGames(lines []string) []*game.Game {
	games := []*game.Game{}
	for _, line := range lines {
		games = append(games, game.NewGame(line))
	}
	return games
}
