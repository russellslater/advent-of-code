package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-05-supply-stacks/crane"
	"github.com/russellslater/advent-of-code/internal/util"
)

func stackInput() [][]string {
	stacks := [][]string{
		{"T", "F", "V", "Z", "C", "W", "S", "Q"},
		{"B", "R", "Q"},
		{"S", "M", "P", "Q", "T", "Z", "B"},
		{"H", "Q", "R", "F", "V", "D"},
		{"P", "T", "S", "B", "D", "L", "G", "J"},
		{"Z", "T", "R", "W"},
		{"J", "R", "F", "S", "N", "M", "Q", "H"},
		{"W", "H", "F", "N", "R"},
		{"B", "P", "R", "Q", "T", "Z", "J"},
	}
	return stacks
}

func main() {
	filename := "./2022/day-05-supply-stacks/input.txt"
	moves := getTransformedInput(filename)

	out := crane.NewCrateMover9000().Operate(stackInput(), moves)
	fmt.Printf("Part One Answer: %v\n", out)

	out = crane.NewCrateMover9001().Operate(stackInput(), moves)
	fmt.Printf("Part Two Answer: %v\n", out)
}

func getTransformedInput(filename string) []crane.Move {
	return crane.BuildMoves(util.LoadInput(filename))
}
