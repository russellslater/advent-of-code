package bliz

type Position struct {
	X int
	Y int
}

func (p Position) North() Position {
	return Position{p.X, p.Y - 1}
}

func (p Position) South() Position {
	return Position{p.X, p.Y + 1}
}

func (p Position) West() Position {
	return Position{p.X - 1, p.Y}
}

func (p Position) East() Position {
	return Position{p.X + 1, p.Y}
}

func (p Position) Neighbours() []Position {
	return []Position{
		p.North(),
		p.South(),
		p.West(),
		p.East(),
	}
}
