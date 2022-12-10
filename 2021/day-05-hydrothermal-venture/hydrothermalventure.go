package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

type Segment struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func (s Segment) getCoveredCoordinates() (coords [][]int) {
	if s.isHorizontal() {
		startX, endX := s.x1, s.x2
		if startX > endX {
			endX, startX = startX, endX
		}
		for x := startX; x <= endX; x++ {
			coords = append(coords, []int{x, s.y1})
		}
		return
	} else if s.isVertical() {
		startY, endY := s.y1, s.y2
		if startY > endY {
			endY, startY = startY, endY
		}
		for y := startY; y <= endY; y++ {
			coords = append(coords, []int{s.x1, y})
		}
		return
	} else if s.isDiagonal() {
		startY, endY := s.y1, s.y2
		startX, endX := s.x1, s.x2

		modY := 1
		if startY > endY {
			modY = -1
		}

		modX := 1
		if startX > endX {
			modX = -1
		}

		for y, x := startY, startX; (modY == -1 && y >= endY) || (modY == 1 && y <= endY); y, x = y+modY, x+modX {
			coords = append(coords, []int{x, y})
		}
	}

	return
}

func (s Segment) isDiagonal() (isDiagonal bool) {
	isDiagonal = s.y1 != s.y2 && s.x1 != s.x2
	return
}

func (s Segment) isHorizontal() (isHorizontal bool) {
	isHorizontal = s.y1 == s.y2
	return
}

func (s Segment) isVertical() (isVertical bool) {
	isVertical = s.x1 == s.x2
	return
}

func CreateSegment(line string) Segment {
	// example: 308,411 -> 656,63
	coords := strings.Split(line, " -> ")
	startCoords := strings.Split(coords[0], ",")
	endCoords := strings.Split(coords[1], ",")

	x1, _ := strconv.Atoi(startCoords[0])
	y1, _ := strconv.Atoi(startCoords[1])
	x2, _ := strconv.Atoi(endCoords[0])
	y2, _ := strconv.Atoi(endCoords[1])

	return Segment{
		x1: x1,
		y1: y1,
		x2: x2,
		y2: y2,
	}
}

func CreateSegments(lines []string) (segments []Segment, maxX int, maxY int) {
	for _, line := range lines {
		s := CreateSegment(line)
		segments = append(segments, s)

		if s.x1 > maxX {
			maxX = s.x1
		} else if s.x2 > maxX {
			maxX = s.x2
		}

		if s.y1 > maxY {
			maxY = s.y1
		} else if s.y2 > maxY {
			maxY = s.y2
		}
	}
	return
}

func PlotSegments(segments []Segment, maxX int, maxY int, includeDiagonals bool) (oceanFloor [][]int) {
	oceanFloor = make([][]int, maxY+1)
	for i := range oceanFloor {
		oceanFloor[i] = make([]int, maxX+1)
	}

	for _, s := range segments {
		if s.isDiagonal() && !includeDiagonals {
			continue
		}

		coords := s.getCoveredCoordinates()
		for _, c := range coords {
			x := c[0]
			y := c[1]

			oceanFloor[y][x] += 1
		}
	}

	return
}

func CountSegmentOverlaps(oceanFloor [][]int) (overlapCount int) {
	for i := 0; i < len(oceanFloor); i++ {
		for j := 0; j < len(oceanFloor[i]); j++ {
			if oceanFloor[i][j] > 1 {
				overlapCount++
			}
		}
	}

	return
}

func ProcessSegmentsForDanger(lines []string, includeDiagonals bool) (overlapCount int) {
	segments, maxX, maxY := CreateSegments(lines)
	oceanFloor := PlotSegments(segments, maxX, maxY, includeDiagonals)
	return CountSegmentOverlaps(oceanFloor)
}

func main() {
	inputLines := util.LoadInput("./2021/day-05-hydrothermal-venture/input.txt")

	overlapCount := ProcessSegmentsForDanger(inputLines, false)

	fmt.Printf("Overlap Count: %d\n", overlapCount)

	overlapCountWithDiagonals := ProcessSegmentsForDanger(inputLines, true)

	fmt.Printf("Overlap Count with Diagonals: %d\n", overlapCountWithDiagonals)
}
