package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-22-monkey-map/input.txt"
	board, moves := getTransformedInput(filename)

	password := calcFinalPassword(board, moves)
	fmt.Printf("Part One Answer: %v\n", password)
}

type Move struct {
	Turn Turn
	Step int
}

type Turn int

const (
	None Turn = iota
	CW
	CCW
)

func getTransformedInput(filename string) ([][]rune, []Move) {
	board := [][]rune{}
	lines := util.LoadInput(filename)
	for _, line := range lines[:len(lines)-2] {
		// TODO: All rows could be the same length for simplicity
		board = append(board, []rune(line))
	}
	moves := parseMoves(lines[len(lines)-1])
	return board, moves
}

func parseMoves(line string) []Move {
	moves := []Move{}
	numChars := ""

	for _, c := range line {
		if c == 'R' || c == 'L' {
			if numChars != "" {
				moves = append(moves, Move{Turn: None, Step: util.MustAtoi(numChars)})
				numChars = ""
			}
			if c == 'R' {
				moves = append(moves, Move{Turn: CW})
			} else {
				moves = append(moves, Move{Turn: CCW})
			}
		} else {
			numChars += string(c)
		}
	}

	if numChars != "" {
		moves = append(moves, Move{Turn: None, Step: util.MustAtoi(numChars)})
		numChars = ""
	}

	return moves
}

func calcFinalPassword(board [][]rune, moves []Move) int {
	fmt.Println(board)
	fmt.Println(moves)

	return 0
}
