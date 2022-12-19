package main

import (
	"fmt"
	"strings"

	"github.com/russellslater/advent-of-code/2022/day-18-boiling-boulders/lava"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-18-boiling-boulders/input.txt"
	droplet := getTransformedInput(filename)

	surfaceArea, edgeCubes := calcDropletSurfaceArea(droplet)
	fmt.Printf("Part One Answer: %v\n", surfaceArea)

	exteriorSurfaceArea := calcExteriorSurfaceArea(droplet, edgeCubes)
	fmt.Printf("Part Two Answer: %v\n", exteriorSurfaceArea)
}

func getTransformedInput(filename string) *lava.LavaDroplet {
	cubes := []lava.Cube{}
	for _, line := range util.LoadInput(filename) {
		s := strings.Split(line, ",")
		cube := lava.Cube{X: util.MustAtoi(s[0]), Y: util.MustAtoi(s[1]), Z: util.MustAtoi(s[2])}
		cubes = append(cubes, cube)
	}
	return lava.NewLavaDroplet(cubes)
}

func calcDropletSurfaceArea(droplet *lava.LavaDroplet) (int, []lava.Cube) {
	edgeCubes := []lava.Cube{}
	surfaceArea := 0

	for _, cube := range droplet.Cubes {
		isExposed := false
		for _, neighbour := range cube.PossibleNeighbours() {
			if !droplet.IsOccupied(neighbour) {
				surfaceArea++
				isExposed = true
			}
		}
		if isExposed {
			edgeCubes = append(edgeCubes, cube)
		}
	}

	return surfaceArea, edgeCubes
}

func calcExteriorSurfaceArea(droplet *lava.LavaDroplet, edgeCubes []lava.Cube) int {
	exteriorSurfaceArea := 0
	var queue []lava.Cube
	visited := make([]lava.Cube, 0)

	queue = append(queue, droplet.MinCube())

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, neighbour := range curr.PossibleNeighbours() {
			if contains(edgeCubes, neighbour) {
				exteriorSurfaceArea++
			} else if !contains(visited, neighbour) && droplet.IsInsideBounds(neighbour) {
				visited = append(visited, neighbour)
				queue = append(queue, neighbour)
			}
		}
	}

	return exteriorSurfaceArea
}

func contains(cubes []lava.Cube, cube lava.Cube) bool {
	for _, c := range cubes {
		if c == cube {
			return true
		}
	}
	return false
}
