package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func IsWinningBoard(markedNums [][]bool) bool {
	// check rows
	for i := 0; i < len(markedNums); i++ {
		markedCount := 0
		for j := 0; j < len(markedNums[i]); j++ {
			if markedNums[i][j] {
				markedCount++
			}
		}

		if markedCount == len(markedNums[i]) {
			return true
		}
	}

	// check columns
	for j := 0; j < len(markedNums[0]); j++ {
		markedCount := 0
		for i := 0; i < len(markedNums); i++ {
			if markedNums[i][j] {
				markedCount++
			}
		}

		if markedCount == len(markedNums) {
			return true
		}
	}

	return false
}

func createMarkedNums(board [][]int) [][]bool {
	markedNums := make([][]bool, len(board), len(board[0]))
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			markedNums[i] = append(markedNums[i], false)
		}
	}
	return markedNums
}

func ScoreBoard(board [][]int, markedNums [][]bool) (score int) {
	for i := 0; i < len(markedNums); i++ {
		for j := 0; j < len(markedNums[i]); j++ {
			if !markedNums[i][j] {
				score += board[i][j]
			}
		}
	}

	return
}

func PlayBoard(board [][]int, calls []int) (isWinner bool, score int, draws int) {
	markedNums := createMarkedNums(board)

	draws = 0

	for _, call := range calls {
		draws++

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == call {
					markedNums[i][j] = true
				}
			}
		}

		if IsWinningBoard(markedNums) {
			isWinner = true
			score = ScoreBoard(board, markedNums) * call
			break
		}
	}

	return
}

func GetBingoBoards(lines []string) (boards [][][]int) {
	boards = [][][]int{}

	board := [][]int{}

	for _, line := range lines {
		nums := strings.Fields(line)

		if len(nums) > 0 {
			row := make([]int, len(nums))
			for i, num := range nums {
				row[i], _ = strconv.Atoi(num)
			}

			board = append(board, row)
		} else {
			// add board and start a new on
			boards = append(boards, board)
			board = [][]int{}
		}
	}

	boards = append(boards, board)

	return
}

func GetBingoCalls(bingoCallLine string) (bingoCalls []int) {
	if bingoCallLine == "" {
		return []int{}
	}

	bingoCallStrs := strings.Split(bingoCallLine, ",")

	bingoCalls = make([]int, len(bingoCallStrs))

	for i, num := range bingoCallStrs {
		bingoCalls[i], _ = strconv.Atoi(num)
	}

	return
}

func main() {
	inputLines := util.LoadInput("./2021/day-04-giant-squid/input.txt")

	bingoCallsInput := "30,35,8,2,39,37,72,7,81,41,25,46,56,18,89,70,0,15,84,75,88,67,42,44,94,71,79,65,58,52,96,83,54,29,14,95,66,61,97,68,57,90,55,32,17,47,20,98,1,69,63,62,31,86,77,85,87,93,26,40,24,19,48,76,73,49,34,45,82,22,80,78,23,6,59,91,64,43,21,51,13,3,53,99,4,28,33,74,12,9,36,50,60,11,27,10,5,16,92,38"
	bingoCalls := GetBingoCalls(bingoCallsInput)

	lowestDrawCount := math.MaxInt32
	scoreOfFastest := 0

	highestDrawCount := 0
	scoreOfSlowest := 0

	boards := GetBingoBoards(inputLines)

	for _, board := range boards {
		if isWinner, score, draws := PlayBoard(board, bingoCalls); isWinner {
			if lowestDrawCount > draws {
				lowestDrawCount = draws
				scoreOfFastest = score
			}

			if highestDrawCount < draws {
				highestDrawCount = draws
				scoreOfSlowest = score
			}
		}
	}

	fmt.Printf("Fastest Score: %d, Slowest Score: %d\n", scoreOfFastest, scoreOfSlowest)
}
