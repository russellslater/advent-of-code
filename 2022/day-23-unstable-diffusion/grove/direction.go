package grove

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case South:
		return "South"
	case West:
		return "West"
	case East:
		return "East"
	default:
		return "Unknown"
	}
}
