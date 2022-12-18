package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-17-pyroclastic-flow/tetris"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-17-pyroclastic-flow/input.txt"
	jetDirections := getTransformedInput(filename)

	output := calcHeightOfFallenRocks(jetDirections, 2022)
	fmt.Printf("Part One Answer: %v\n", output)
}

func getTransformedInput(filename string) []rune {
	return []rune(util.LoadInput(filename)[0])
}

func calcHeightOfFallenRocks(jetDirections []rune, stoppedRocksTarget int) int {
	b := tetris.NewBoard(7, 4)

	stoppedRocks := 0
	jetIdx := 0

	var currShape tetris.Shape
	var x, y int
	isFalling := false

	// Game loop
	for {
		// New shape appears?
		if !isFalling {
			currShape = b.NextShape()
			b.Resize(currShape)

			// Initial Placement of shape
			x, y = 2, 0
			b.PlaceShape(currShape, x, y)

			isFalling = true
		} else {
			// Check if shape cannot fall any further
			if !b.CanPlaceShape(currShape, x, y+1) {
				b.FixShape(currShape, x, y)
				isFalling = false
				stoppedRocks++

				// Target reached?
				if stoppedRocks == stoppedRocksTarget {
					return b.Height - b.HighestRockPosition()
				}

				continue
			}

			x, y = b.MoveDown(currShape, x, y)
		}

		// Cycle through jet directions
		direction := jetDirections[jetIdx]
		jetIdx = (jetIdx + 1) % len(jetDirections)

		if direction == '>' {
			// Attempt move right
			if b.CanPlaceShape(currShape, x+1, y) {
				x, y = b.MoveRight(currShape, x, y)
			}
		} else if direction == '<' {
			// Attempt move left
			if b.CanPlaceShape(currShape, x-1, y) {
				x, y = b.MoveLeft(currShape, x, y)
			}
		}
	}
}
