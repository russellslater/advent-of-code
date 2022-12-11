package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2015/day-03-perfectly-spherical-houses-in-a-vacuum/input.txt"
	moves := getTransformedInput(filename)

	total := totalHousesWithPresent(moves)
	fmt.Printf("Part One Answer: %d\n", total)

	total = totalHousesWithPresentWithRoboSanta(moves)
	fmt.Printf("Part Two Answer: %d\n", total)
}

func getTransformedInput(filename string) []rune {
	moves := []rune{}
	for _, r := range util.LoadInput(filename)[0] {
		moves = append(moves, r)
	}
	return moves
}

type location struct{ x, y int }

func totalHousesWithPresent(moves []rune) int {
	presents := make(map[location]bool)

	x, y := 0, 0

	presents[location{x, y}] = true

	for _, mv := range moves {
		switch mv {
		case '>':
			x++
		case '<':
			x--
		case 'v':
			y++
		case '^':
			y--
		}
		presents[location{x, y}] = true
	}

	return len(presents)
}

func totalHousesWithPresentWithRoboSanta(moves []rune) int {
	presents := make(map[location]bool)

	santa := location{0, 0}
	roboSanta := location{0, 0}

	presents[santa] = true

	for i, mv := range moves {
		deliverer := &roboSanta
		if i%2 == 0 {
			deliverer = &santa
		}
		switch mv {
		case '>':
			deliverer.x++
		case '<':
			deliverer.x--
		case 'v':
			deliverer.y++
		case '^':
			deliverer.y--
		}
		presents[*deliverer] = true
	}

	return len(presents)
}
