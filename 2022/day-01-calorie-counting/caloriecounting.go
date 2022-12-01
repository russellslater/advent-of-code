package main

import (
	"fmt"
	"strconv"

	"github.com/russellslater/advent-of-code/2022/day-01-calorie-counting/elfcal"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-01-calorie-counting/input.txt"
	ec := getTransformedInput(filename)

	top := elfcal.TopElfTotalCalories(ec, 1)
	fmt.Printf("Part One Answer: %d\n", top)

	topThree := elfcal.TopElfTotalCalories(ec, 3)
	fmt.Printf("Part 2 Answer: %d\n", topThree)
}

func getTransformedInput(filename string) []elfcal.ElfCalories {
	ec := []elfcal.ElfCalories{{}}
	for _, line := range util.LoadInput(filename) {
		if line == "" {
			ec = append(ec, elfcal.ElfCalories{})
			continue
		}
		num, _ := strconv.Atoi(line)
		ec[len(ec)-1] = append(ec[len(ec)-1], num)
	}
	return ec
}
