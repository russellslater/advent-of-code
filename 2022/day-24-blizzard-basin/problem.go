package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-24-blizzard-basin/bliz"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-24-blizzard-basin/input.txt"
	basin := getTransformedInput(filename)

	// Debugging
	w, h := basin.InternalWidthHeight()
	fmt.Println("Internal Width/Height", w, h)

	fmt.Println("Basin at t=0")
	basin.PrintAtTime(0)

	fmt.Println("Basin at t=1")
	basin.PrintAtTime(1)

	fmt.Printf("Basin at t=%d\n", w*h)
	basin.PrintAtTime(w * h)

	output := basin.FastestTraversal()
	fmt.Printf("Part One Answer: %v\n", output)
}

func getTransformedInput(filename string) *bliz.Basin {
	basin := bliz.NewBasin()

	y := 0
	var prev rune
	for _, line := range util.LoadInput(filename) {
		x := 0
		prev = 0
		for _, char := range line {
			if char == '#' {
				basin.Walls.Add(bliz.Position{X: x, Y: y})
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
				basin.Blizzards.Add(bliz.Position{X: x, Y: y}, dx, dy)
			} else if char == '.' && prev == '#' {
				if y == 0 {
					basin.Start = bliz.Position{X: x, Y: y}
				} else {
					basin.End = bliz.Position{X: x, Y: y}
				}
			}
			prev = char
			x++
		}
		y++
	}

	return basin
}
