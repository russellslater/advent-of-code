package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2015/day-02-i-was-told-there-would-be-no-math/input.txt"
	boxes := getTransformedInput(filename)

	totalSurfaceArea := totalSquareFeetWrappingPaper(boxes)
	fmt.Printf("Part One Answer: %d\n", totalSurfaceArea)

	totalFeetRibbon := totalFeetRibbon(boxes)
	fmt.Printf("Part Two Answer: %d\n", totalFeetRibbon)
}

type box struct {
	l int
	w int
	h int
}

func (b *box) surfaceArea() int {
	return 2*b.l*b.w + 2*b.w*b.h + 2*b.h*b.l
}

func (b *box) areaSmallestSide() int {
	smallest := b.l * b.w
	if smallest > b.w*b.h {
		smallest = b.w * b.h
	}
	if smallest > b.h*b.l {
		smallest = b.h * b.l
	}
	return smallest
}

func (b *box) perimeterSmallestSide() int {
	smallest := 2*b.l + 2*b.w
	if smallest > 2*b.w+2*b.h {
		smallest = 2*b.w + 2*b.h
	}
	if smallest > 2*b.h+2*b.l {
		smallest = 2*b.h + 2*b.l
	}
	return smallest
}

func (b *box) cubicVolume() int {
	return b.l * b.w * b.h
}

func getTransformedInput(filename string) []*box {
	boxes := []*box{}
	for _, line := range util.LoadInput(filename) {
		b := &box{}
		fmt.Sscanf(line, "%dx%dx%d", &b.l, &b.w, &b.h)
		boxes = append(boxes, b)
	}
	return boxes
}

func totalSquareFeetWrappingPaper(boxes []*box) int {
	total := 0
	for _, b := range boxes {
		total += b.surfaceArea() + b.areaSmallestSide()
	}
	return total
}

func totalFeetRibbon(boxes []*box) int {
	total := 0
	for _, b := range boxes {
		total += b.perimeterSmallestSide() + b.cubicVolume()
	}
	return total
}
