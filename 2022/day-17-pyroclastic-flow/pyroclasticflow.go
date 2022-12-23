package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/2022/day-17-pyroclastic-flow/detect"
	"github.com/russellslater/advent-of-code/2022/day-17-pyroclastic-flow/tetris"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-17-pyroclastic-flow/input.txt"
	jetDirections := getTransformedInput(filename)

	d := detect.NewStoppedRocksDetector(2022)
	rainTetrisBlocks(jetDirections, d)
	fmt.Printf("Part One Answer: %v\n", d.TowerHeight())

	rd := detect.NewRepetitionDetector(1_000_000_000_000)
	rainTetrisBlocks(jetDirections, rd)
	fmt.Printf("Part Two Answer: %v\n", rd.TowerHeight())
}

func getTransformedInput(filename string) []rune {
	return []rune(util.LoadInput(filename)[0])
}

func rainTetrisBlocks(jetDirections []rune, detector detect.Detector) {
	b := tetris.NewBoard(7, 4)

	var currShape tetris.Shape
	var x, y int
	isFalling := false
	jetIdx := 0

	// Game loop
	for {
		// New shape appears?
		if !isFalling {
			detector.Detect(b)

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

				if finished := detector.IncrementStoppedRockCount(b); finished {
					return
				}

				continue
			}

			x, y = b.MoveDown(currShape, x, y)

			detector.IncrementFallCount(b)
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
