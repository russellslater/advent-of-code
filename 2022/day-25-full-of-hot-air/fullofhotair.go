package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-25-full-of-hot-air/fu"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-25-full-of-hot-air/input.txt"
	snafus := getTransformedInput(filename)

	result := calcRequiredSNAFU(snafus)
	fmt.Printf("Part One Answer: %v\n", result)
}

func getTransformedInput(filename string) []fu.SNAFU {
	snafus := []fu.SNAFU{}
	for _, line := range util.LoadInput(filename) {
		snafus = append(snafus, fu.SNAFU(line))
	}
	return snafus
}

func calcRequiredSNAFU(snafus []fu.SNAFU) fu.SNAFU {
	total := 0
	for _, snafu := range snafus {
		total += snafu.Decimal()
	}
	return fu.NewSNAFU(total)
}
