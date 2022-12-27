package twod

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

func (d Direction) Relative() RelativeDirection {
	switch d {
	case North:
		return Up
	case South:
		return Down
	case West:
		return Left
	case East:
		return Right
	default:
		return 0
	}
}

func (d Direction) Turn(turn RelativeDirection) Direction {
	switch turn {
	case Left:
		switch d {
		case North:
			return West
		case South:
			return East
		case West:
			return South
		case East:
			return North
		}
	case Right:
		switch d {
		case North:
			return East
		case South:
			return West
		case West:
			return North
		case East:
			return South
		}
	}
	return d
}

type RelativeDirection int

const (
	Up RelativeDirection = iota
	Down
	Left
	Right
)
