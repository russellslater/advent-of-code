package main

import (
	"fmt"
	"math"

	"github.com/russellslater/advent-of-code/2022/day-15-beacon-exclusion-zone/sense"
	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-15-beacon-exclusion-zone/input.txt"
	sensors := getTransformedInput(filename)

	count := calcPositionsNotContainingBeacon(sensors, 2000000)
	fmt.Printf("Part One Answer: %d\n", count)
}

func getTransformedInput(filename string) []*sense.Sensor {
	sensors := []*sense.Sensor{}
	for _, line := range util.LoadInput(filename) {
		sensor := &sense.Sensor{}
		sensor.ClosestBeacon = &sense.Beacon{}
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.X, &sensor.Y, &sensor.ClosestBeacon.X, &sensor.ClosestBeacon.Y)
		sensors = append(sensors, sensor)
	}
	return sensors
}

func findBounds(sensors []*sense.Sensor) (int, int) {
	maxX := 0
	minX := math.MaxInt
	for _, sensor := range sensors {
		if sensor.MaxX() > maxX {
			maxX = sensor.MaxX()
		}
		if sensor.MinX() < minX {
			minX = sensor.MinX()
		}
	}

	return minX, maxX
}

func calcPositionsNotContainingBeacon(sensors []*sense.Sensor, targetY int) int {
	// Filter out sensors that are clearly out of range
	inRangeSensors := []*sense.Sensor{}
	for _, s := range sensors {
		if s.IsInVerticalRange(targetY) {
			inRangeSensors = append(inRangeSensors, s)
		}
	}

	minX, maxX := findBounds(inRangeSensors)
	noBeaconCount := 0

	for x := minX; x <= maxX; x++ {
		for _, s := range inRangeSensors {
			if s.IsBeaconLocation(x, targetY) {
				break
			}

			if s.IsInRange(x, targetY) {
				noBeaconCount++
				break
			}
		}
	}

	return noBeaconCount
}
