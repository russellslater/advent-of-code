package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/twod"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2016/day-01-no-time-for-a-taxicab/input.txt"
	moves := getTransformedInput(filename)

	dist := calcDistanceInBlocks(moves)
	fmt.Printf("Part One Answer: %d\n", dist)

	dist = calcDistanceToFirstLocationVisitedTwice(moves)
	fmt.Printf("Part Two Answer: %d\n", dist)
}

type Move struct {
	Turn  twod.RelativeDirection
	Steps int
}

func getTransformedInput(filename string) []Move {
	moves := []Move{}
	for _, mv := range strings.Split(util.LoadInput(filename)[0], ", ") {
		move := Move{}
		switch mv[0] {
		case 'L':
			move.Turn = twod.Left
		case 'R':
			move.Turn = twod.Right
		}
		move.Steps = util.MustAtoi(mv[1:])
		moves = append(moves, move)
	}
	return moves
}

func calcDistanceInBlocks(moves []Move) int {
	x, y := 0, 0
	dir := twod.North

	for _, mv := range moves {
		dir = dir.Turn(mv.Turn)
		switch dir {
		case twod.North:
			y += mv.Steps
		case twod.South:
			y -= mv.Steps
		case twod.West:
			x -= mv.Steps
		case twod.East:
			x += mv.Steps
		}
	}

	return util.ManhattanDistance(0, 0, x, y)
}

func calcDistanceToFirstLocationVisitedTwice(moves []Move) int {
	x, y := 0, 0
	dir := twod.North
	visited := map[twod.Position]bool{}

	for _, mv := range moves {
		dir = dir.Turn(mv.Turn)
		dx, dy := 0, 0
		switch dir {
		case twod.North:
			dy = 1
		case twod.South:
			dy = -1
		case twod.West:
			dx = -1
		case twod.East:
			dx = 1
		}

		for i := 0; i < mv.Steps; i++ {
			x += dx
			y += dy
			if visited[twod.Position{X: x, Y: y}] {
				return util.ManhattanDistance(0, 0, x, y)
			}
			visited[twod.Position{X: x, Y: y}] = true
		}
	}

	return -1
}
