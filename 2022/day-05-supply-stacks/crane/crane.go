package crane

import (
	"strings"

	"github.com/russellslater/advent-of-code/internal/util"
)

type Crane struct {
	handleMultiple bool
}

func NewCrateMover9000() *Crane {
	return &Crane{false}
}

func NewCrateMover9001() *Crane {
	return &Crane{true}
}

func (c *Crane) Operate(stacks [][]string, moves []Move) string {
	for _, mv := range moves {
		stacks = c.moveCrates(stacks, mv)
	}
	return c.result(stacks)
}

func (c *Crane) moveCrates(stacks [][]string, mv Move) [][]string {
	moved := stacks[mv.Origin-1][:mv.Count]
	if c.handleMultiple {
		moved = util.Reverse(moved)
	}
	stacks[mv.Origin-1] = stacks[mv.Origin-1][mv.Count:]       // remove from origin
	stacks[mv.Dest-1] = util.Prepend(stacks[mv.Dest-1], moved) // add to destination
	return stacks
}

func (c *Crane) result(stacks [][]string) string {
	var b strings.Builder
	for _, s := range stacks {
		b.WriteString(s[0])
	}
	return b.String()
}
