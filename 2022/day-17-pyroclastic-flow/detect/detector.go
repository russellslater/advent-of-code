package detect

import "github.com/russellslater/advent-of-code/2022/day-17-pyroclastic-flow/tetris"

type Detector interface {
	Detect(board *tetris.Board)
	IncrementFallCount(board *tetris.Board)
	IncrementStoppedRockCount(board *tetris.Board) bool // Returns true if the detector has detected the answer
}
