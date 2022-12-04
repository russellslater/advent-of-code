package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-04-camp-cleanup/input.txt"
	assignments := getTransformedInput(filename)

	count := impactedAssignmentCount(assignments, rangeContains)
	fmt.Printf("Part One Answer: %d\n", count)

	count = impactedAssignmentCount(assignments, rangeOverlaps)
	fmt.Printf("Part Two Answer: %d\n", count)
}

type assignment struct {
	range1 []int
	range2 []int
}

func getTransformedInput(filename string) []*assignment {
	input := []*assignment{}
	for _, line := range util.LoadInput(filename) {
		parts := strings.Split(line, ",")

		a := &assignment{
			range1: parsePair(parts[0]),
			range2: parsePair(parts[1]),
		}

		input = append(input, a)
	}
	return input
}

func parsePair(pair string) []int {
	var start, end int
	fmt.Sscanf(pair, "%d-%d", &start, &end)
	r := []int{}
	for i := start; i <= end; i++ {
		r = append(r, i)
	}
	return r
}

func impactedAssignmentCount(assignments []*assignment, checker rangeChecker) int {
	count := 0
	for _, a := range assignments {
		if checker(a.range1, a.range2) || checker(a.range2, a.range1) {
			count++
		}
	}
	return count
}

type rangeChecker func([]int, []int) bool

func rangeContains(target []int, other []int) bool {
	return containsInt(target, other[0]) && containsInt(target, other[len(other)-1])
}

func rangeOverlaps(target []int, other []int) bool {
	return containsInt(target, other[0]) || containsInt(target, other[len(other)-1])
}

func containsInt(arr []int, num int) bool {
	for _, el := range arr {
		if el == num {
			return true
		}
	}
	return false
}
