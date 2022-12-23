package detect

import "github.com/russellslater/advent-of-code/2022/day-17-pyroclastic-flow/tetris"

type RepetitionDetector struct {
	// Number of times the max fall count must be repeated to be consider a pattern
	threshhold         int
	stoppedRocksTarget int
	stoppedRocks       int

	maxFallCount       int
	fallCount          int
	maxFallRepetitions [][]int // [stoppedRocksPerRepeat, heightPerRepeat]

	thresholdHit                    bool
	postThresholdStoppedRocks       int
	postThresholdStoppedRocksTarget int
	postThresholdHeight             int

	// Cached to compute answer
	stoppedRocksBeforeRepeat int
	heightBeforeRepeat       int
	stoppedRocksPerRepeat    int
	heightPerRepeat          int

	answer int
}

func NewRepetitionDetector(stoppedRocksTarget int) *RepetitionDetector {
	return &RepetitionDetector{
		threshhold:         5,
		stoppedRocksTarget: stoppedRocksTarget,
		maxFallRepetitions: [][]int{},
	}
}

func (r *RepetitionDetector) TowerHeight() int {
	return r.answer
}

func (r *RepetitionDetector) IncrementFallCount(board *tetris.Board) {
	r.fallCount++
}

func (r *RepetitionDetector) IncrementStoppedRockCount(board *tetris.Board) bool {
	r.stoppedRocks++

	if r.thresholdHit {
		r.postThresholdStoppedRocks++

		// Cannot reasonably execute until 1_000_000_000_000 has been hit.
		// Check if we have reached the target number of stopped rocks inside of one repeat then calculate the answer.
		if r.postThresholdStoppedRocks == r.postThresholdStoppedRocksTarget {
			repeats := (r.stoppedRocksTarget - r.stoppedRocksBeforeRepeat) / r.stoppedRocksPerRepeat
			heightOfRepeats := repeats * r.heightPerRepeat
			heightFollowingLastFullRepeat := (board.Height - board.HighestRockPosition()) - r.postThresholdHeight
			r.answer = heightOfRepeats + r.heightBeforeRepeat + heightFollowingLastFullRepeat
			return true
		}
	}
	return false
}

func (r *RepetitionDetector) Detect(board *tetris.Board) {
	// Reset fall count if greater than max
	if r.fallCount > r.maxFallCount {
		r.maxFallCount = r.fallCount
		r.maxFallRepetitions = [][]int{}
	}

	if r.maxFallCount == r.fallCount {
		r.maxFallRepetitions = append(r.maxFallRepetitions, []int{
			r.stoppedRocks, board.Height - board.HighestRockPosition(),
		})

		// Hit threshold yet?
		if len(r.maxFallRepetitions) == r.threshhold {
			r.thresholdHit = true
			curr := r.maxFallRepetitions[len(r.maxFallRepetitions)-1]
			prev := r.maxFallRepetitions[len(r.maxFallRepetitions)-2]

			// Cache values ahead of computing answer
			r.stoppedRocksBeforeRepeat = r.maxFallRepetitions[0][0]
			r.heightBeforeRepeat = r.maxFallRepetitions[0][1]
			r.stoppedRocksPerRepeat = curr[0] - prev[0]
			r.heightPerRepeat = curr[1] - prev[1]

			// Calculate to determine how far to continue through next repeat
			r.postThresholdStoppedRocksTarget = (r.stoppedRocksTarget - r.stoppedRocksBeforeRepeat) % r.stoppedRocksPerRepeat
			r.postThresholdHeight = curr[1]
		}
	}

	r.fallCount = 0
}
