package hills

type Position struct {
	X int
	Y int
}

func (t Position) Add(o Position) Position {
	return Position{t.X + o.X, t.Y + o.Y}
}
