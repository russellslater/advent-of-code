package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-03-rucksack-reorganization/rksk"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-03-rucksack-reorganization/input.txt"
	rucksacks := getTransformedInput(filename)

	score := rearrangementPriorityScore(rucksacks)
	fmt.Printf("Part One Answer: %d\n", score)

	score = badgeStickerPriorityScore(rucksacks)
	fmt.Printf("Part Two Answer: %d\n", score)
}

func getTransformedInput(filename string) []*rksk.Rucksack {
	input := []*rksk.Rucksack{}
	for _, line := range util.LoadInput(filename) {
		input = append(input, &rksk.Rucksack{Items: line})
	}
	return input
}

func rearrangementPriorityScore(rucksacks []*rksk.Rucksack) int {
	score := 0
	for _, r := range rucksacks {
		score += r.ReoccuringItem().Score()
	}
	return score
}

func badgeStickerPriorityScore(rucksacks []*rksk.Rucksack) int {
	score := 0
	for _, g := range rksk.GroupElves(rucksacks, 3) {
		score += g.CommonItem().Score()
	}
	return score
}
