package ropy

type Rope struct {
	head  *Knot
	knots []*Knot
}

func NewRope(numKnots int) *Rope {
	r := &Rope{
		head:  NewKnot(nil),
		knots: []*Knot{},
	}

	head := r.head
	for i := 0; i < numKnots-1; i++ {
		knot := NewKnot(head)
		knot.recordVisit()
		r.knots = append(r.knots, knot)
		head = knot
	}

	return r
}

func (r *Rope) MoveHead(x int, y int, dx int, dy int) {
	for i := 0; i < x; i++ {
		r.head.x += dx
		r.moveTail()
	}

	for i := 0; i < y; i++ {
		r.head.y += dy
		r.moveTail()
	}
}

func (r *Rope) moveTail() {
	for _, knot := range r.knots {
		knot.move()
	}
}

func (r *Rope) TailPositionVisitCount() int {
	return len(r.knots[len(r.knots)-1].visits)
}
