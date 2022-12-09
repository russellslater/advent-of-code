package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-08-treetop-tree-house/forestry"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-08-treetop-tree-house/input.txt"
	forest := getTransformedInput(filename)

	count := forestry.CountVisibleArbores(forest)
	fmt.Printf("Part One Answer: %d\n", count)

	score := forestry.MaxScenicScore(forest)
	fmt.Printf("Part Two Answer: %d\n", score)
}

func getTransformedInput(filename string) forestry.Forest {
	return util.LoadInput(filename)
}
