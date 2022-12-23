package detect

import "github.com/russellslater/advent-of-code/2022/day-17-pyroclastic-flow/tetris"

type StoppedRocksDetector struct {
	stoppedRocksTarget int
	stoppedRocks       int
	answer             int
}

func NewStoppedRocksDetector(stoppedRocksTarget int) *StoppedRocksDetector {
	return &StoppedRocksDetector{
		stoppedRocksTarget: stoppedRocksTarget,
		stoppedRocks:       0,
	}
}

func (d *StoppedRocksDetector) TowerHeight() int {
	return d.answer
}

func (d *StoppedRocksDetector) Detect(board *tetris.Board) {
	// No-op
}

func (d *StoppedRocksDetector) IncrementFallCount(board *tetris.Board) {
	// No-op
}

func (d *StoppedRocksDetector) IncrementStoppedRockCount(board *tetris.Board) bool {
	d.stoppedRocks++
	d.answer = board.Height - board.HighestRockPosition()
	return d.stoppedRocks == d.stoppedRocksTarget
}
