package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2023/day-03-gear-ratios/engine"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2023/day-03-gear-ratios/input.txt"

	e := engine.NewEngineAssessor()
	e.AssessEngine(engine.Schematic(util.LoadInput(filename)))

	fmt.Printf("Part One Answer: %d\n", e.PartSum)
	fmt.Printf("Part Two Answer: %d\n", e.GearRatioSum)
}
