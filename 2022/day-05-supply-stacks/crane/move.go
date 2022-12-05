package crane

import "fmt"

type Move struct {
	Count  int
	Origin int
	Dest   int
}

func NewMove(instruction string) Move {
	mv := Move{}
	fmt.Sscanf(instruction, "move %d from %d to %d", &mv.Count, &mv.Origin, &mv.Dest)
	return mv
}

func BuildMoves(instructions []string) []Move {
	moves := []Move{}
	for _, ins := range instructions {
		moves = append(moves, NewMove(ins))
	}
	return moves
}
