package cave

import (
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

type Path struct {
	Points []Point
}

type Point struct {
	X int
	Y int
}

// Example input: "498,4 -> 498,6 -> 496,6"
func NewPath(input string) *Path {
	p := &Path{}
	for _, pt := range strings.Split(input, " -> ") {
		parts := strings.Split(pt, ",")
		p.Points = append(p.Points, Point{util.MustAtoi(parts[0]), util.MustAtoi(parts[1])})
	}
	return p
}
