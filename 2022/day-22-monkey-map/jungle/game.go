package jungle

import "fmt"

type Game struct {
	Path             map[Position]Direction
	Board            [][]rune
	Moves            []Move
	PlayerX, PlayerY int
	PlayerDir        Direction
	Is3DCube         bool
}

func New3DGame(board [][]rune, moves []Move) *Game {
	return NewGame(board, moves, true)
}

func New2DGame(board [][]rune, moves []Move) *Game {
	return NewGame(board, moves, false)
}

func NewGame(board [][]rune, moves []Move, is3DCube bool) *Game {
	return &Game{
		Board:     board,
		Moves:     moves,
		PlayerDir: Right,
		Is3DCube:  is3DCube,
		Path:      make(map[Position]Direction),
	}
}

func (g *Game) PrintBoard() {
	for y, row := range g.Board {
		for x, c := range row {
			if c == 0 {
				c = ' '
			}

			// For debugging purposes, plot course of player
			if dir, ok := g.Path[Position{x, y}]; ok {
				switch dir {
				case Up:
					c = '^'
				case Down:
					c = 'v'
				case Left:
					c = '<'
				case Right:
					c = '>'
				}
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

	// For debugging purposes
	g.Path[Position{g.PlayerX, g.PlayerY}] = g.PlayerDir

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

		// For debugging purposes
		g.Path[Position{g.PlayerX, g.PlayerY}] = g.PlayerDir
	}
}

func (g *Game) movePlayerRight() {
	// Player not facing right?
	if g.PlayerDir != Right {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerX+1 >= len(g.Board[g.PlayerY]) || g.Board[g.PlayerY][g.PlayerX+1] == 0 {
		g.attemptRightWithTranspose()
		return
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY][g.PlayerX+1] == '#' {
		return
	}

	g.PlayerX++
}

func (g *Game) attemptRightWithTranspose() {
	if !g.Is3DCube {
		if newX, ok := g.checkFromLeft(g.PlayerY); ok {
			g.PlayerX = newX
		}
		return
	}

	// From Side #6 to Side #4
	if g.PlayerX == 49 && g.PlayerY >= 150 && g.PlayerY < 200 {
		newX := g.PlayerY - 100
		if newY, ok := g.checkFromBottom(newX); ok {
			g.updatePlayer(newX, newY, Up)
		}
		return
	}

	// From Side #4 to Side #1
	if g.PlayerX == 99 && g.PlayerY >= 100 && g.PlayerY < 150 {
		newY := -(g.PlayerY - 149)
		if newX, ok := g.checkFromRight(newY); ok {
			g.updatePlayer(newX, newY, Left)
		}
		return
	}

	// From Side #1 to Side #4
	if g.PlayerX == 149 && g.PlayerY >= 0 && g.PlayerY < 50 {
		newY := 149 - g.PlayerY
		if newX, ok := g.checkFromRight(newY); ok {
			g.updatePlayer(newX, newY, Left)
		}
		return
	}

	// From Side #3 to Side #1
	if g.PlayerX == 99 && g.PlayerY >= 50 && g.PlayerY < 100 {
		newX := g.PlayerY + 50
		if newY, ok := g.checkFromBottom(newX); ok {
			g.updatePlayer(newX, newY, Up)
		}
		return
	}
}

func (g *Game) movePlayerDown() {
	// Player not facing down?
	if g.PlayerDir != Down {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerY+1 >= len(g.Board) || g.Board[g.PlayerY+1][g.PlayerX] == 0 {
		g.attemptDown()
		return
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY+1][g.PlayerX] == '#' {
		return
	}

	g.PlayerY++
}

func (g *Game) attemptDown() {
	if !g.Is3DCube {
		if newY, ok := g.checkFromTop(g.PlayerX); ok {
			g.PlayerY = newY
		}
		return
	}

	// From Side #4 to Side #6
	if g.PlayerY == 149 && g.PlayerX >= 50 && g.PlayerX < 100 {
		newY := g.PlayerX + 100
		if newX, ok := g.checkFromRight(newY); ok {
			g.updatePlayer(newX, newY, Left)
		}
		return
	}

	// From Side #6 to Side #1
	if g.PlayerY == 199 && g.PlayerX >= 0 && g.PlayerX < 50 {
		newX := g.PlayerX + 100
		if newY, ok := g.checkFromTop(newX); ok {
			g.updatePlayer(newX, newY, Down)
		}
		return
	}

	// From Side #1 to Side #3
	if g.PlayerY == 49 && g.PlayerX >= 100 && g.PlayerX < 150 {
		newY := g.PlayerX - 50
		if newX, ok := g.checkFromRight(newY); ok {
			g.updatePlayer(newX, newY, Left)
		}
		return
	}
}

func (g *Game) movePlayerLeft() {
	// Player not facing left?
	if g.PlayerDir != Left {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerX-1 < 0 || g.Board[g.PlayerY][g.PlayerX-1] == 0 {
		g.attemptLeft()
		return
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY][g.PlayerX-1] == '#' {
		return
	}

	g.PlayerX--
}

func (g *Game) attemptLeft() {
	if !g.Is3DCube {
		if newX, ok := g.checkFromRight(g.PlayerY); ok {
			g.PlayerX = newX
		}
		return
	}

	// From Side #2 to Side #5
	if g.PlayerX == 50 && g.PlayerY >= 0 && g.PlayerY < 50 {
		newY := -(g.PlayerY - 149)
		if newX, ok := g.checkFromLeft(newY); ok {
			g.updatePlayer(newX, newY, Right)
		}
		return
	}

	// From Side #5 to Side #2
	if g.PlayerX == 0 && g.PlayerY >= 100 && g.PlayerY < 150 {
		newY := 149 - g.PlayerY
		if newX, ok := g.checkFromLeft(newY); ok {
			g.updatePlayer(newX, newY, Right)
		}
		return
	}

	// From Side #6 to Side #2
	if g.PlayerX == 0 && g.PlayerY >= 150 && g.PlayerY < 200 {
		newX := g.PlayerY - 100
		if newY, ok := g.checkFromTop(newX); ok {
			g.updatePlayer(newX, newY, Down)
		}
		return
	}

	// From Side #3 to Side #5
	if g.PlayerX == 50 && g.PlayerY >= 50 && g.PlayerY < 100 {
		newX := g.PlayerY - 50
		if newY, ok := g.checkFromTop(newX); ok {
			g.updatePlayer(newX, newY, Down)
		}
		return
	}
}

func (g *Game) movePlayerUp() {
	// Player not facing up?
	if g.PlayerDir != Up {
		return
	}

	// About to go off edge of board or the end of the traversable tiles?
	if g.PlayerY-1 < 0 || g.Board[g.PlayerY-1][g.PlayerX] == 0 {

		g.attemptUp()
		return
	}

	// Player blocked by wall?
	if g.Board[g.PlayerY-1][g.PlayerX] == '#' {
		return
	}

	g.PlayerY--
}

func (g *Game) attemptUp() {
	if !g.Is3DCube {
		if newY, ok := g.checkFromBottom(g.PlayerX); ok {
			g.PlayerY = newY
		}
		return
	}

	// From Side #2 to Side #6
	if g.PlayerY == 0 && g.PlayerX >= 50 && g.PlayerX < 100 {
		newY := 150 + g.PlayerX - 50
		if newX, ok := g.checkFromLeft(newY); ok {
			g.updatePlayer(newX, newY, Right)
		}
		return
	}

	// From Side #1 to Side #6
	if g.PlayerY == 0 && g.PlayerX >= 100 && g.PlayerX < 150 {
		newX := g.PlayerX - 100
		if newY, ok := g.checkFromBottom(newX); ok {
			g.updatePlayer(newX, newY, Up)
		}
		return
	}

	// From Side #5 to Side #3
	if g.PlayerY == 100 && g.PlayerX >= 0 && g.PlayerX < 50 {
		newY := g.PlayerX + 50
		if newX, ok := g.checkFromLeft(newY); ok {
			g.updatePlayer(newX, newY, Right)
		}
		return
	}
}

func (g *Game) checkFromTop(playerX int) (int, bool) {
	for y := 0; y < len(g.Board); y++ {
		if g.Board[y][playerX] != 0 {
			if g.Board[y][playerX] != '#' {
				// Wrap around
				return y, true
			}
			return 0, false
		}
	}
	return 0, false
}

func (g *Game) checkFromBottom(playerX int) (int, bool) {
	for y := len(g.Board) - 1; y >= 0; y-- {
		if g.Board[y][playerX] != 0 {
			if g.Board[y][playerX] != '#' {
				// Wrap around
				return y, true
			}
			return 0, false
		}
	}
	return 0, false
}

func (g *Game) checkFromLeft(playerY int) (int, bool) {
	for x := 0; x < len(g.Board[playerY]); x++ {
		if g.Board[playerY][x] != 0 {
			if g.Board[playerY][x] != '#' {
				// Wrap around
				return x, true
			}
			return 0, false
		}
	}
	return 0, false
}

func (g *Game) checkFromRight(playerY int) (int, bool) {
	for x := len(g.Board[playerY]) - 1; x >= 0; x-- {
		if g.Board[playerY][x] != 0 {
			if g.Board[playerY][x] != '#' {
				// Wrap around
				return x, true
			}
			return 0, false
		}
	}
	return 0, false
}

func (g *Game) updatePlayer(newX, newY int, newDir Direction) {
	g.PlayerX = newX
	g.PlayerY = newY
	g.PlayerDir = newDir
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
