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

	freq := calcDistressSignalTuningFrequency(sensors, 4000000, 4000000)
	fmt.Printf("Part Two Answer: %d\n", freq)
}

func getTransformedInput(filename string) []*sense.Sensor {
	sensors := []*sense.Sensor{}
	for _, line := range util.LoadInput(filename) {
		sensor := &sense.Sensor{}
		sensor.ClosestBeacon = &sense.Beacon{}

		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.X, &sensor.Y, &sensor.ClosestBeacon.X, &sensor.ClosestBeacon.Y)

		sensor.MaxDistance = util.ManhattanDistance(sensor.X, sensor.Y, sensor.ClosestBeacon.X, sensor.ClosestBeacon.Y)

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

	// Count positions that cannot contain a beacon
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

func calcDistressSignalTuningFrequency(sensors []*sense.Sensor, maxX, maxY int) int {
	// Searching for the only location that isn't in an exclusion zone
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			isExcluded := false
			for _, s := range sensors {
				if s.IsInRange(x, y) {
					// Skip to end of sensors exclusion zone
					// Distance diminishes as y-offset increases
					skipDist := s.MaxDistance - util.Abs(s.Y-y)
					x = s.X + skipDist

					isExcluded = true
					break
				}
			}

			if !isExcluded {
				return x*4000000 + y
			}
		}
	}

	return -1 // Not found
}
