package jungle

import "fmt"

type Game struct {
	Board            [][]rune
	Moves            []Move
	PlayerX, PlayerY int
	PlayerDir        Direction
}

func NewGame(board [][]rune, moves []Move) *Game {
	return &Game{
		Board:     board,
		Moves:     moves,
		PlayerDir: Right,
	}
}

func (g *Game) PrintBoard() {
	for y, row := range g.Board {
		for x, c := range row {
			if c == 0 {
				c = ' '
			}
			if x == g.PlayerX && y == g.PlayerY {
				c = 'P'
			}
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

func (g *Game) CalcFinalPassword() int {
	g.startPlayer()
	for _, move := range g.Moves {
		if move == ClockwiseTurn {
			g.PlayerDir = (g.PlayerDir + 1) % 4
		} else if move == CounterclockwiseTurn {
			g.PlayerDir = (g.PlayerDir + 3) % 4
		} else {
			g.walkPlayer(move.Steps)
		}
	}
	return 1000*(g.PlayerY+1) + 4*(g.PlayerX+1) + int(g.PlayerDir)
}

func (g *Game) walkPlayer(steps int) {
	// Player takes steps in the direction it is facing.
	// If player hits a wall, it stops.
	// Player appears on other side of board if it goes off the end.
	for i := 0; i < steps; i++ {
		switch g.PlayerDir {
		case Right:
			g.movePlayerRight()
		case Down:
			g.movePlayerDown()
		case Left:
			g.movePlayerLeft()
		case Up:
			g.movePlayerUp()
		}
	}
}

func (g *Game) movePlayerRight() {
	// Player not facing right?
	if g.PlayerDir != Right {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerX+1 >= len(g.Board[g.PlayerY]) || g.Board[g.PlayerY][g.PlayerX+1] == 0 {
		// Make sure first open space isn't blocked
		for x := 0; x < len(g.Board[g.PlayerY]); x++ {
			if g.Board[g.PlayerY][x] != 0 {
				if g.Board[g.PlayerY][x] != '#' {
					// Wrap around
					g.PlayerX = x
				}
				return
			}
		}
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY][g.PlayerX+1] == '#' {
		return
	}

	g.PlayerX++
}

func (g *Game) movePlayerDown() {
	// Player not facing down?
	if g.PlayerDir != Down {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerY+1 >= len(g.Board) || g.Board[g.PlayerY+1][g.PlayerX] == 0 {
		// Make sure first open space isn't blocked
		for y := 0; y < len(g.Board); y++ {
			if g.Board[y][g.PlayerX] != 0 {
				if g.Board[y][g.PlayerX] != '#' {
					// Wrap around
					g.PlayerY = y
				}
				return
			}
		}
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY+1][g.PlayerX] == '#' {
		return
	}

	g.PlayerY++
}

func (g *Game) movePlayerLeft() {
	// Player not facing left?
	if g.PlayerDir != Left {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerX-1 < 0 || g.Board[g.PlayerY][g.PlayerX-1] == 0 {
		// Make sure first open space isn't blocked
		for x := len(g.Board[g.PlayerY]) - 1; x >= 0; x-- {
			if g.Board[g.PlayerY][x] != 0 {
				if g.Board[g.PlayerY][x] != '#' {
					// Wrap around
					g.PlayerX = x
				}
				return
			}
		}
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY][g.PlayerX-1] == '#' {
		return
	}

	g.PlayerX--
}

func (g *Game) movePlayerUp() {
	// Player not facing up?
	if g.PlayerDir != Up {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerY-1 < 0 || g.Board[g.PlayerY-1][g.PlayerX] == 0 {
		// Make sure first open space isn't blocked
		for y := len(g.Board) - 1; y >= 0; y-- {
			if g.Board[y][g.PlayerX] != 0 {
				if g.Board[y][g.PlayerX] != '#' {
					// Wrap around
					g.PlayerY = y
				}
				return
			}
		}
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY-1][g.PlayerX] == '#' {
		return
	}

	g.PlayerY--
}

func (g *Game) startPlayer() {
	isPlaced := false
	for y, row := range g.Board {
		for x, c := range row {
			if c == '.' {
				g.PlayerX = x
				g.PlayerY = y
				isPlaced = true
				break
			}
		}
		if isPlaced {
			break
		}
	}
}
