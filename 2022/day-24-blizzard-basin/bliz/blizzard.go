package bliz

type Blizzard struct {
	Position Position
	DX       int
	DY       int
}

type BlizzardSet map[Blizzard]struct{}

func (b BlizzardSet) Add(p Position, dx int, dy int) {
	b[Blizzard{Position: p, DX: dx, DY: dy}] = struct{}{}
}

func (b BlizzardSet) Contains(o Blizzard) bool {
	_, ok := b[o]
	return ok
}
