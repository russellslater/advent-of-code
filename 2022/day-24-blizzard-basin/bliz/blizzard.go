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

func (b BlizzardSet) ContainsPosition(p Position) bool {
	for blizzard := range b {
		if blizzard.Position == p {
			return true
		}
	}
	return false
}
