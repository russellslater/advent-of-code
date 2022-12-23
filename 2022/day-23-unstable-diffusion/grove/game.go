package grove

import "fmt"

type Game struct {
	startingGrid [][]rune
}

func NewGame(startingGrid [][]rune) *Game {
	return &Game{
		startingGrid: startingGrid,
	}
}

func (g *Game) Solve() int {
	fmt.Println("Grid Height:", len(g.startingGrid))
	fmt.Println("Grid Weight:", len(g.startingGrid[0]))

	for _, row := range g.startingGrid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}

	return 0
}
