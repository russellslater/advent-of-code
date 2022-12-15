package main

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

func main() {
	filename := "./2022/day-15-beacon-exclusion-zone/input.txt"
	inputLines := getTransformedInput(filename)
	output := solveProblem(inputLines)
	fmt.Printf("Solution output: %v\n", output)
}

type Sensor struct {
	X             int
	Y             int
	ClosestBeacon *Beacon
}

func (s *Sensor) DistanceToBeacon() int {
	return util.ManhattanDistance(s.X, s.Y, s.ClosestBeacon.X, s.ClosestBeacon.Y)
}

func (s *Sensor) String() string {
	return fmt.Sprintf("Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d [MD=%d]]",
		s.X, s.Y, s.ClosestBeacon.X, s.ClosestBeacon.Y, s.DistanceToBeacon())
}

type Beacon struct {
	X int
	Y int
}

func getTransformedInput(filename string) []*Sensor {
	sensors := []*Sensor{}
	for _, line := range util.LoadInput(filename) {
		sensor := &Sensor{}
		sensor.ClosestBeacon = &Beacon{}
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&sensor.X, &sensor.Y, &sensor.ClosestBeacon.X, &sensor.ClosestBeacon.Y)
		sensors = append(sensors, sensor)
	}
	return sensors
}

func solveProblem(sensors []*Sensor) interface{} {
	for _, sensor := range sensors {
		fmt.Println(sensor)
	}

	return true
}
