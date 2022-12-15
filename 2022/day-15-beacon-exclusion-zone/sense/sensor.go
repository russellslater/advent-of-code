package sense

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

type Beacon struct {
	X int
	Y int
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

func (s *Sensor) IsInRange(x, y int) bool {
	return s.DistanceToBeacon() >= util.ManhattanDistance(s.X, s.Y, x, y)
}

func (s *Sensor) IsInVerticalRange(y int) bool {
	return util.Abs(s.Y-y) <= s.DistanceToBeacon()
}

func (s *Sensor) IsBeaconLocation(x, y int) bool {
	return s.ClosestBeacon.X == x && s.ClosestBeacon.Y == y
}

func (s *Sensor) MaxX() int {
	return s.X + s.DistanceToBeacon() + 1
}

func (s *Sensor) MinX() int {
	return s.X - s.DistanceToBeacon() - 1
}
