package bliz

type Wall struct {
	Position Position
}

type WallSet map[Wall]struct{}

func (w WallSet) Contains(p Position) bool {
	_, ok := w[Wall{Position: p}]
	return ok
}

func (w WallSet) Add(p Position) {
	w[Wall{Position: p}] = struct{}{}
}
