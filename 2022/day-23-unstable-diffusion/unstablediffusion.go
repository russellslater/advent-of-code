package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-23-unstable-diffusion/grove"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-23-unstable-diffusion/input.txt"
	grid := getTransformedInput(filename)

	g := grove.NewGame(grid)

	emptyTileCount := g.SolveForEmptyTiles(10)
	fmt.Printf("Part One Answer: %v\n", emptyTileCount)

	rounds := g.SolveForFirstRoundNoMovement()
	fmt.Printf("Part Two Answer: %v\n", rounds)
}

func getTransformedInput(filename string) [][]rune {
	grid := [][]rune{}
	for _, line := range util.LoadInput(filename) {
		grid = append(grid, []rune(line))
	}
	return grid
}
