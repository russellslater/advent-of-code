package lava

import (
	"math"

	"github.com/russellslater/advent-of-code/internal/util"
)

type LavaDroplet struct {
	Cubes          []Cube
	Bounds         [6]int
	occupiedSpaces map[Cube]bool
}

func NewLavaDroplet(cubes []Cube) *LavaDroplet {
	d := &LavaDroplet{Cubes: cubes, Bounds: calcBounds(cubes)}
	d.occupiedSpaces = map[Cube]bool{}
	for _, cube := range d.Cubes {
		d.occupiedSpaces[cube] = true
	}
	return d
}

func (d *LavaDroplet) IsInsideBounds(cube Cube) bool {
	// Ensure cubes at the extremities are included (-1, +1)
	return cube.X >= d.Bounds[0]-1 && cube.X <= d.Bounds[1]+1 &&
		cube.Y >= d.Bounds[2]-1 && cube.Y <= d.Bounds[3]+1 &&
		cube.Z >= d.Bounds[4]-1 && cube.Z <= d.Bounds[5]+1
}

func (d *LavaDroplet) MinCube() Cube {
	return Cube{d.Bounds[0], d.Bounds[2], d.Bounds[4]}
}

func (d *LavaDroplet) IsOccupied(cube Cube) bool {
	return d.occupiedSpaces[cube]
}

func calcBounds(cubes []Cube) [6]int {
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt
	minZ, maxZ := math.MaxInt, math.MinInt

	for _, cube := range cubes {
		minX = util.Min(minX, cube.X)
		maxX = util.Max(maxX, cube.X)
		minY = util.Min(minY, cube.Y)
		maxY = util.Max(maxY, cube.Y)
		minZ = util.Min(minZ, cube.Z)
		maxZ = util.Max(maxZ, cube.Z)
	}

	return [6]int{minX, maxX, minY, maxY, minZ, maxZ}
}

type Cube struct {
	X int
	Y int
	Z int
}

func (c Cube) PossibleNeighbours() []Cube {
	return []Cube{
		{c.X + 1, c.Y, c.Z},
		{c.X - 1, c.Y, c.Z},
		{c.X, c.Y + 1, c.Z},
		{c.X, c.Y - 1, c.Z},
		{c.X, c.Y, c.Z + 1},
		{c.X, c.Y, c.Z - 1},
	}
}
