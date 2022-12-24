package bliz

import (
	"fmt"
	"math"
)

type Basin struct {
	Walls           WallSet
	Blizzards       BlizzardSet
	Start           Position
	End             Position
	BlizzardsByTime map[int]BlizzardSet
}

func NewBasin() *Basin {
	return &Basin{
		Walls:     WallSet{},
		Blizzards: BlizzardSet{},
	}
}

func (b *Basin) InternalWidthHeight() (int, int) {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt

	for wall := range b.Walls {
		if wall.Position.X < minX {
			minX = wall.Position.X
		}
		if wall.Position.X > maxX {
			maxX = wall.Position.X
		}
		if wall.Position.Y < minY {
			minY = wall.Position.Y
		}
		if wall.Position.Y > maxY {
			maxY = wall.Position.Y
		}
	}

	return maxX - minX - 1, maxY - minY - 1
}

func (b *Basin) PrecomputeBlizzardPositions() {
	w, h := b.InternalWidthHeight()
	lcm := w * h

	b.BlizzardsByTime = map[int]BlizzardSet{}

	for blizzard := range b.Blizzards {
		for t := 0; t < lcm; t++ {
			if _, ok := b.BlizzardsByTime[t]; !ok {
				b.BlizzardsByTime[t] = BlizzardSet{}
			}

			newPosX := (blizzard.Position.X + blizzard.DX*t) % w
			newPosY := (blizzard.Position.Y + blizzard.DY*t) % h

			// Wrap around
			if newPosX <= 0 {
				newPosX += 100
			}
			if newPosX > 100 {
				newPosX %= 100
				newPosX++
			}

			if newPosY <= 0 {
				newPosY += 35
			}
			if newPosY > 35 {
				newPosY %= 35
				newPosY++
			}

			b.BlizzardsByTime[t].Add(Position{X: newPosX, Y: newPosY}, blizzard.DX, blizzard.DY)
		}
	}
}

func (b *Basin) FastestTraversal() int {
	return 0
}

func (b *Basin) PrintAtTime(time int) {
	if b.BlizzardsByTime == nil {
		b.PrecomputeBlizzardPositions()
	}

	// Pattern repeats every w * h
	w, h := b.InternalWidthHeight()
	lcm := w * h

	if blizzards, ok := b.BlizzardsByTime[time%lcm]; ok {
		b.Print(blizzards)
	}
}

func (b *Basin) Print(blizzards BlizzardSet) {
	// TODO: Print not account for overlapping blizzards - likely unimportant for solution
	w, h := b.InternalWidthHeight()
	for y := 0; y < h+2; y++ {
		for x := 0; x < w+2; x++ {
			if b.Walls.Contains(Position{X: x, Y: y}) {
				fmt.Print("#")
			} else if blizzards.Contains(Blizzard{Position{X: x, Y: y}, -1, 0}) {
				fmt.Print("<")
			} else if blizzards.Contains(Blizzard{Position{X: x, Y: y}, 1, 0}) {
				fmt.Print(">")
			} else if blizzards.Contains(Blizzard{Position{X: x, Y: y}, 0, -1}) {
				fmt.Print("^")
			} else if blizzards.Contains(Blizzard{Position{X: x, Y: y}, 0, 1}) {
				fmt.Print("v")
			} else if b.Start.X == x && b.Start.Y == y {
				fmt.Print("S")
			} else if b.End.X == x && b.End.Y == y {
				fmt.Print("E")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
