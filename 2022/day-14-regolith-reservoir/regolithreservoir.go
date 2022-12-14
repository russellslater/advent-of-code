package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-14-regolith-reservoir/cave"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-14-regolith-reservoir/input.txt"
	paths := getTransformedInput(filename)

	count := calcUnitsOfSandAtRest(paths, false)
	fmt.Printf("Part One Answer: %d\n", count)

	count = calcUnitsOfSandAtRest(paths, true)
	fmt.Printf("Part Two Answer: %d\n", count)
}

func getTransformedInput(filename string) []*cave.Path {
	paths := []*cave.Path{}
	for _, line := range util.LoadInput(filename) {
		paths = append(paths, cave.NewPath(line))
	}
	return paths
}

func calcUnitsOfSandAtRest(paths []*cave.Path, includeFloor bool) int {
	cave := cave.BuildCave(paths, includeFloor)

	// printGrid(grid)

	count := 0
	for {
		// Sand pours in from the top of the cave (y=0) at a fixed point (x=500)
		if !cave.DropSand(500, 0) {
			break
		}
		count++
	}

	// printGrid(grid)

	return count
}
