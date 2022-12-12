package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/fzipp/astar"
	"github.com/russellslater/advent-of-code/2022/day-12-hill-climbing-algorithm/hills"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	hmap := getTransformedInput("./2022/day-12-hill-climbing-algorithm/input.txt")

	steps := shortestPathFromAvailableStarts(hmap, false)
	fmt.Printf("Part One Answer: %d\n", steps)

	steps = shortestPathFromAvailableStarts(hmap, true)
	fmt.Printf("Part Two Answer: %d\n", steps)
}

func shortestPathFromAvailableStarts(hmap hills.Heightmap, includeAltStarts bool) int {
	starts, dest := hmap.GetStartsAndDest(includeAltStarts)
	shortest := math.MaxInt
	for _, p := range starts {
		path := astar.FindPath[hills.Position](hmap, p, dest, hmap.Cost, hmap.Cost)
		if len(path) == 0 {
			continue // no path found
		}
		if len(path) < shortest {
			shortest = len(path)
		}
	}
	return shortest - 1
}

func getTransformedInput(filename string) hills.Heightmap {
	input := [][]string{}
	for _, line := range util.LoadInput(filename) {
		input = append(input, strings.Split(line, ""))
	}
	return input
}
