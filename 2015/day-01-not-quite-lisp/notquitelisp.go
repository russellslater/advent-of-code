package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2015/day-01-not-quite-lisp/input.txt"

	floor := computeFloor(filename)
	fmt.Printf("Part One Answer: %d\n", floor)

	floor = positionAtFirstBasementEntry(filename)
	fmt.Printf("Part Two Answer: %d\n", floor)
}

type Santa struct {
	floor int
}

func (s *Santa) move(direction rune) {
	switch direction {
	case '(':
		s.floor++
	case ')':
		s.floor--
	}
}

func (s *Santa) isInBasement() bool {
	return s.floor <= -1
}

func computeFloor(filename string) int {
	s := Santa{}
	for _, r := range util.LoadInput(filename)[0] {
		s.move(r)
	}
	return s.floor
}

func positionAtFirstBasementEntry(filename string) int {
	position := 1
	s := Santa{}
	for _, r := range util.LoadInput(filename)[0] {
		s.move(r)
		if s.isInBasement() {
			return position
		}
		position++
	}
	return s.floor
}
