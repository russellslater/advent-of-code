package grove

type Position struct {
	X int
	Y int
}

func (p Position) North() Position {
	return Position{p.X, p.Y - 1}
}

func (p Position) Northwest() Position {
	return Position{p.X - 1, p.Y - 1}
}

func (p Position) Northeast() Position {
	return Position{p.X + 1, p.Y - 1}
}

func (p Position) Northward() []Position {
	return []Position{
		p.North(),
		p.Northwest(),
		p.Northeast(),
	}
}

func (p Position) South() Position {
	return Position{p.X, p.Y + 1}
}

func (p Position) Southwest() Position {
	return Position{p.X - 1, p.Y + 1}
}

func (p Position) Southeast() Position {
	return Position{p.X + 1, p.Y + 1}
}

func (p Position) Southward() []Position {
	return []Position{
		p.South(),
		p.Southwest(),
		p.Southeast(),
	}
}

func (p Position) West() Position {
	return Position{p.X - 1, p.Y}
}

func (p Position) Westward() []Position {
	return []Position{
		p.West(),
		p.Northwest(),
		p.Southwest(),
	}
}

func (p Position) East() Position {
	return Position{p.X + 1, p.Y}
}

func (p Position) Eastward() []Position {
	return []Position{
		p.East(),
		p.Northeast(),
		p.Southeast(),
	}
}

func (p Position) Neighbours() []Position {
	return []Position{
		p.North(),
		p.Northwest(),
		p.Northeast(),
		p.South(),
		p.Southwest(),
		p.Southeast(),
		p.West(),
		p.East(),
	}
}
