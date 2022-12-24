package bliz

import (
	"fmt"
	"math"
)

type Basin struct {
	Start           Position
	End             Position
	InternalWidth   int
	InternalHeight  int
	Walls           WallSet
	Blizzards       BlizzardSet
	blizzardsByTime map[int]BlizzardSet
}

func NewBasin(start, end Position, walls WallSet, blizzards BlizzardSet) *Basin {
	b := &Basin{
		Start:     start,
		End:       end,
		Walls:     walls,
		Blizzards: blizzards,
	}
	b.InternalWidth, b.InternalHeight = b.internalWidthHeight()
	b.precomputeBlizzardPositions()
	return b
}

func (b *Basin) internalWidthHeight() (int, int) {
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

func (b *Basin) precomputeBlizzardPositions() {
	// Pattern convienently repeats every w * h units of time
	lcm := b.InternalWidth * b.InternalHeight

	b.blizzardsByTime = map[int]BlizzardSet{}

	for blizzard := range b.Blizzards {
		for t := 0; t < lcm; t++ {
			if _, ok := b.blizzardsByTime[t]; !ok {
				b.blizzardsByTime[t] = BlizzardSet{}
			}

			newPosX := (blizzard.Position.X + blizzard.DX*t) % b.InternalWidth
			newPosY := (blizzard.Position.Y + blizzard.DY*t) % b.InternalHeight

			// Wrap around
			if newPosX <= 0 {
				newPosX += b.InternalWidth
			}
			if newPosX > b.InternalWidth {
				newPosX %= b.InternalWidth
				newPosX++
			}

			if newPosY <= 0 {
				newPosY += b.InternalHeight
			}
			if newPosY > b.InternalHeight {
				newPosY %= b.InternalHeight
				newPosY++
			}

			b.blizzardsByTime[t].Add(Position{X: newPosX, Y: newPosY}, blizzard.DX, blizzard.DY)
		}
	}
}

func (b *Basin) PrintAtTime(time int) {
	if b.blizzardsByTime == nil {
		b.precomputeBlizzardPositions()
	}

	// Pattern repeats every w * h
	lcm := b.InternalWidth * b.InternalHeight

	if blizzards, ok := b.blizzardsByTime[time%lcm]; ok {
		b.Print(blizzards)
	}
}

func (b *Basin) Print(blizzards BlizzardSet) {
	// Print does not account for overlapping blizzards; unimportant for solution
	for y := 0; y < b.InternalHeight+2; y++ {
		for x := 0; x < b.InternalWidth+2; x++ {
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

type State struct {
	Position Position
	Time     int
}

func (b *Basin) FastestTraversal() int {
	return b.bfs(b.Start, b.End, 0)
}

func (b *Basin) bfs(start Position, end Position, time int) int {
	if b.blizzardsByTime == nil {
		b.precomputeBlizzardPositions()
	}

	queue := []State{{start, time}}
	visited := make(map[State]bool, 0)

	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]

		for _, nextState := range b.nextPossibleStates(state) {
			if visited[nextState] {
				continue
			}

			// Reached end state?
			if nextState.Position == end {
				return nextState.Time
			}

			visited[nextState] = true
			queue = append(queue, nextState)
		}
	}

	return -1
}

func (b *Basin) FastestThereAndBackTraversal() int {
	totalTime := 0
	starts := []Position{b.Start, b.End, b.Start}
	ends := []Position{b.End, b.Start, b.End}

	for i := 0; i < len(starts); i++ {
		totalTime = b.bfs(starts[i], ends[i], totalTime)
		if totalTime == -1 {
			return -1
		}
	}

	return totalTime
}

func (b *Basin) nextPossibleStates(state State) []State {
	nextStates := []State{}

	currPos := state.Position
	time := state.Time

	blizzards := b.blizzardsByTime[time+1]

	neighbours := append(currPos.Neighbours(), currPos)

	// Possible moves
	for _, neighbour := range neighbours {
		if b.Walls.Contains(neighbour) {
			continue
		}
		if blizzards.ContainsPosition(neighbour) {
			continue
		}
		// Cannot move out of bounds
		if neighbour.X < 1 || neighbour.X > b.InternalWidth || neighbour.Y < 0 || neighbour.Y > b.InternalHeight+1 {
			continue
		}
		nextStates = append(nextStates, State{neighbour, time + 1})
	}

	return nextStates
}
