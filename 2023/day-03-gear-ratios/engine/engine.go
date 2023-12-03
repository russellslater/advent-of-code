package engine

import (
	"github.com/russellslater/advent-of-code/internal/twod"
)

type Schematic []string

func (s Schematic) IsOutOfBounds(p twod.Position) bool {
	if p.Y < 0 || p.Y >= len(s) {
		return true
	}

	if p.X < 0 || p.X >= len(s[p.Y]) {
		return true
	}

	return false
}

func (s Schematic) IsSymbol(p twod.Position) bool {
	if s.IsOutOfBounds(p) {
		return false
	}

	// a symbol is any character other than 0-9 or .
	return s[p.Y][p.X] != '.' && (s[p.Y][p.X] < '0' || s[p.Y][p.X] > '9')
}

func (s Schematic) IsGear(p twod.Position) bool {
	if s.IsOutOfBounds(p) {
		return false
	}

	return s[p.Y][p.X] == '*'
}

func (s Schematic) isNeighbouringSymbol(positions []twod.Position) bool {
	for _, p := range positions {
		for _, n := range p.OctilinearNeighbours() {
			if s.IsSymbol(n) {
				return true
			}
		}
	}
	return false
}

func (s Schematic) isNeighbouringGear(positions []twod.Position) (bool, []twod.Position) {
	found := false
	gearPositions := []twod.Position{}

	for _, p := range positions {
		for _, n := range p.OctilinearNeighbours() {
			if s.IsGear(n) {
				found = true

				alreadyFound := false
				for _, gp := range gearPositions {
					if gp == n {
						alreadyFound = true
						break
					}
				}
				if !alreadyFound {
					gearPositions = append(gearPositions, n)
				}
			}
		}
	}

	return found, gearPositions
}

type EngineAssessor struct {
	PartSum            int
	GearRatioSum       int
	PossibleGearRatios map[twod.Position][]int
}

func NewEngineAssessor() *EngineAssessor {
	return &EngineAssessor{
		PartSum:            0,
		GearRatioSum:       0,
		PossibleGearRatios: map[twod.Position][]int{},
	}
}

func (e *EngineAssessor) AssessEngine(schematic Schematic) {
	for y, row := range schematic {
		isDigit := false
		currentNumber := 0
		positions := []twod.Position{}
		for x, c := range row {
			if c >= '0' && c <= '9' {
				isDigit = true
				currentNumber = currentNumber*10 + int(c-'0')
				positions = append(positions, twod.Position{X: x, Y: y})
			} else {
				// engine part identified?
				if isDigit {
					e.processEnginePart(currentNumber, schematic, positions)
				}
				isDigit = false
				currentNumber = 0
				positions = []twod.Position{}
			}
		}

		// end of row — engine part identified?
		if isDigit {
			e.processEnginePart(currentNumber, schematic, positions)
		}
	}

	e.calculateGearRatio(e.PossibleGearRatios)
}

func (e *EngineAssessor) processEnginePart(num int, schematic Schematic, positions []twod.Position) {
	if schematic.isNeighbouringSymbol(positions) {
		e.PartSum += num

		ok, gearPositions := schematic.isNeighbouringGear(positions)
		if ok {
			for _, gp := range gearPositions {
				e.PossibleGearRatios[gp] = append(e.PossibleGearRatios[gp], num)
			}
		}
	}
}

func (e *EngineAssessor) calculateGearRatio(possibleGearRatios map[twod.Position][]int) {
	for _, nums := range possibleGearRatios {
		// must be exactly two numbers to calculate a gear ratio
		if len(nums) == 2 {
			e.GearRatioSum += nums[0] * nums[1]
		}
	}
}
