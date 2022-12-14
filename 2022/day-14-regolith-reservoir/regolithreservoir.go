package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-14-regolith-reservoir/input.txt"
	paths := getTransformedInput(filename)

	count := calcUnitsOfSandAtRest(paths)
	fmt.Printf("Part One Answer: %d\n", count)
}

type Path struct {
	Points []Point
}

type Point struct {
	X int
	Y int
}

// Example input: "498,4 -> 498,6 -> 496,6"
func parsePath(input string) Path {
	p := Path{}
	for _, pt := range strings.Split(input, " -> ") {
		fmt.Println(pt, strings.Split(pt, ","))
		parts := strings.Split(pt, ",")
		p.Points = append(p.Points, Point{util.MustAtoi(parts[0]), util.MustAtoi(parts[1])})
	}
	return p
}

func getTransformedInput(filename string) []Path {
	paths := []Path{}
	for _, line := range util.LoadInput(filename) {
		paths = append(paths, parsePath(line))
	}
	return paths
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%c", cell)
			}
		}
		fmt.Println()
	}
}

func buildGrid(paths []Path) ([][]rune, int) {
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
					grid[i][curr.X-left] = '#'
				}
			} else if curr.Y == next.Y {
				// Draw horizontal line in grid
				if curr.X > next.X {
					curr, next = next, curr
				}
				for i := curr.X; i <= next.X; i++ {
					grid[curr.Y][i-left] = '#'
				}
			}
		}
	}

	return grid, left
}

func calcUnitsOfSandAtRest(paths []Path) interface{} {
	grid, leftOffset := buildGrid(paths)

	printGrid(grid)

	// sand pours in from the top (y=0) at a fixed point (x=500)
	start := Point{500, 0}

	endIdx := 0
	for i := 0; i < 1000; i++ {
		if !dropSand(grid, leftOffset, start) {
			endIdx = i
			break
		}
	}

	fmt.Println()
	printGrid(grid)

	return endIdx
}

func isOutOfBounds(grid [][]rune, x, y int) bool {
	return x < 0 || x >= len(grid[0]) || y > len(grid)
}

func dropSand(grid [][]rune, leftOffset int, start Point) bool {
	x, y := start.X-leftOffset, start.Y

	downLeft := Point{-1, 1}
	downRight := Point{1, 1}

	for {
		if isOutOfBounds(grid, x, y) {
			return false
		}

		// blocked below?
		if grid[y+1][x] != 0 {
			// travel down-left?
			if isOutOfBounds(grid, x+downLeft.X, y+downLeft.Y) {
				return false
			}
			if grid[y+downLeft.Y][x+downLeft.X] == 0 {
				x += downLeft.X
				y += downLeft.Y
				continue
			}
			// travel down-right?
			if isOutOfBounds(grid, x+downRight.X, y+downRight.Y) {
				return false
			}
			if grid[y+downRight.Y][x+downRight.X] == 0 {
				x += downRight.X
				y += downRight.Y
				continue
			}

			// blocked all ways, deposit sand
			grid[y][x] = 'o'
			return true
		}

		y += 1 // continue down
	}
}
