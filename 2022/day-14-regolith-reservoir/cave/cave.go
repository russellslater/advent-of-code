package cave

import (
	"fmt"
	"math"

	"github.com/russellslater/advent-of-code/internal/util"
)

const (
	Rock = '#'
	Sand = 'o'
	Air  = '.'
)

type Cave struct {
	grid       [][]rune
	leftOffset int
}

func (c *Cave) IsEmpty(x, y int) bool {
	return c.grid[y][x] == 0
}

func (c *Cave) IsOutOfBounds(x, y int) bool {
	return x < 0 || x >= len(c.grid[0]) || y > len(c.grid)
}

func (c *Cave) DropSand(startX int, startY int) bool {
	x, y := startX-c.leftOffset, startY

	// Source of sand blocked?
	if c.grid[y][x] == 'o' {
		return false
	}

	downLeft := Point{-1, 1}
	downRight := Point{1, 1}

	for {
		if c.IsOutOfBounds(x, y) {
			return false
		}

		// Blocked below?
		if !c.IsEmpty(x, y+1) {
			// Travel down-left?
			if c.IsOutOfBounds(x+downLeft.X, y+downLeft.Y) {
				return false
			}
			if c.IsEmpty(x+downLeft.X, y+downLeft.Y) {
				x += downLeft.X
				y += downLeft.Y
				continue
			}
			// Travel down-right?
			if c.IsOutOfBounds(x+downRight.X, y+downRight.Y) {
				return false
			}
			if c.IsEmpty(x+downRight.X, y+downRight.Y) {
				x += downRight.X
				y += downRight.Y
				continue
			}

			// Blocked all ways, deposit sand
			c.grid[y][x] = 'o'
			return true
		}

		y += 1 // Continue down
	}
}

func (c *Cave) Print() {
	fmt.Println()
	for _, row := range c.grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Printf("%c", Air)
			} else {
				fmt.Printf("%c", cell)
			}
		}
		fmt.Println()
	}
}

func BuildCave(paths []*Path, includeFloor bool) *Cave {
	// Find the lowest point
	lowest := 0
	for _, path := range paths {
		for _, pt := range path.Points {
			if pt.Y > lowest {
				lowest = pt.Y
			}
		}
	}

	// Find the left-most and right-most points
	left := math.MaxInt
	right := 0
	for _, path := range paths {
		for _, pt := range path.Points {
			if pt.X < left {
				left = pt.X
			}
			if pt.X > right {
				right = pt.X
			}
		}
	}

	if includeFloor {
		lowest += 2
		// Build more width to accomodate maximum size of falling sand triangle
		right = util.Max(right, 500+lowest)
		left = util.Min(left, 500-lowest)
	}

	// Build grid
	grid := make([][]rune, lowest+1)
	for i := range grid {
		grid[i] = make([]rune, right-left+1)
	}

	// Place rock
	for _, path := range paths {
		for idx := 0; idx < len(path.Points)-1; idx++ {
			curr := path.Points[idx]
			next := path.Points[idx+1]
			if curr.X == next.X {
				// Draw vertical line in grid
				if curr.Y > next.Y {
					curr, next = next, curr
				}
				for i := curr.Y; i <= next.Y; i++ {
					grid[i][curr.X-left] = Rock
				}
			} else if curr.Y == next.Y {
				// Draw horizontal line in grid
				if curr.X > next.X {
					curr, next = next, curr
				}
				for i := curr.X; i <= next.X; i++ {
					grid[curr.Y][i-left] = Rock
				}
			}
		}
	}

	if includeFloor {
		for i := 0; i <= right-left; i++ {
			grid[lowest][i] = Rock
		}
	}

	return &Cave{grid, left}
}
