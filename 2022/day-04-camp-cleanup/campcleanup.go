package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-04-camp-cleanup/asspair"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-04-camp-cleanup/input.txt"
	assignments := getTransformedInput(filename)

	count := impactedAssignmentCount(assignments, containedMatch)
	fmt.Printf("Part One Answer: %d\n", count)

	count = impactedAssignmentCount(assignments, overlapMatch)
	fmt.Printf("Part Two Answer: %d\n", count)
}

func getTransformedInput(filename string) []*asspair.AssignmentPair {
	input := []*asspair.AssignmentPair{}
	for _, line := range util.LoadInput(filename) {
		a := &asspair.AssignmentPair{}
		fmt.Sscanf(line, "%d-%d,%d-%d", &a.One.Start, &a.One.End, &a.Two.Start, &a.Two.End)
		input = append(input, a)
	}
	return input
}

func impactedAssignmentCount(assignments []*asspair.AssignmentPair, m match) int {
	count := 0
	for _, a := range assignments {
		if m(a) {
			count++
		}
	}
	return count
}

type match func(*asspair.AssignmentPair) bool

func overlapMatch(a *asspair.AssignmentPair) bool { return a.IsOverlap() }

func containedMatch(a *asspair.AssignmentPair) bool { return a.IsContained() }
