package hills

type Heightmap [][]string

func (h Heightmap) Neighbours(p Position) []Position {
	moves := []Position{
		{0, -1}, // up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}
	res := make([]Position, 0, 4)
	for _, mv := range moves {
		q := p.Add(mv)
		if h.isPermissible(p, q) {
			res = append(res, q)
		}
	}
	return res
}

func (h Heightmap) isPermissible(from, to Position) bool {
	if h.isInBounds(to) {
		fromValue := h[from.Y][from.X][0]
		toValue := h[to.Y][to.X][0]

		if h[from.Y][from.X] == "S" {
			fromValue = 'a'
		}

		if h[to.Y][to.X] == "E" {
			toValue = 'z'
		}

		// can only move to a lower or equal height
		if int(toValue)-1 <= int(fromValue) {
			return true
		}
	}
	return false
}

func (h Heightmap) isInBounds(p Position) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(h) && p.X < len(h[p.Y])
}

func (h Heightmap) Cost(from, to Position) float64 {
	return 1 // all moves are of equal cost
}

func (h Heightmap) GetStartsAndDest(includeAltStarts bool) ([]Position, Position) {
	points := []Position{}
	var dest Position
	for y, row := range h {
		for x, el := range row {
			if (includeAltStarts && el == "a") || el == "S" {
				points = append(points, Position{x, y})
			} else if el == "E" {
				dest = Position{x, y}
			}
		}
	}
	return points, dest
}
