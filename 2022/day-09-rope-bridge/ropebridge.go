package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-09-rope-bridge/ropy"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-09-rope-bridge/input.txt"

	twoKnotRope := ropy.NewRope(2)
	tenKnotRope := ropy.NewRope(10)

	simulateRopeMotions(filename, twoKnotRope, tenKnotRope)

	fmt.Printf("Part One Answer: %d\n", twoKnotRope.TailPositionVisitCount())
	fmt.Printf("Part Two Answer: %d\n", tenKnotRope.TailPositionVisitCount())
}

func simulateRopeMotions(filename string, ropes ...*ropy.Rope) {
	for _, line := range util.LoadInput(filename) {
		var direction rune
		var distance int
		fmt.Sscanf(line, "%c %d", &direction, &distance)

		xDist, yDist, dx, dy := 0, 0, 0, 0

		switch direction {
		case 'L':
			dx = -1
			xDist = distance
		case 'R':
			dx = 1
			xDist = distance
		case 'U':
			dy = -1
			yDist = distance
		case 'D':
			dy = 1
			yDist = distance
		}

		for _, r := range ropes {
			r.MoveHead(xDist, yDist, dx, dy)
		}
	}
}
