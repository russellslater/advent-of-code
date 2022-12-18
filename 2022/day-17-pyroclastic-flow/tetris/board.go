package tetris

import (
	"fmt"

	"github.com/russellslater/advent-of-code/internal/util"
)

type Board struct {
	Width        int
	Height       int
	Cells        [][]rune
	ShapeOrdinal int
}

func NewBoard(width, height int) *Board {
	b := &Board{
		Width:  width,
		Height: height,
		Cells:  make([][]rune, height),
	}
	for i := range b.Cells {
		b.Cells[i] = make([]rune, width)
	}
	return b
}

func (b *Board) NextShape() Shape {
	shape := ShapeInventory[b.ShapeOrdinal]
	b.ShapeOrdinal = (b.ShapeOrdinal + 1) % 5
	return shape
}

func (b *Board) Print() {
	for _, row := range b.Cells {
		fmt.Printf("|")
		for _, cell := range row {
			if cell == 0 {
				cell = '.'
			}
			fmt.Printf("%c", cell)
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+")
}

func (b *Board) PlaceShape(shape Shape, x, y int) {
	for i, row := range shape {
		for j, cell := range row {
			if cell == '@' {
				b.Cells[y+i][x+j] = '@'
			}
		}
	}
}

func (b *Board) RemoveShape(shape Shape, x, y int) {
	for i, row := range shape {
		for j, cell := range row {
			if cell == '@' {
				b.Cells[y+i][x+j] = '.'
			}
		}
	}
}

func (b *Board) FixShape(shape Shape, x, y int) {
	for i, row := range shape {
		for j, cell := range row {
			if cell == '@' {
				b.Cells[y+i][x+j] = '#'
			}
		}
	}
}

func (b *Board) Resize(shape Shape) {
	if b.HighestRockPosition() == -1 {
		return
	}

	heightRequired := -(b.HighestRockPosition() - shape.Height() - 3)

	// Shrink the board
	if heightRequired < 0 {
		b.Cells = b.Cells[-(heightRequired):]
		b.Height += heightRequired
		return
	}

	height := util.Abs(heightRequired)

	// Add new rows to the top of the board
	newCells := make(Shape, height)
	for i := range newCells {
		newCells[i] = make([]rune, b.Width)
	}

	b.Cells = append(newCells, b.Cells...)

	b.Height += height
}

func (b *Board) CanPlaceShape(shape Shape, x, y int) bool {
	for i, row := range shape {
		for j, cell := range row {
			if cell == '@' {
				// Not through the floor
				if y+i >= b.Height {
					return false
				}

				// Not through the side
				if x+j < 0 || x+j >= b.Width {
					return false
				}

				// Not on top of another shape
				if b.Cells[y+i][x+j] == '#' {
					return false
				}
			}
		}
	}
	return true
}

func (b *Board) HighestRockPosition() int {
	for y := range b.Cells {
		for x := range b.Cells[y] {
			if b.Cells[y][x] == '#' {
				return y
			}
		}
	}
	return -1
}

func (b *Board) MoveDown(shape Shape, x, y int) (int, int) {
	return b.moveShape(shape, x, y, 0, 1)
}

func (b *Board) MoveLeft(shape Shape, x, y int) (int, int) {
	return b.moveShape(shape, x, y, -1, 0)
}

func (b *Board) MoveRight(shape Shape, x, y int) (int, int) {
	return b.moveShape(shape, x, y, 1, 0)
}

func (b *Board) moveShape(shape Shape, x, y, dx, dy int) (int, int) {
	b.RemoveShape(shape, x, y)
	x += dx
	y += dy
	b.PlaceShape(shape, x, y)
	return x, y
}
