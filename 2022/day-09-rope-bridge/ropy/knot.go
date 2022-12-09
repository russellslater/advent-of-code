package ropy

import "github.com/russellslater/advent-of-code/internal/util"

type position struct {
	x int
	y int
}

type Knot struct {
	head   *Knot
	x, y   int
	visits map[position]bool
}

func NewKnot(head *Knot) *Knot {
	return &Knot{
		head:   head,
		x:      0,
		y:      0,
		visits: map[position]bool{},
	}
}

func (k *Knot) isTouchingHead() bool {
	dx := util.Abs(k.x - k.head.x)
	dy := util.Abs(k.y - k.head.y)
	return dx <= 1 && dy <= 1
}

func (k *Knot) move() {
	if k.isTouchingHead() {
		return
	}

	if k.x < k.head.x {
		k.x++
	} else if k.x > k.head.x {
		k.x--
	}

	if k.y < k.head.y {
		k.y++
	} else if k.y > k.head.y {
		k.y--
	}

	k.recordVisit()
}

func (k *Knot) recordVisit() {
	k.visits[position{k.x, k.y}] = true
}
