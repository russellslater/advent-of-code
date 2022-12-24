package main

import (
	"fmt"
	"time"

	"github.com/russellslater/advent-of-code/2022/day-24-blizzard-basin/bliz"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-24-blizzard-basin/input.txt"
	basin := getTransformedInput(filename)

	t := time.Now()

	fastestTime := basin.FastestTraversal()
	fmt.Printf("Part One Answer: %d (%s)\n", fastestTime, time.Since(t))
}

func getTransformedInput(filename string) *bliz.Basin {
	walls := bliz.WallSet{}
	blizzards := bliz.BlizzardSet{}
	var start, end bliz.Position

	y := 0
	var prev rune
	for _, line := range util.LoadInput(filename) {
		x := 0
		prev = 0
		for _, char := range line {
			if char == '#' {
				walls.Add(bliz.Position{X: x, Y: y})
			} else if char == '<' || char == '>' || char == '^' || char == 'v' {
				dx := 0
				dy := 0
				switch char {
				case '<':
					dx = -1
				case '>':
					dx = 1
				case '^':
					dy = -1
				case 'v':
					dy = 1
				}
				blizzards.Add(bliz.Position{X: x, Y: y}, dx, dy)
			} else if char == '.' && prev == '#' {
				if y == 0 {
					start = bliz.Position{X: x, Y: y}
				} else {
					end = bliz.Position{X: x, Y: y}
				}
			}
			prev = char
			x++
		}
		y++
	}

	return bliz.NewBasin(start, end, walls, blizzards)
}
