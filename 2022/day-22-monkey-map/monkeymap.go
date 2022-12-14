package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-22-monkey-map/jungle"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-22-monkey-map/input.txt"
	board, moves := getTransformedInput(filename)

	twoDimensionalGame := jungle.New2DGame(board, moves)
	password := twoDimensionalGame.CalcFinalPassword()
	fmt.Printf("Part One Answer: %v\n", password)

	// Note: Did not attempt a general purpose solution for Part Two
	// that would work for any input. There are multiple configurations
	// for 6 squares to fold into a cube.

	threeDimensionalGame := jungle.New3DGame(board, moves)
	password = threeDimensionalGame.CalcFinalPassword()
	fmt.Printf("Part Two Answer: %v\n", password)
}

func getTransformedInput(filename string) ([][]rune, []jungle.Move) {
	board := [][]rune{}
	lines := util.LoadInput(filename)
	maxRowLength := 0

	for _, line := range lines[:len(lines)-2] {
		if len(line) > maxRowLength {
			maxRowLength = len(line)
		}
		board = append(board, []rune(line))
	}

	// Resize rows to be the same length and replace spaces with 0s
	for i := range board {
		for len(board[i]) < maxRowLength {
			board[i] = append(board[i], 0)
		}
		for j, c := range board[i] {
			if c == ' ' {
				board[i][j] = 0
			}
		}
	}

	return board, parseMoves(lines[len(lines)-1])
}

func parseMoves(line string) []jungle.Move {
	moves := []jungle.Move{}
	numChars := ""

	for _, c := range line {
		if c == 'R' || c == 'L' {
			if numChars != "" {
				moves = append(moves, jungle.Move{Steps: util.MustAtoi(numChars)})
				numChars = ""
			}
			if c == 'R' {
				moves = append(moves, jungle.ClockwiseTurn)
			} else {
				moves = append(moves, jungle.CounterclockwiseTurn)
			}
		} else {
			numChars += string(c)
		}
	}

	if numChars != "" {
		moves = append(moves, jungle.Move{Steps: util.MustAtoi(numChars)})
		numChars = ""
	}

	return moves
}
