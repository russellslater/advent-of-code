package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2015/day-05-doesnt-he-have-intern-elves-for-this/santastr"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2015/day-05-doesnt-he-have-intern-elves-for-this/input.txt"

	strs := getTransformedInput(filename, santastr.NewSantaStringV1)
	count := countNiceStrings(strs)
	fmt.Printf("Part One Answer: %d\n", count)

	strs = getTransformedInput(filename, santastr.NewSantaStringV2)
	count = countNiceStrings(strs)
	fmt.Printf("Part Two Answer: %d\n", count)
}

func getTransformedInput(filename string, factory func(string) santastr.SantaString) []santastr.SantaString {
	santaStrs := []santastr.SantaString{}
	for _, line := range util.LoadInput(filename) {
		santaStrs = append(santaStrs, factory(line))
	}
	return santaStrs
}

func countNiceStrings(strs []santastr.SantaString) int {
	niceCount := 0
	for _, str := range strs {
		if str.IsNice() {
			niceCount++
		}
	}
	return niceCount
}
