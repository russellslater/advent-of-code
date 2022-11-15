package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2020/day-03-toboggan-trajectory/input.txt"
	landscape := getInputLandscape(filename)

	output := treeEncounters(landscape, move{3, 1})

	fmt.Printf("Tree encounters for first part: %v\n", output)

	moves := []move{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	product := 1
	for _, mv := range moves {
		product *= treeEncounters(landscape, mv)
	}

	fmt.Printf("Tree encounter product for second part: %v\n", product)
}

type move struct {
	right int
	down  int
}

func getInputLandscape(filename string) [][]rune {
	// Can assume that each row has a uniform length
	landscape := [][]rune{}
	for _, line := range util.LoadInput(filename) {
		landscape = append(landscape, []rune(line))
	}
	return landscape
}

func treeEncounters(landscape [][]rune, mv move) int {
	height := len(landscape)
	width := len(landscape[0])
	x, y := 0, 0
	count := 0

	for {
		// fmt.Printf("?: %c, y: %d, x: %d\n", landscape[y][x], y, x)
		if landscape[y][x] == '#' {
			count++
		}

		// Modulo operator to wrap back around as if multi-dim array repeated to the right
		x = (x + mv.right) % width
		y += mv.down

		if y >= height {
			break
		}
	}

	return count
}
