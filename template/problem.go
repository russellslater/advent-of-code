package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./template/input.txt"
	inputLines := getTransformedInput(filename)
	output := solveProblem(inputLines)
	fmt.Printf("Solution output: %v\n", output)
}

// Custom input return type.
type inputLine struct {
	line string
}

// Update return type as required.
func getTransformedInput(filename string) []*inputLine {
	input := []*inputLine{}
	for _, line := range util.LoadInput(filename) {
		// Code to manipulate input here.
		input = append(input, &inputLine{line})
	}
	return input
}

func solveProblem(inputLines []*inputLine) interface{} {
	// Code to solve problem here.
	return true
}
