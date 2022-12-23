package grove

import (
	"fmt"
	"math"
)

type Game struct {
	startingGrid     [][]rune
	currentGrid      [][]rune
	primaryDirection Direction
	proposedMoves    map[Position][]Position
	noMovesLastRound bool
}

func NewGame(startingGrid [][]rune) *Game {
	startingGridCopy := make([][]rune, len(startingGrid))
	copy(startingGridCopy, startingGrid)
	return &Game{
		startingGrid:     startingGrid,
		primaryDirection: North,
		currentGrid:      startingGridCopy,
	}
}

func (g *Game) reset() {
	startingGridCopy := make([][]rune, len(g.startingGrid))
	copy(startingGridCopy, g.startingGrid)
	g.currentGrid = startingGridCopy

	g.primaryDirection = North
	g.noMovesLastRound = false
}

func (g *Game) SolveForEmptyTiles(rounds int) int {
	g.reset()
	for i := 0; i < rounds; i++ {
		g.playRound()
	}
	return g.countEmptyTiles()
}

func (g *Game) SolveForFirstRoundNoMovement() int {
	g.reset()
	i := 0
	for {
		i++
		g.playRound()
		if g.noMovesLastRound {
			break
		}
	}
	return i
}

func (g *Game) nextDirection() Direction {
	dir := g.primaryDirection
	g.primaryDirection = (g.primaryDirection + 1) % 4
	return dir
}

func (g *Game) playRound() {
	g.growGrid()
	g.proposeThenMove()
	g.shrinkGrid()
}

func (g *Game) PrintCurrentGrid() {
	for _, row := range g.currentGrid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func (g *Game) growGrid() {
	height := len(g.currentGrid)
	width := len(g.currentGrid[0])

	newGrid := make([][]rune, height+2)
	for i := range newGrid {
		newGrid[i] = make([]rune, width+2)
	}

	// Fill the new grid with empty tiles
	for i := 0; i < height+2; i++ {
		for j := 0; j < width+2; j++ {
			newGrid[i][j] = '.'
		}
	}

	// Copy the current grid into the new grid
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newGrid[i+1][j+1] = g.currentGrid[i][j]
		}
	}

	g.currentGrid = newGrid
}

func (g *Game) shrinkGrid() {
	minX := math.MaxInt32
	maxX := 0
	minY := math.MaxInt32
	maxY := 0

	for y := 0; y < len(g.currentGrid); y++ {
		for x := 0; x < len(g.currentGrid[0]); x++ {
			if g.currentGrid[y][x] == '#' {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	newGrid := make([][]rune, maxY-minY+1)
	for i := range newGrid {
		newGrid[i] = make([]rune, maxX-minX+1)
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			newGrid[y-minY][x-minX] = g.currentGrid[y][x]
		}
	}

	g.currentGrid = newGrid
}

func (g *Game) proposeThenMove() {
	g.proposedMoves = make(map[Position][]Position)
	leadDirection := g.nextDirection()

	// Calculate proposed moves for each elf
	for y := 0; y < len(g.currentGrid); y++ {
		for x := 0; x < len(g.currentGrid[0]); x++ {
			dir := leadDirection
			// Is position occupied by an elf?
			if g.currentGrid[y][x] == '#' {
				elfPosition := Position{x, y}
				// Don't move if all neighbouring cells are clear
				if g.checkNeighbouringCellsClear(elfPosition) {
					continue
				}
			Proposal:
				// Cycle through directions in order
				for i := 0; i < 4; i++ {
					if ok := g.proposeMove(elfPosition, dir); ok {
						break Proposal
					}
					dir = (dir + 1) % 4
				}
			}
		}
	}

	// Move elves
	mvCount := 0
	for pos, elves := range g.proposedMoves {
		// If collision, don't move
		if len(elves) > 1 {
			continue
		}
		// Occupy new position and clear up old position
		elf := elves[0]
		g.currentGrid[pos.Y][pos.X] = '#'
		g.currentGrid[elf.Y][elf.X] = '.'
		mvCount++
	}

	if mvCount == 0 {
		g.noMovesLastRound = true
	}
}

func (g *Game) proposeMove(currPos Position, dir Direction) bool {
	var checkPos []Position
	var targetPos Position

	switch dir {
	case North:
		targetPos = currPos.North()
		checkPos = currPos.Northward()
	case South:
		targetPos = currPos.South()
		checkPos = currPos.Southward()
	case West:
		targetPos = currPos.West()
		checkPos = currPos.Westward()
	case East:
		targetPos = currPos.East()
		checkPos = currPos.Eastward()
	}

	for _, p := range checkPos {
		// If cell is occupied, don't move
		if g.currentGrid[p.Y][p.X] == '#' {
			return false
		}
	}
	g.proposedMoves[targetPos] = append(g.proposedMoves[targetPos], currPos)
	return true
}

func (g *Game) checkNeighbouringCellsClear(pos Position) bool {
	for _, n := range pos.Neighbours() {
		if g.currentGrid[n.Y][n.X] == '#' {
			return false
		}
	}
	return true
}

func (g *Game) countEmptyTiles() int {
	count := 0
	for y := 0; y < len(g.currentGrid); y++ {
		for x := 0; x < len(g.currentGrid[0]); x++ {
			if g.currentGrid[y][x] == '.' {
				count++
			}
		}
	}
	return count
}
